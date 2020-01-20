package middleware

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const UserContextKey = MiddlewareKey("user")

type UserData struct {
	ID    primitive.ObjectID `bson:"_id, omitempty"`
	Email string
	Name  string
}

func WithUser(ctx context.Context, db *mongo.Database) context.Context {
	user := ctx.Value(UserContextKey).(UserData)

	collection := db.Collection("users")
	err := collection.FindOne(context.TODO(), bson.M{
		"_id": user.ID,
	}).Decode(&user)

	if err != nil {
		panic(err)
	}

	return context.WithValue(ctx, UserContextKey, UserData{
		Name: user.Name,
	})
}
