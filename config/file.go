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

	err = json.Unmarshal(file, &c)
	if err != nil {
		panic(err)
	}

	return c
}
