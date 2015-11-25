package config

import (
	"errors"
	"os"
)

// Default Config
type Default struct {
	ConfigName string
	ConfigPath string
}

// SetConfig for set config to read
func (d *Default) SetConfig(filename string) error {
	d.ConfigName = PATH + filename + ExtName

	_, err := os.Stat(d.ConfigName)
	if err == nil {
		return nil
	}

	if os.IsNotExist(err) {
		return errors.New(d.ConfigName + ": no such file or directory")
	}
	return nil
}

// Get key value
func (d *Default) Get(key string) interface{} {
	return PATH + d.ConfigName
}
