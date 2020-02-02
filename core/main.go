package core

import (
	"context"
	"fmt"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/core/intents"
	"github.com/dominik-zeglen/geralt/parser"
)

type Core struct {
	intentPredictor intents.IntentPredictor
}

func (c *Core) Init() {
	c.intentPredictor.Init()
}

func (c Core) Reply(ctx context.Context, text string) string {
	var handler handlers.ReplyHandler

	parsedText := parser.Transform(ctx, text)
	intentProbs := c.intentPredictor.GetIntent(parsedText)

	user := handlers.GetUserFromContext(ctx)
	fmt.Println(user.FlowState.Current())
	if user.FlowState.Current() != flow.Default.String() {
		switch user.FlowState.Current() {
		case flow.SettingBotName.String():
			handler = handlers.HandleBotNameSet
		}
	} else {
		intent, intentProb := intentProbs.Max()
		if intentProb > .4 {
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

	// for intent, intentProb := range intentProbs {
	// 	fmt.Printf("%s: %0.5f\n", intent, intentProb)
	// }

	return handler(ctx, parsedText)
}
