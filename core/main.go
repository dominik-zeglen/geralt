package core

import (
	"fmt"

	"github.com/dominik-zeglen/geralt/core/intents"
)

type Core struct {
	intentPredictor intents.IntentPredictor
}

func (c *Core) Init() {
	c.intentPredictor.Init()
}

func (c Core) Reply(text string) {
	intentProbs := c.intentPredictor.GetIntent(text)

	for intent, intentProb := range intentProbs {
		fmt.Printf("%s: %0.5f\n", intent, intentProb)
	}
}
