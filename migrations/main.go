package main

import (
	"context"
	"time"

	"github.com/dominik-zeglen/geralt/models"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	dbConfig := models.GetDBConfig()

	opt := options.
		Client().
		ApplyURI(dbConfig.Hostname).
		SetAuth(options.Credential{
			Username: dbConfig.Username,
			Password: dbConfig.Password,
		})

	client, err := mongo.NewClient(opt)
	if err != nil {
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		panic(err)
	}

	db := client.Database(dbConfig.DBName)
	migrate.SetDatabase(db)

	if err := migrate.Up(migrate.AllAvailable); err != nil {
		panic(err)
	}
}
