package tests

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"

	"github.com/lindell/connect-iq-sdk-manager-cli/cmd"
	"github.com/lindell/connect-iq-sdk-manager-cli/internal/connectiq"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

	assert.NoError(t, run("device", "download", "--manifest", "manifest.test.xml", "--include-fonts"))
	files, err = os.ReadDir(connectiq.FontsPath)
	assert.NoError(t, err)
	assert.Greater(t, len(files), 10)

	assert.NoError(t, run("sdk", "set", "^8.0.0"))

	_, _, err = execCmd("go", "install", "..")
	require.NoError(t, err)

	stdOut, _, err := execCmd("connect-iq-sdk-manager-cli", "sdk", "current-path", "--bin")
	require.NoError(t, err)
	binPath := strings.TrimSpace(stdOut)

	monkeyCBin := filepath.Join(binPath, "./monkeyc")
	_, _, err = execCmd(monkeyCBin, "--help")
	require.NoError(t, err)
}

func run(args ...string) error {
	command := cmd.RootCmd()
	command.SetArgs(args)
	return command.Execute()
}

func execCmd(cmd string, args ...string) (string, string, error) {
	c := exec.Command(cmd, args...)
	stdout := bytes.NewBuffer(nil)
	stderr := bytes.NewBuffer(nil)
	c.Stdout = stdout
	c.Stderr = stderr
	err := c.Run()
	return stdout.String(), stderr.String(), err
}
