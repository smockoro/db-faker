// Package repository provides ...
package repository

import (
	"fmt"
	"strings"

	"github.com/elastic/go-elasticsearch/v7"
	"github.com/smockoro/db-faker/config"
)

// ElasticsearchFakerRepository :
type ElasticsearchFakerRepository struct {
	Client *elasticsearch.Client
}

// NewElasticsearchFakerRepository :
func NewElasticsearchFakerRepository(cfg *config.Config) FakerRepository {
	hosts := make([]string, 0)
	for _, v := range strings.Split(cfg.Db.Host, ",") {
		hosts = append(hosts, fmt.Sprintf("http://%s:%d", v, cfg.Db.Port))
	}

	ecfg := elasticsearch.Config{
		Addresses: hosts,
		Username:  cfg.Db.User,
		Password:  cfg.Db.Password,
	}

	client, _ := elasticsearch.NewClient(ecfg)
	return &ElasticsearchFakerRepository{Client: client}
}

// Ping : ping to database
func (mfr *ElasticsearchFakerRepository) Ping() error {
	_, err := mfr.Client.Ping()
	return err
}
