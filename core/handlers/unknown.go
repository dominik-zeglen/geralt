package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

const unknownHandlerName = "unknown"

var UnknownHandler ReplyHandler

func init() {
	templates := []string{
		"I don't understand",
		"Can you rephrase? I'm still learning",
		"Sorry, I don't understand",
		"I'm not sure what do you mean",
	}

	responseTemplates.RegisterHandlerResponses(unknownHandlerName, templates)

	UnknownHandler = createReplyHandler(unknownHandlerName, handleUnknown)
}

func handleUnknown(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(unknownHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
