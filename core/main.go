package core

import (
	"context"
	"fmt"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/core/intents"
	"github.com/dominik-zeglen/geralt/parser"
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
	var handler handlers.ReplyHandler

	parsedText := parser.Transform(ctx, text)
	intentProbs := c.intentPredictor.GetIntent(parsedText)

	user := handlers.GetUserFromContext(ctx)
	intent, intentProb := intentProbs.Max()
	if intentProb > intentThreshold && intent == intents.Back {
		handler = handlers.HandleBack
	} else {
		if user.FlowState.Current() != flow.Default.String() {
			switch user.FlowState.Current() {
			case flow.SettingBotName.String():
				handler = handlers.HandleBotNameSet
			case flow.SettingSpeakerName.String():
				handler = handlers.HandleSpeakerNameSet
			}
		} else {
			if intentProb > intentThreshold {
				switch intent {
				case intents.Hello:
					handler = handlers.HandleHello
					break
				case intents.Status:
					handler = handlers.HandleStatus
					break

				case intents.GetSpeakerName:
					handler = handlers.HandleGetSpeakerName
					break

				case intents.SetSpeakerName:
					handler = handlers.HandleSetSpeakerName
					break

				case intents.GetBotName:
					handler = handlers.HandleGetBotName
					break

				case intents.SetBotName:
					handler = handlers.HandleSetBotName
					break

				default:
					handler = handlers.HandleUnknown
				}
			} else {
				handler = handlers.HandleUnknown
			}
		}
	}

	for intent, intentProb := range intentProbs {
		fmt.Printf("%s: %0.5f\n", intent, intentProb)
	}

	return handler(ctx, c.db, parsedText)
}
