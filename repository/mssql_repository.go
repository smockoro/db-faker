// Package repository provides ...
package repository

import (
	"fmt"
	"net/url"

	_ "github.com/denisenkom/go-mssqldb" // for mssql driver import
	"github.com/jmoiron/sqlx"
	"github.com/smockoro/db-faker/config"
)

// MSSQLFakerRepository :
type MSSQLFakerRepository struct {
	DB *sqlx.DB
}

// NewMSSQLFakerRepository :
func NewMSSQLFakerRepository(cfg *config.Config) FakerRepository {
	query := url.Values{}
	query.Add("database", cfg.Db.Name)

	u := &url.URL{
		Scheme:   cfg.Db.Schema,
		User:     url.UserPassword(cfg.Db.User, cfg.Db.Password),
		Host:     fmt.Sprintf("%s:%d", cfg.Db.Host, cfg.Db.Port),
		RawQuery: query.Encode(),
	}
	db, _ := sqlx.Open(cfg.Db.Schema, u.String())
	return &MSSQLFakerRepository{DB: db}
}

// Ping : ping to database
func (mfr *MSSQLFakerRepository) Ping() error {
	return mfr.DB.Ping()
}
