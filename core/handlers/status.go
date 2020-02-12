package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

const statusHandlerName = "status"

func init() {
	templates := []string{
		"Nothing much",
		"Everything's fine, thanks",
		"It's okay",
		"Could have been better, but it's okay",
	}

	responseTemplates.RegisterHandlerResponses(statusHandlerName, templates)
}

func HandleStatus(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(statusHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
