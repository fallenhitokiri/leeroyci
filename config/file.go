// Config based on a JSON formatted file.
package config

import (
	"encoding/json"
	"io/ioutil"
)

// Read a configuration file and create a new Config instance.
func FromFile(name string) Config {
	var c Config

	file, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}

	return c
}

// Write the current configuration to the standard configuration file.
func (c *Config) ToFile() error {
	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.path, data, 0600)
}
