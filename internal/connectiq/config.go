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

type ConfigEntity struct {
	Key   string
	Value string
}

func StoreConfigKeyVal(key, val string) error {
	return StoreConfigKeyVals(ConfigEntity{key, val})
}

func StoreConfigKeyVals(values ...ConfigEntity) error {
	configMutex.Lock()
	defer configMutex.Unlock()

	cfg, err := ini.Load(ConfigPath)
	if os.IsNotExist(err) {
		cfg = ini.Empty()
	} else if err != nil {
		return errors.WithMessage(err, "could not open config file")
	}

	for _, v := range values {
		cfg.Section("").Key(v.Key).SetValue(v.Value)
	}

	if err := cfg.SaveTo(ConfigPath); err != nil {
		return errors.WithMessage(err, "could not write to config file")
	}

	return nil
}

func LoadConfigVals(keys ...string) []string {
	configMutex.Lock()
	defer configMutex.Unlock()

	cfg, err := ini.Load(ConfigPath)
	if os.IsNotExist(err) {
		cfg = ini.Empty()
	}

	ret := []string{}
	for _, key := range keys {
		ret = append(ret, cfg.Section("").Key(key).String())
	}
	return ret
}
