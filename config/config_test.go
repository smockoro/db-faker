package config_test

import (
	"testing"

	"github.com/smockoro/db-faker/config"
	"github.com/stretchr/testify/assert"
)

func TestReadYamlConfigFile(t *testing.T) {
	cases := []struct {
		name       string
		path       string
		cfg        *config.Config
		errorIsNil bool
	}{
		{
			name: "success read yml file",
			path: "../testdata/mysql_test.yml",
			cfg: &config.Config{
				Db: struct {
					Schema   string `yaml:"schema"`
					User     string `yaml:"user"`
					Password string `yaml:"password"`
					Host     string `yaml:"host"`
					Port     int    `yaml:"port"`
					Name     string `yaml:"name"`
				}{
					Schema:   "mysql",
					User:     "root",
					Password: "password",
					Host:     "localhost",
					Port:     3306,
					Name:     "sampleDB",
				},
			},
			errorIsNil: true,
		},
		{
			name:       "yml file not found",
			path:       "../testdata/not_found.yml",
			errorIsNil: false,
		},
		{
			name:       "invalid yml structure",
			path:       "../testdata/invalid_structure.yml",
			errorIsNil: false,
		},
	}

	for _, c := range cases {
		c := c // cascading
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			cfg, err := config.ReadYamlConfigFile(c.path)
			if c.errorIsNil {
				assert.Nil(t, err)
				assert.EqualValues(t, c.cfg, cfg)
			} else {
				assert.NotNil(t, err)
				assert.Nil(t, cfg)
			}
		})
	}
}
