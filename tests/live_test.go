package tests

import (
	"os"
	"testing"

	"github.com/lindell/connect-iq-sdk-manager-cli/cmd"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/stretchr/testify/assert"
)

func TestStory(t *testing.T) {
	assert.NoError(t, os.RemoveAll(connectiq.RootPath))
	username, _ := os.LookupEnv("TEST_GARMIN_USERNAME")
	password, _ := os.LookupEnv("TEST_GARMIN_PASSWORD")

	assert.Error(t, run("login", "--username", username, "--password", password))
	assert.Error(t, run("device", "list"))

	assert.NoError(t, run("agreement", "view"))
	assert.NoError(t, run("agreement", "accept"))

	assert.NoError(t, run("login", "--username", username, "--password", password))

	assert.NoError(t, run("device", "list"))

	// Download a set of devices
	assert.NoError(t, run("device", "download", "--manifest", "manifest.test.xml"))
	files, err := os.ReadDir(connectiq.DevicesPath)
	assert.NoError(t, err)
	assert.Len(t, files, 5) // Manifest file contains 5 entities

	assert.NoError(t, run("device", "download", "--include-fonts"))
	files, err = os.ReadDir(connectiq.DevicesPath)
	assert.NoError(t, err)
	assert.Greater(t, len(files), 5)
	files, err = os.ReadDir(connectiq.FontsPath)
	assert.NoError(t, err)
	assert.Greater(t, len(files), 10)

	assert.NoError(t, run("sdk", "download", "^6.0.0"))
}

func run(args ...string) error {
	command := cmd.RootCmd()
	command.SetArgs(args)
	return command.Execute()
}
