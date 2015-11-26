package config

import (
	"errors"

	"github.com/spf13/viper"
)

// Viper new struct
type Viper struct {
	ConfigName string
	ConfigPath string
	Objs       *viper.Viper
}

// SetConfig for set config to read
func (v *Viper) SetConfig(filename string) error {
	v.ConfigName = filename
	v.Objs = viper.New()
	v.Objs.SetConfigName(filename)
	v.Objs.AddConfigPath(PATH)
	v.Objs.SetConfigType("json")

	err := v.Objs.ReadInConfig()
	if err != nil {
		return errors.New(v.ConfigName + ": no such file or directory")
	}
	return nil
}

// Get key value
func (v *Viper) Get(key string) interface{} {
	return v.Objs.Get(key)
}

// Set key value - not implemented
func (v *Viper) Set(key string, value interface{}) bool {
	return true
}
