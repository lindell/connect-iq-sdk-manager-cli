package manager

import (
	"archive/zip"
	"crypto/md5" //nolint:gosec
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func unzip(source, destination string) error {
	archive, err := zip.OpenReader(source)
	if err != nil {
		return err
	}
	defer archive.Close()
	for _, file := range archive.Reader.File {
		reader, err := file.Open()
		if err != nil {
			return err
		}
		defer reader.Close()

		path := filepath.Join(destination, file.Name) // nolint: gosec
		if strings.Contains(path, "..") {
			continue
		}

		// Remove file if it already exists; no problem if it doesn't; other cases can error out below
		_ = os.Remove(path)

		// Create a directory, including parents
		if file.FileInfo().IsDir() {
			err = os.MkdirAll(path, os.ModePerm)
			if err != nil {
				return err
			}
			continue
		}

		// Create a directory at path, including parents
		err = os.MkdirAll(filepath.Dir(path), os.ModePerm)
		if err != nil {
			return err
		}

		// and create the actual file.  This ensures that the parent directories exist!
		// An archive may have a single file with a nested path, rather than a file for each parent dir
		writer, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, file.Mode())
		if err != nil {
			return err
		}
		defer writer.Close()
		_, err = io.Copy(writer, reader) //nolint:gosec
		if err != nil {
			return err
		}
	}
	return nil
}

func fetchAndExtract(r io.Reader, destination string, filename string) (md5sum string, err error) {
	// Save the zip/dmg to a temporary file
	ext := filepath.Ext(filename)
	f, err := os.CreateTemp(os.TempDir(), "*"+ext)
	if err != nil {
		return "", errors.WithMessage(err, "could not create tmp device file")
	}
	defer os.Remove(f.Name())
	defer f.Close()

	logrus.Debug("Downloading zip to temporary a location")
	m := md5.New() //nolint:gosec
	w := io.MultiWriter(f, m)
	if _, err := io.Copy(w, r); err != nil {
		return "", err
	}
	md5hash := fmt.Sprintf("%x", m.Sum(nil))

	logger := logrus.
		WithField("hash", md5hash).
		WithField("destination", destination)

	if ext == ".dmg" {
		logger.Debugf("Installing DMG file")
		return md5hash, installDMG(f.Name(), destination)
	}

	logger.Debugf("Unzipping file")
	return md5hash, unzip(f.Name(), destination)
}

func installDMG(source, destination string) error {
	// Mount the DMG
	// hdiutil attach -nobrowse -mountpoint /path/to/mountpoint /path/to/dmg
	mountPoint, err := os.MkdirTemp("", "dmg-mount")
	if err != nil {
		return errors.Wrap(err, "failed to create mount point")
	}
	defer os.Remove(mountPoint)

	cmd := exec.Command("hdiutil", "attach", "-nobrowse", "-mountpoint", mountPoint, source)
	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrapf(err, "failed to mount dmg: %s", string(output))
	}
	defer func() {
		// Detach the DMG
		// hdiutil detach /path/to/mountpoint
		cmd := exec.Command("hdiutil", "detach", mountPoint, "-force")
		if output, err := cmd.CombinedOutput(); err != nil {
			logrus.Errorf("failed to detach dmg: %s", string(output))
		}
	}()

	// Copy the contents
	// For SDKs, there is usually a folder inside the DMG. We want to copy the contents
	// of that folder to destination. Because of how unzip works (it unzips the
	// folder structure), we expect 'destination' to be the directory that WILL
	// contain the SDK contents. But the SDK zip/dmg usually contains a root folder
	// (e.g. connectiq-sdk-mac-4.1.4). So we need to look into the mount point,
	// find the single directory there, and copy IT explicitly.

	failed := true
	defer func() {
		if failed {
			os.RemoveAll(destination)
		}
	}()

	entries, err := os.ReadDir(mountPoint)
	if err != nil {
		return errors.Wrap(err, "failed to read mount point")
	}

	// Filter out hidden files
	var visibleEntries []os.DirEntry
	for _, entry := range entries {
		if !strings.HasPrefix(entry.Name(), ".") {
			visibleEntries = append(visibleEntries, entry)
		}
	}

	if len(visibleEntries) != 1 || !visibleEntries[0].IsDir() {
		return errors.Errorf("expected exactly one directory in DMG, found %d", len(visibleEntries))
	}

	srcDir := filepath.Join(mountPoint, visibleEntries[0].Name())

	// Copy the directory found in the DMG to the destination path.
	// This mimics the behavior of unzip by placing the SDK root folder at the
	// destination.

	if err := os.MkdirAll(filepath.Dir(destination), 0755); err != nil {
		return errors.Wrap(err, "failed to create destination parent dir")
	}

	cmd = exec.Command("cp", "-R", srcDir, destination)
	if output, err := cmd.CombinedOutput(); err != nil {
		return errors.Wrapf(err, "failed to copy files: %s", string(output))
	}

	failed = false
	return nil
}

func isNotFound(err error) bool {
	if err == nil {
		return false
	}

	if errNotFound, ok := err.(interface{ NotFound() bool }); ok && errNotFound.NotFound() {
		return true
	}

	return isNotFound(errors.Unwrap(err))
}
