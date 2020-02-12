package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

type handlerName string

type ReplyHandler func(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string
