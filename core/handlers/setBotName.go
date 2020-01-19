package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
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
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(setBotNameHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
