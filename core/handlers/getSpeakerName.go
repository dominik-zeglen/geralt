package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
)

const getSpeakerNameHandlerName = "getSpeakerName"

func init() {
	templates := []string{
		"You are {{.User.Name}}",
		"{{.User.Name}}",
	}

	responseTemplates.RegisterHandlerResponses(getSpeakerNameHandlerName, templates)
}

func HandleGetSpeakerName(
	ctx context.Context,
	sentence parser.ParsedSentence,
) string {
	tmpl := responseTemplates.GetRandomResponse(getSpeakerNameHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
