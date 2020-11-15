package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

const setSpeakerNameHandlerName = "setSpeakerName"

var SetSpeakerNameHandler ReplyHandler

func init() {
	templates := []string{
		"Ok, so what's your name?",
	}

	responseTemplates.RegisterHandlerResponses(
		setSpeakerNameHandlerName,
		templates,
	)

	SetSpeakerNameHandler = createReplyHandler(
		setSpeakerNameHandlerName,
		handleSetSpeakerName,
	)
}

func handleSetSpeakerName(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	user := GetUserFromContext(ctx)
	user.FlowState.Event(flow.ToSpeakerNameSetting.String())

	tmpl := responseTemplates.GetRandomResponse(setSpeakerNameHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
