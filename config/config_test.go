package config

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestNewConfigDefault(t *testing.T) {
	cfg, err := NewConfig("")

	if err != nil {
		t.Error(err.Error())
	}

	if cfg.URL != "http://localhost" {
		t.Error("Wrong URL", cfg.URL)
	}

	if cfg.ParallelBuilds != 1 {
		t.Error("Wrong value for parallel builds", cfg.ParallelBuilds)
	}

	if cfg.cfgPath != "" {
		t.Error("cfgPath set", cfg.cfgPath)
	}
}

func TestNewConfig(t *testing.T) {
	dir, err := ioutil.TempDir("", "leeroy_config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	tmpCfg := filepath.Join(dir, "config.json")
	if err = ioutil.WriteFile(tmpCfg, []byte(defaultConfig), 0666); err != nil {
		t.Fatal(err)
	}

	cfg, err := NewConfig(tmpCfg)

	if err != nil {
		t.Error(err.Error())
	}

	if cfg.URL != "http://localhost" {
		t.Error("Wrong URL", cfg.URL)
	}

	if cfg.ParallelBuilds != 1 {
		t.Error("Wrong value for parallel builds", cfg.ParallelBuilds)
	}

	if cfg.cfgPath == "" {
		t.Error("No cfgPath set")
	}
}

func TestNewConfigInvalidJSON(t *testing.T) {
	dir, err := ioutil.TempDir("", "leeroy_config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	tmpCfg := filepath.Join(dir, "config.json")
	if err = ioutil.WriteFile(tmpCfg, []byte("fj28:"), 0666); err != nil {
		t.Fatal(err)
	}

	_, err = NewConfig(tmpCfg)

	if err == nil {
		t.Error("No error for invalid JSON")
	}
}

func TestNewConfigCannotReadFile(t *testing.T) {
	dir, err := ioutil.TempDir("", "leeroy_config_test")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(dir)

	tmpCfg := filepath.Join(dir, "config.json")
	if err = ioutil.WriteFile(tmpCfg, []byte("fj28:"), 0000); err != nil {
		t.Fatal(err)
	}

	_, err = NewConfig(tmpCfg)

	if err == nil {
		t.Error("No error for config file that cannot be read")
	}
}
