// Package repository provides ...
package repository

import (
	"context"
	"fmt"
	"net/url"
	"time"

	"github.com/smockoro/db-faker/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoFakerRepository :
type MongoFakerRepository struct {
	Client *mongo.Client
}

// NewMongoFakerRepository :
func NewMongoFakerRepository(cfg *config.Config) FakerRepository {
	u := &url.URL{
		Scheme: cfg.Db.Schema,
		User:   url.UserPassword(cfg.Db.User, cfg.Db.Password),
		Host:   fmt.Sprintf("%s:%d", cfg.Db.Host, cfg.Db.Port),
		Path:   cfg.Db.Name,
		//RawQuery: query.Encode(),
	}
	client, _ := mongo.NewClient(options.Client().ApplyURI(u.String()))
	return &MongoFakerRepository{Client: client}
}

// Ping : ping to database
func (mfr *MongoFakerRepository) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	fmt.Println("client: ", mfr.Client)
	return mfr.Client.Ping(ctx, readpref.Primary())
}
