// Package repository provides ...
package repository

import (
	"context"
	"fmt"
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
	uri := fmt.Sprintf("mongodb://%s:%s@%s:%d/%s",
		cfg.Db.User, cfg.Db.Password, cfg.Db.Host, cfg.Db.Port, cfg.Db.Name)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	return &MongoFakerRepository{Client: client}
}

// Ping : ping to database
func (mfr *MongoFakerRepository) Ping() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	return mfr.Client.Ping(ctx, readpref.Primary())
}
