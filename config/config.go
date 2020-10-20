package config

import (
	"io/ioutil"

	"github.com/eankeen/bamboo/types"
	"github.com/pelletier/go-toml"
	"github.com/safinsingh/stat"
)

func Parse(location string, verbose bool) (types.Config, error) {
	read, err := ioutil.ReadFile(location)
	if err != nil {
		stat.Fail("Failed to read config file:", location)
	}

	var conf types.Config
	err = toml.Unmarshal(read, &conf)

	return conf, err
}
