package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

type handlerName string

const handlerOkSuffix = "Ok"
const handlerNotOkSuffix = "NotOk"

type replyHandlerFunc func(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string

type replyHandlerData struct {
	handlerFunc replyHandlerFunc
	name        string
}

func (h replyHandlerData) Exec(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	return h.handlerFunc(ctx, db, sentence)
}
func (h replyHandlerData) GetName() string {
	return h.name
}

type ReplyHandler interface {
	Exec(
		ctx context.Context,
		db *mongo.Database,
		sentence parser.ParsedSentence,
	) string
	GetName() string
}

func createReplyHandler(
	name string,
	handlerFunc replyHandlerFunc,
) ReplyHandler {
	handler := replyHandlerData{
		handlerFunc: handlerFunc,
		name:        name,
	}

	return handler
}
