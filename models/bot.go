package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const BotCollectionID = "bot"

type Bot struct {
	ID   string `bson:"_id"`
	Name string
}

func getBot(db *mongo.Database) (Bot, error) {
	bot := Bot{}
	collection := db.Collection(GlobalsCollectionKey)
	err := collection.FindOne(context.TODO(), bson.M{
		"_id": BotCollectionID,
	}).Decode(&bot)

	return bot, err
}
