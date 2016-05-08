// Package config contains all data models used for LeeroyCI.
package config

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"strings"
)

var defaultConfig = `
{
	"url": "http://localhost",
	"port": 8000,
	"parallel_builds": 1
}
`

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
}

// NewConfig initializes LeeroyCIs configuration.
func NewConfig(path string) (*Config, error) {
	var config *Config
	cfgDecoder, err := getConfigDecoder(path)

	if err != nil {
		return nil, err
	}

	if err := cfgDecoder.Decode(&config); err != nil {
		return nil, err
	}

	config.cfgPath = path
	return config, nil
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

	cfgDecoder := json.NewDecoder(cfgFile)
	return cfgDecoder, nil
}

// Save saves the configuration file to c.cfgPath
func (c *Config) Save() error {
	data, err := json.Marshal(c)

	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(c.cfgPath, data, 0600); err != nil {
		return err
	}

	return nil
}
