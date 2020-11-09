// Package repository provides ...
package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // for postgresql driver loaded
	"github.com/smockoro/db-faker/config"
)

// PostgreSQLFakerRepository :
type PostgreSQLFakerRepository struct {
	DB *sqlx.DB
}

// NewPostgreSQLFakerRepository :
func NewPostgreSQLFakerRepository(cfg *config.Config) FakerRepository {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)
	db, _ := sqlx.Open(cfg.Db.Schema, url)
	return &PostgreSQLFakerRepository{DB: db}
}

// Ping : ping to database
func (mfr *PostgreSQLFakerRepository) Ping() error {
	return mfr.DB.Ping()
}
