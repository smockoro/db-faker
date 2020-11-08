package repository_test

import (
	"testing"

	"github.com/smockoro/db-faker/config"
	"github.com/smockoro/db-faker/repository"
	"github.com/stretchr/testify/assert"
)

func TestNewMSSQLFakerRepository(t *testing.T) {
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
					Schema:   "sqlserver",
					User:     "sa",
					Password: "password",
					Host:     "localhost",
					Port:     1433,
					Name:     "",
				},
			},
			errorIsNil: true,
		},
	}

	for _, c := range cases {
		c := c // cascading
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			repo := repository.NewMSSQLFakerRepository(c.cfg)
			if c.errorIsNil {
				assert.NotNil(t, repo)
			} else {
				assert.Nil(t, repo)
			}
		})
	}
}

func TestMSSQLPing(t *testing.T) {
	cfg := &config.Config{
		Db: struct {
			Schema   string `yaml:"schema"`
			User     string `yaml:"user"`
			Password string `yaml:"password"`
			Host     string `yaml:"host"`
			Port     int    `yaml:"port"`
			Name     string `yaml:"name"`
		}{
			Schema:   "sqlserver",
			User:     "sa",
			Password: "password",
			Host:     "localhost",
			Port:     1433,
			Name:     "",
		},
	}

	repo := repository.NewMSSQLFakerRepository(cfg)
	cases := []struct {
		name       string
		errorIsNil bool
	}{
		{
			name:       "success ping",
			errorIsNil: true,
		},
	}

	for _, c := range cases {
		c := c // cascading
		t.Run(c.name, func(t *testing.T) {
			t.Parallel()
			err := repo.Ping()
			if c.errorIsNil {
				assert.Nil(t, err)
			} else {
				assert.NotNil(t, err)
			}
		})
	}
}
