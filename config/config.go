package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

// Config : yaml config
type Config struct {
	Db struct {
		Schema   string `yaml:"schema"`
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Host     string `yaml:"host"`
		Port     int    `yaml:"port"`
		Name     string `yaml:"name"`
	} `yaml:"db"`
}

// ReadYamlConfigFile :
func ReadYamlConfigFile(path string) (*Config, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	cfg := &Config{}
	err = yaml.Unmarshal(buf, cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
