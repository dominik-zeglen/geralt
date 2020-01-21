package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
)

const unknownHandlerName = "unknown"

func init() {
	templates := []string{
		"I don't understand",
		"Can you rephrase? I'm still learning",
		"Sorry, I don't understand",
		"I'm not sure what do you mean",
	}
	responseTemplates.RegisterHandlerResponses(unknownHandlerName, templates)
}

func HandleUnknown(
	ctx context.Context,
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(unknownHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}