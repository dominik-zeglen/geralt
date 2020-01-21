package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BaseModel struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`
}

const GlobalsCollectionKey = "globals"
