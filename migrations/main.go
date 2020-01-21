package main

import (
	"context"
	"time"

	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	opt := options.
		Client().
		ApplyURI("mongodb://localhost:27017").
		SetAuth(options.Credential{
			Username: "geralt",
			Password: "geralt",
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

	db := client.Database("geralt")
	migrate.SetDatabase(db)

	if err := migrate.Up(migrate.AllAvailable); err != nil {
		panic(err)
	}
}
