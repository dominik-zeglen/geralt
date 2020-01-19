package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
)

const getBotNameHandlerName = "Hello"

func init() {
	templates := []string{
		"I'm {{.Bot.Name}}",
		"My name is {{.Bot.Name}}",
		"I'm {{.Bot.Name}}, autonomic bot",
	}

	responseTemplates.RegisterHandlerResponses(getBotNameHandlerName, templates)
}

func HandleGetBotName(
	ctx context.Context,
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(getBotNameHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
