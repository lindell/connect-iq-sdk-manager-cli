package manager

import (
	"archive/zip"
	"context"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/lindell/connect-iq-sdk-manager-cli/internal/client"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type DownloadConfig struct {
	DeviceFilters DeviceFilters
}

func (m *Manager) Download(ctx context.Context, config DownloadConfig) error {
	var err error
	if ctx, err = m.setTokenToCtx(ctx); err != nil {
		return err
	}

	deviceInfos, err := client.GetDeviceInfo(ctx)
	if err != nil {
		return err
	}

	deviceInfos, err = filterDevices(deviceInfos, config.DeviceFilters)
	if err != nil {
		return err
	}

	log.Infof("Downloading %d devices.", len(deviceInfos))

	for _, device := range deviceInfos {
		log := log.WithField("device", device.Name)
		if err := m.fetchDevice(ctx, log, device); err != nil {
			return err
		}
	}

	return nil
}

func (m *Manager) fetchDevice(ctx context.Context, log log.FieldLogger, device client.DeviceInfo) error {
	rootFolder, err := connectiq.RootGarminFolder()
	if err != nil {
		return err
	}
	deviceDir := path.Join(rootFolder, "Devices", device.Name)

	if _, err := os.Stat(deviceDir); !os.IsNotExist(err) {
		log.Info("Device folder already exist")
		return nil
	}

	log.Info("Downloading device zip")
	r, err := client.DownloadDevice(ctx, device)
	if err != nil {
		return err
	}
	defer r.Close()

	// Save the zip to a temporary file
	f, err := os.CreateTemp(os.TempDir(), "device-*.zip")
	if err != nil {
		return errors.WithMessage(err, "could not create tmp device file")
	}
	defer os.Remove(f.Name())
	defer f.Close()
	if _, err := io.Copy(f, r); err != nil {
		return err
	}

	log.Info("Extracting device zip")
	return unzip(f.Name(), deviceDir)
}

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
		// Create a directory at path, including parents
		err = os.MkdirAll(path, os.ModePerm)
		if err != nil {
			return err
		}
		// If file is _supposed_ to be a directory, we're done
		if file.FileInfo().IsDir() {
			continue
		}
		// otherwise, remove that directory (_not_ including parents)
		err = os.Remove(path)
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
		_, err = io.Copy(writer, reader)
		if err != nil {
			return err
		}
	}
	return nil
}
