package config

// Default Config
type Default struct {
	ConfigName string
	ConfigPath string
}

// SetConfig for set config to read
func (d *Default) SetConfig(filename string) error {
	d.ConfigName = filename
	return nil
}

// Get key value
func (d *Default) Get(key string) interface{} {
	return PATH + d.ConfigName + ExtName
}
