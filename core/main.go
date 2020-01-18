package core

import (
	"fmt"

	"github.com/dominik-zeglen/geralt/core/intents"
	"github.com/dominik-zeglen/geralt/parser"
)

type Core struct {
	intentPredictor intents.IntentPredictor
}

func (c *Core) Init() {
	c.intentPredictor.Init()
}

func (c Core) Reply(text string) {
	parsedSentence := parser.Transform(text)
	intentProbs := c.intentPredictor.GetIntent(parsedSentence)

	for intent, intentProb := range intentProbs {
		fmt.Printf("%s: %0.5f\n", intent, intentProb)
	}
}
