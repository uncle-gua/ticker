package models

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var (
	TickerCollection *mongo.Collection
)

const (
	DBName = "ticker"
	DBHost = "mongodb://localhost:27017"
)

func Init() (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(DBHost))
	if err != nil {
		return err
	}

	TickerCollection = Client.Database(DBName).Collection("ETHUSDT_1m")

	return nil
}
