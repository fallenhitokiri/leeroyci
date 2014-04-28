// Config based on a JSON formatted file.
package config

import (
	"encoding/json"
	"io/ioutil"
)

func FromFile(name string) Config {
	var c Config

	file, err := ioutil.ReadFile(name)
	if err != nil {
		panic(err)
	}

	json.Unmarshal(file, &c)

	return c
}
