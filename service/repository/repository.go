package repository

import (
	"common"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"strings"
	"time"
)

var Client *mongo.Client
var Database *mongo.Database

type config struct {
	Database databaseConfig `yaml:"database"`
}

type databaseConfig struct {
	Url string `yaml:"url"`
}

func init() {
	cfg := config{}
	err := common.ReadConfig(&cfg)
	if err != nil {
		panic(err)
	}

	lastSlash := strings.LastIndex(cfg.Database.Url, "/")
	server := cfg.Database.Url[:lastSlash]
	database := cfg.Database.Url[lastSlash+1:]

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(server))
	if err != nil {
		panic(err)
	}
	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	Client = client
	Database = client.Database(database)
}
