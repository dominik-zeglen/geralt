package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
)

const setSpeakerNameHandlerName = "setSpeakerName"

func init() {
	templates := []string{
		"Nice to meet you, {{.User.Name}}",
		"It is a pleasure meet you, {{.User.Name}}",
		"Hello {{.User.Name}}, I'm {{.Bot.Name}}",
	}

	responseTemplates.RegisterHandlerResponses(setSpeakerNameHandlerName, templates)
}

func HandleSetSpeakerName(
	ctx context.Context,
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(setBotNameHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
