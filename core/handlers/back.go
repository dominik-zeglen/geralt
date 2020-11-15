package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/parser"
	"go.mongodb.org/mongo-driver/mongo"
)

const backHandlerName = "back"

var BackHandler ReplyHandler

func init() {
	templates := []string{
		"Ok",
		"Ok, nevermind",
		"As you wish",
	}

	responseTemplates.RegisterHandlerResponses(backHandlerName, templates)

	BackHandler = createReplyHandler(
		backHandlerName,
		handleBack,
	)
}

func handleBack(
	ctx context.Context,
	db *mongo.Database,
	sentence parser.ParsedSentence,
) string {
	user := GetUserFromContext(ctx)
	user.FlowState.SetState(flow.Default.String())
	tmpl := responseTemplates.GetRandomResponse(backHandlerName)

	return execTemplateWithContext(ctx, tmpl)
}
