package main

import (
	"context"

	"github.com/dominik-zeglen/geralt/models"
	migrate "github.com/xakep666/mongo-migrate"
	"go.mongodb.org/mongo-driver/mongo"
)

func init() {
	migrate.Register(func(db *mongo.Database) error {
		bot := models.Bot{
			ID:   models.BotCollectionID,
			Name: "Geralt",
		}

		_, err := db.
			Collection(models.GlobalsCollectionKey).
			InsertOne(context.TODO(), &bot)

		user := models.User{
			Email: "admin@example.com",
			Name:  "Admin",
		}

		_, err = db.
			Collection(models.UsersCollectionKey).
			InsertOne(context.TODO(), &user)

		return err
	}, func(db *mongo.Database) error {
		return nil
	})
}
