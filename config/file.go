// Package config takes care of the whole configuration.
package config

import (
	"encoding/json"
	"io/ioutil"
)

// FromFile reads a configuration file and creates a new Config instance.
func FromFile(name string) *Config {
	var c Config

	file, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}

	c.path = name

	return &c
}

// ToFile writes the current configuration to the standard configuration file.
func (c *Config) ToFile() error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.path, data, 0600)
}
