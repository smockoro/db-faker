// Package repository provides ...
package repository

import (
	"fmt"
	"net/url"

	_ "github.com/jackc/pgx/stdlib" // for postgresql driver loaded
	"github.com/jmoiron/sqlx"
	"github.com/smockoro/db-faker/config"
)

// PostgreSQLFakerRepository :
type PostgreSQLFakerRepository struct {
	DB *sqlx.DB
}

// NewPostgreSQLFakerRepository :
func NewPostgreSQLFakerRepository(cfg *config.Config) FakerRepository {
	u := &url.URL{
		Scheme: cfg.Db.Schema,
		User:   url.UserPassword(cfg.Db.User, cfg.Db.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Db.Host, cfg.Db.Port),
		Path:   cfg.Db.Name,
		//RawQuery: query.Encode(),
	}
	db, _ := sqlx.Open("pgx", u.String())
	return &PostgreSQLFakerRepository{DB: db}
}

// Ping : ping to database
func (mfr *PostgreSQLFakerRepository) Ping() error {
	return mfr.DB.Ping()
}
