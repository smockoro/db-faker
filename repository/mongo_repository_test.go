package repository_test

import (
	"testing"

	"github.com/smockoro/db-faker/config"
	"github.com/smockoro/db-faker/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewMongoFakerRepository(t *testing.T) {
	cases := []struct {
		name       string
		cfg        *config.Config
		errorIsNil bool
	}{
		{
			name: "success create repository",
			cfg: &config.Config{
				Db: struct {
					Schema   string `yaml:"schema"`
					User     string `yaml:"user"`
					Password string `yaml:"password"`
					Host     string `yaml:"host"`
					Port     int    `yaml:"port"`
					Name     string `yaml:"name"`
				}{
					Schema:   "mongodb",
					User:     "admin",
					Password: "P@ssw0rd",
					Host:     "localhost",
					Port:     27017,
					Name:     "admin",
				},
			},
			errorIsNil: true,
		},
	}

	for _, c := range cases {
		c := c // cascading
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			repo := repository.NewMongoFakerRepository(c.cfg)
			if c.errorIsNil {
				assert.NotNil(t, repo)
			} else {
				assert.Nil(t, repo)
			}
		})
	}
}

func TestMongoPing(t *testing.T) {

	cases := []struct {
		name       string
		cfg        *config.Config
		errorIsNil bool
	}{
		{
			name: "success ping",
			cfg: &config.Config{
				Db: struct {
					Schema   string `yaml:"schema"`
					User     string `yaml:"user"`
					Password string `yaml:"password"`
					Host     string `yaml:"host"`
					Port     int    `yaml:"port"`
					Name     string `yaml:"name"`
				}{
					Schema:   "mongodb",
					User:     "admin",
					Password: "P@ssw0rd",
					Host:     "localhost",
					Port:     27017,
					Name:     "admin",
				},
			},
			errorIsNil: true,
		},
		{
			name: "disconnect",
			cfg: &config.Config{
				Db: struct {
					Schema   string `yaml:"schema"`
					User     string `yaml:"user"`
					Password string `yaml:"password"`
					Host     string `yaml:"host"`
					Port     int    `yaml:"port"`
					Name     string `yaml:"name"`
				}{
					Schema:   "mongodb",
					User:     "",
					Password: "",
					Host:     "localhost",
					Port:     7017,
					Name:     "",
				},
			},
			errorIsNil: false,
		},
	}

	for _, c := range cases {
		c := c // cascading
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			repo := repository.NewMongoFakerRepository(c.cfg)
			err := repo.Ping()
			if c.errorIsNil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
