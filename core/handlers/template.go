package handlers

import (
	"bytes"
	"context"
	"math/rand"
	"text/template"

	"github.com/dominik-zeglen/geralt/core/intents"
	"github.com/dominik-zeglen/geralt/core/middleware"
	"github.com/dominik-zeglen/geralt/models"
)

type responseTemplateData struct {
	Bot  middleware.BotData
	User models.User
}

func responsesToTemplates(
	intent intents.Intent,
	responses []string,
) []string {
	templates := make([]template.Template, len(responses))
	for responseIndex, response := range responses {
		tmpl, err := template.New(string(intent)).Parse(response)
		if err != nil {
			panic(err)
		}

		templates[responseIndex] = *tmpl
	}

	return responses
}

type responseArray map[handlerName]([](*template.Template))
type handlerResponseTemplates struct {
	initialized bool
	responses   responseArray
}

var responseTemplates handlerResponseTemplates

func (t *handlerResponseTemplates) init() {
	t.responses = responseArray{}
	t.initialized = true

}

func (t *handlerResponseTemplates) RegisterHandlerResponses(
	name handlerName,
	handlerResponseTemplates []string,
) {
	if !t.initialized {
		t.init()
	}
	templates := make([]*template.Template, len(handlerResponseTemplates))
	for tmplIndex, tmpl := range handlerResponseTemplates {
		parsedTemplate, err := template.
			New(setBotNameHandlerName + string(tmplIndex)).
			Parse(tmpl)

		if err != nil {
			panic(err)
		}

		templates[tmplIndex] = parsedTemplate
	}

	t.responses[name] = templates
}

func (t handlerResponseTemplates) GetRandomResponse(
	name handlerName,
) *template.Template {
	responses := t.responses[name]

	return responses[rand.Intn(len(responses))]
}

func execTemplateWithContext(ctx context.Context, t *template.Template) string {
	bot := ctx.Value(middleware.BotContextKey).(middleware.BotData)
	user := ctx.Value(middleware.UserContextKey).(models.User)

	response := bytes.Buffer{}
	err := t.Execute(&response, responseTemplateData{
		bot,
		user,
	})

	if err != nil {
		panic(err)
	}

	return response.String()
}
