package handlers

import (
	"context"
	"math/rand"

	"github.com/dominik-zeglen/geralt/parser"
)

func HandleHello(
	ctx context.Context,
	sentence []parser.ParsedSentence,
) string {
	greetings := []string{
		"Hi",
		"Hello",
		"Oh, hi",
		"Hey",
	}

	return greetings[rand.Intn(len(greetings))]
}
