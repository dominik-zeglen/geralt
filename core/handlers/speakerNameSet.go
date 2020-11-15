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

const speakerNameSetHandlerName = "speakerNameSet"

var SpeakerNameSetHandler ReplyHandler

func init() {
	okTemplates := []string{
		"Ok, {{.User.Name}}",
		"It's very nice to meet you, {{.User.Name}}",
	}

	responseTemplates.RegisterHandlerResponses(
		speakerNameSetHandlerName+handlerOkSuffix,
		okTemplates,
	)

	notOkTemplates := []string{
		"I don't think it's a legit name",
	}

	responseTemplates.RegisterHandlerResponses(
		speakerNameSetHandlerName+handlerNotOkSuffix,
		notOkTemplates,
	)

	SpeakerNameSetHandler = createReplyHandler(
		speakerNameSetHandlerName,
		handleSpeakerNameSet,
	)
}

func handleSpeakerNameSet(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	var tmpl *template.Template

	if len(sentence.Tokens) > 1 {
		tmpl = responseTemplates.GetRandomResponse(
			speakerNameSetHandlerName + handlerNotOkSuffix,
		)
	} else {
		user := GetUserFromContext(ctx)
		user.FlowState.Event(flow.SpeakerNameSet.String())

		dbSpan, _ := opentracing.StartSpanFromContext(
			ctx,
			"db-call",
		)
		users := db.Collection(models.UsersCollectionKey)
		user.Data.Name = sentence.Text
		r, updateErr := users.UpdateOne(
			context.TODO(),
			bson.M{
				"_id": user.Data.ID,
			},
			bson.M{
				"$set": user.Data,
			},
		)

		if updateErr != nil || r.MatchedCount == 0 {
			panic(updateErr)
		}
		dbSpan.Finish()

		ctx = context.WithValue(ctx, UserContextKey, user)
		tmpl = responseTemplates.GetRandomResponse(
			speakerNameSetHandlerName + handlerOkSuffix,
		)
	}

	return execTemplateWithContext(ctx, tmpl)
}
