// Package repository provides ...
package repository

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql" // for mysql driver loaded
	"github.com/jmoiron/sqlx"
	"github.com/smockoro/db-faker/config"
)

// MySQLFakerRepository :
type MySQLFakerRepository struct {
	DB *sqlx.DB
}

// NewMySQLFakerRepository :
func NewMySQLFakerRepository(cfg *config.Config) FakerRepository {
	dsn := fmt.Sprintf("%s:%s@%s:%d/%s",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)
	db, _ := sqlx.Open(cfg.Db.Schema, dsn)
	return &MySQLFakerRepository{DB: db}
}

// Ping : ping to database
func (mfr *MySQLFakerRepository) Ping() error {
	return mfr.DB.Ping()
}
