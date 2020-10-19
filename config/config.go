package config

import (
	"io/ioutil"

	"github.com/BurntSushi/toml"
	"github.com/eankeen/bamboo/types"
	"github.com/safinsingh/stat"
)

func Parse(location string) types.Config {
	read, err := ioutil.ReadFile(location)
	if err != nil {
		stat.Fail("Failed to read config file: " + location)
	}

	var conf types.Config
	if _, err := toml.Decode(string(read), &conf); err != nil {
		stat.Fail("Failed to unmarshal TOML data: " + err.Error())
	}
	return conf
}
