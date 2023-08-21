package connectiq

import (
	"os"
	"sync"

	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
)

func init() {
	// For some reason the ini package does only allow this to be set globally
	ini.PrettyFormat = false
}

var configMutex = sync.Mutex{}

func StoreConfigKeyVal(key, val string) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	cfg, err := ini.Load(ConfigPath)
	if os.IsNotExist(err) {
		cfg = ini.Empty()
	} else if err != nil {
		return errors.WithMessage(err, "could not open config file")
	}

	cfg.Section("").Key(key).SetValue(val)
	if err := cfg.SaveTo(ConfigPath); err != nil {
		return errors.WithMessage(err, "could not write to config file")
	}

	return nil
}
