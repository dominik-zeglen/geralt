package handlers

import (
	"context"
	"text/template"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/models"
	"github.com/dominik-zeglen/geralt/parser"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

const botNameSetHandlerName = "botNameSet"

var BotNameSetHandler ReplyHandler

func init() {
	okTemplates := []string{
		"{{.Bot.Name}}, I like it!",
		"Perfect!",
		"Then I'm {{.Bot.Name}}",
	}

	responseTemplates.RegisterHandlerResponses(
		botNameSetHandlerName+handlerOkSuffix,
		okTemplates,
	)

	notOkTemplates := []string{
		"I don't think it's a legit name",
	}

	responseTemplates.RegisterHandlerResponses(
		botNameSetHandlerName+handlerNotOkSuffix,
		notOkTemplates,
	)

	BotNameSetHandler = createReplyHandler(
		botNameSetHandlerName,
		handleBotNameSet,
	)
}

func handleBotNameSet(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	user := GetUserFromContext(ctx)
	user.FlowState.Event(flow.BotNameSet.String())

	var tmpl *template.Template
	if len(sentence.Tokens) > 1 {
		tmpl = responseTemplates.GetRandomResponse(
			botNameSetHandlerName + handlerNotOkSuffix,
		)
	} else {
		dbSpan, _ := opentracing.StartSpanFromContext(
			ctx,
			"db-call",
		)
		globals := db.Collection(models.GlobalsCollectionKey)
		bot := GetBotFromContext(ctx)
		bot.Name = sentence.Text
		r, updateErr := globals.UpdateOne(
			context.TODO(),
			bson.M{
				"_id": bot.ID,
			}, bson.M{
				"$set": bson.M{
					"name": bot.Name,
				},
			},
		)

		if updateErr != nil || r.MatchedCount == 0 {
			panic(updateErr)
		}
		dbSpan.Finish()

		ctx = context.WithValue(ctx, BotContextKey, bot)
		tmpl = responseTemplates.GetRandomResponse(
			botNameSetHandlerName + handlerOkSuffix,
		)
	}

	return execTemplateWithContext(ctx, tmpl)
}
