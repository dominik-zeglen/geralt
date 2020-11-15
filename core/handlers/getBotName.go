package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/mongo"
)

const getBotNameHandlerName = "getBotName"

var GetBotNameHandler ReplyHandler

func init() {
	templates := []string{
		"I'm {{.Bot.Name}}",
		"My name is {{.Bot.Name}}",
		"I'm {{.Bot.Name}}, an autonomic bot",
	}

	responseTemplates.RegisterHandlerResponses(getBotNameHandlerName, templates)

	GetBotNameHandler = createReplyHandler(
		getBotNameHandlerName,
		handleGetBotName,
	)
}

func handleGetBotName(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	span, _ := opentracing.StartSpanFromContext(
		ctx,
		"handler-bot-name-get",
	)
	defer span.Finish()

	tmpl := responseTemplates.GetRandomResponse(getBotNameHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
