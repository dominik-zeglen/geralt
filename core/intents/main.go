package intents

import (
	"math/rand"

	"github.com/goml/gobrain"
)

// Intent is a representation of speaker's intentions in the given sentence
type Intent string

type IntentPredictor struct {
	intents    []Intent
	bagOfWords map[string]int
	classifier gobrain.FeedForward
}

type IntentPrediction map[Intent]float64

func init() {
	rand.Seed(0)
}
