package core

import (
	"context"
	"fmt"

	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/core/intents"
	"github.com/dominik-zeglen/geralt/core/middleware"
	"github.com/dominik-zeglen/geralt/parser"
)

type Core struct {
	middlewares     []middleware.Middleware
	intentPredictor intents.IntentPredictor
}

func (c *Core) Init() {
	c.middlewares = []middleware.Middleware{
		middleware.WithBot,
		middleware.WithUser,
	}
	c.intentPredictor.Init()
}

func (c Core) handleReply(ctx context.Context, text string) string {
	parsedText := parser.Transform(ctx, text)
	intentProbs := c.intentPredictor.GetIntent(parsedText)

	for intent, intentProb := range intentProbs {
		fmt.Printf("%s: %0.5f\n", intent, intentProb)
	}

	var handler handlers.ReplyHandler
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

	return handler(ctx, parsedText)
}

func (c Core) Reply(ctx context.Context, text string) string {
	for _, withMiddleware := range c.middlewares {
		ctx = withMiddleware(ctx)
	}
	return c.handleReply(ctx, text)
}
