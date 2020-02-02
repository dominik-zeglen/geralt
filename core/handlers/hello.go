package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

const helloHandlerName = "hello"

func init() {
	templates := []string{
		"Hi",
		"Hello",
		"Oh, hi",
		"Hey",
	}

	responseTemplates.RegisterHandlerResponses(helloHandlerName, templates)
}

func HandleHello(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(helloHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
