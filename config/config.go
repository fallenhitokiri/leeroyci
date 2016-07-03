// Package config contains all data models used for LeeroyCI.
package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"
)

var cfg *Config

var defaultConfig = `
{
	"url": "http://localhost",
	"port": 8000,
	"parallel_builds": 1
}
`

const errorNoConfigPath = "No config path"

// Config stores everything configuraiton related to the process, including
// users and job configurations.
type Config struct {
	URL            string `json:"url"`
	Port           int    `json:"port"`
	Secret         string `json:"secret"`
	SSLCert        string `json:"ssl_cert"`
	SSLKey         string `json:"ssl_key"`
	ParallelBuilds int    `json:"parallel_builds"`
	cfgPath        string

	Users []*User `json:"users"`

	MailServer *MailServer `json:"mail_server"`

	Projects []*Project `json:"projects"`

	inMemory bool
}

// NewConfig initializes LeeroyCIs configuration.
func NewConfig(path string) error {
	cfgDecoder, err := getConfigDecoder(path)

	if err != nil {
		return err
	}

	if err := cfgDecoder.Decode(&cfg); err != nil {
		return err
	}

	cfg.inMemory = false
	cfg.cfgPath = path
	return nil
}

// Returns a json.Decoder for the configurations or an error.
func getConfigDecoder(path string) (*json.Decoder, error) {
	// if no config exists use the default config
	if _, err := os.Stat(path); os.IsNotExist(err) {
		cfgReader := strings.NewReader(defaultConfig)
		cfgDecoder := json.NewDecoder(cfgReader)
		return cfgDecoder, nil
	}

	cfgFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	return json.NewDecoder(cfgFile), nil
}

// Save saves the configuration file to c.cfgPath
func (c *Config) Save() error {
	if c.inMemory == true {
		return nil
	}

	if c.cfgPath == "" {
		return errors.New(errorNoConfigPath)
	}

	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	return ioutil.WriteFile(c.cfgPath, data, 0600)
}
