package core

import (
	"context"
	"fmt"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/core/intents"
	"github.com/dominik-zeglen/geralt/parser"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.mongodb.org/mongo-driver/mongo"
)

const intentThreshold = .4

type Core struct {
	db              *mongo.Database
	intentPredictor intents.IntentPredictor
}

func (c *Core) Init(db *mongo.Database) {
	c.db = db
	c.intentPredictor.Init()
}

func (c Core) Reply(ctx context.Context, text string) string {
	span, spanCtx := opentracing.StartSpanFromContext(
		ctx,
		"core-reply",
	)
	defer span.Finish()

	var handler handlers.ReplyHandler

	parsedText := parser.Transform(spanCtx, text)

	predictSpan := opentracing.StartSpan(
		"core-predict-intent",
		opentracing.ChildOf(span.Context()),
	)
	intentProbs := c.intentPredictor.GetIntent(parsedText)
	predictSpan.Finish()

	user := handlers.GetUserFromContext(spanCtx)
	intent, intentProb := intentProbs.Max()
	if intentProb > intentThreshold && intent == intents.Back {
		handler = handlers.BackHandler
	} else {
		if user.FlowState.Current() != flow.Default.String() {
			switch user.FlowState.Current() {
			case flow.SettingBotName.String():
				handler = handlers.BotNameSetHandler
			case flow.SettingSpeakerName.String():
				handler = handlers.SpeakerNameSetHandler
			}
		} else {
			if intentProb > intentThreshold {
				switch intent {
				case intents.Hello:
					handler = handlers.HelloHandler
					break
				case intents.Status:
					handler = handlers.StatusHandler
					break

				case intents.GetSpeakerName:
					handler = handlers.GetSpeakerNameHandler
					break

				case intents.SetSpeakerName:
					handler = handlers.SetSpeakerNameHandler
					break

				case intents.GetBotName:
					handler = handlers.GetBotNameHandler
					break

				case intents.SetBotName:
					handler = handlers.SetBotNameHandler
					break

				default:
					handler = handlers.UnknownHandler
				}
			} else {
				handler = handlers.UnknownHandler
			}
		}
	}

	for intent, intentProb := range intentProbs {
		fmt.Printf("%s: %0.5f\n", intent, intentProb)
	}

	handlerSpan, handlerSpanCtx := opentracing.StartSpanFromContext(
		spanCtx,
		"handler",
	)
	handlerSpan.LogFields(
		log.String("handler-name", handler.GetName()),
	)
	response := handler.Exec(handlerSpanCtx, c.db, parsedText)
	handlerSpan.Finish()

	return response
}
