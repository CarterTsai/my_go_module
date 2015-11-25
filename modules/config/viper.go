package config

// Viper new struct
type Viper struct {
	ConfigName string
	ConfigPath string
}

// SetConfig for set config to read
func (v *Viper) SetConfig(filename string) error {
	v.ConfigName = filename
	return nil
}

// Get key value
func (v *Viper) Get(key string) interface{} {
	return PATH + v.ConfigName + ExtName
}
