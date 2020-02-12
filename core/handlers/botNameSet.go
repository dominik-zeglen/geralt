package handlers

import (
	"context"
	"text/template"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/models"
	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	user := GetUserFromContext(ctx)
	user.FlowState.Event(flow.BotNameSet.String())

	var tmpl *template.Template
	if len(sentence.Tokens) > 1 {
		tmpl = responseTemplates.GetRandomResponse(botNameSetHandlerNameNotOk)
	} else {
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

		ctx = context.WithValue(ctx, BotContextKey, bot)
		tmpl = responseTemplates.GetRandomResponse(botNameSetHandlerNameOk)
	}

	return execTemplateWithContext(ctx, tmpl)
}
