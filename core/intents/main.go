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

func (prediction IntentPrediction) Max() (Intent, float64) {
	var maxP float64
	var maxI Intent

	for intent, probability := range prediction {
		if probability > maxP {
			maxP = probability
			maxI = intent
		}
	}

	return maxI, maxP
}

func init() {
	rand.Seed(0)
}
