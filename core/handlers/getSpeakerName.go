package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

const getSpeakerNameHandlerName = "getSpeakerName"

var GetSpeakerNameHandler ReplyHandler

func init() {
	templates := []string{
		"You are {{.User.Name}}",
		"{{.User.Name}}",
	}

	responseTemplates.RegisterHandlerResponses(
		getSpeakerNameHandlerName,
		templates,
	)

	GetSpeakerNameHandler = createReplyHandler(
		getSpeakerNameHandlerName,
		handleGetSpeakerName,
	)
}

func handleGetSpeakerName(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(getSpeakerNameHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
