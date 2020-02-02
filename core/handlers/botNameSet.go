package handlers

import (
	"context"
	"text/template"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/parser"
)

const botNameSetHandlerNameOk = "botNameSetOk"
const botNameSetHandlerNameNotOk = "botNameSetNotOk"

func init() {
	okTemplates := []string{
		"{{.Bot.Name}}, I like it!",
		"Perfect!",
		"Then I'm {{.Bot.Name}}",
	}

	responseTemplates.RegisterHandlerResponses(botNameSetHandlerNameOk, okTemplates)

	notOkTemplates := []string{
		"I don't think it's a legit name",
	}

	responseTemplates.RegisterHandlerResponses(botNameSetHandlerNameNotOk, notOkTemplates)
}

func HandleBotNameSet(
	ctx context.Context,
	sentence parser.ParsedSentence,
) string {
	user := GetUserFromContext(ctx)
	user.FlowState.Event(flow.BotNameSet.String())

	var tmpl *template.Template
	if len(sentence.Tokens) > 1 {
		tmpl = responseTemplates.GetRandomResponse(botNameSetHandlerNameNotOk)
	} else {
		tmpl = responseTemplates.GetRandomResponse(botNameSetHandlerNameOk)
	}

	return execTemplateWithContext(ctx, tmpl)
}
