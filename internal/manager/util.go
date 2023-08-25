package manager

import (
	"archive/zip"
	"crypto/md5" //nolint:gosec
	"fmt"
	"io"
	"os"
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

func fetchAndExtract(r io.Reader, destination string) (md5sum string, err error) {
	// Save the zip to a temporary file
	f, err := os.CreateTemp(os.TempDir(), "*.zip")
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

	logrus.
		WithField("hash", md5hash).
		WithField("destination", destination).
		Debugf("Unzipping file")

	return md5hash, unzip(f.Name(), destination)
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
