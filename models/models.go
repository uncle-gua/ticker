package models

import (
	"context"
	"time"

	"github.com/uncle-gua/log"
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

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(DBHost))
	log.Fatal(err)

	TickerCollection = Client.Database(DBName).Collection("ETHUSDT_1m")
}
