package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

const setBotNameHandlerName = "setBotName"

func init() {
	templates := []string{
		"Ok, how am I going to be called?",
		"Sure, to what?",
	}

	responseTemplates.RegisterHandlerResponses(setBotNameHandlerName, templates)
}

func HandleSetBotName(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	user := GetUserFromContext(ctx)
	user.FlowState.Event(flow.ToBotNameSetting.String())

	tmpl := responseTemplates.GetRandomResponse(setBotNameHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
