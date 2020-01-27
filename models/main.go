package models

import (
	"github.com/dominik-zeglen/geralt/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BaseModel struct {
	ID primitive.ObjectID `bson:"_id, omitempty"`
}

const GlobalsCollectionKey = "globals"

type DBConfig struct {
	Username string
	Password string
	Hostname string
	DBName   string
}

func GetDBConfig() DBConfig {
	return DBConfig{
		Username: utils.GetEnvOrPanic("DB_USERNAME"),
		Password: utils.GetEnvOrPanic("DB_PASSWORD"),
		Hostname: utils.GetEnvOrPanic("DB_HOSTNAME"),
		DBName:   utils.GetEnvOrPanic("DB_NAME"),
	}
}
