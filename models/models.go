package models

import (
	"context"
	"time"

	"ticker/config"

	"github.com/uncle-gua/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Client *mongo.Client
var (
	Database *mongo.Database
)

func init() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var err error
	Client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Database.Host))
	if err != nil {
		log.Fatal(err)
	}

	Database = Client.Database(config.Database.Name)
}
