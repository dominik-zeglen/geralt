package parser

import (
	"context"
	"strings"

	"github.com/caneroj1/stemmer"
	"gopkg.in/jdkato/prose.v2"
)

type ParsedToken struct {
	Value string
	Tag   string
}
type ParsedSentence struct {
	Text   string
	Tokens []ParsedToken
}

func init() {
	prose.WithExtraction(false)
	prose.WithSegmentation(false)
}

func Transform(ctx context.Context, text string) ParsedSentence {
	doc, _ := prose.NewDocument(strings.ToUpper(text))

	tokens := make([]ParsedToken, len(doc.Tokens()))
	for tokenIndex, token := range doc.Tokens() {
		tokens[tokenIndex] = ParsedToken{
			Value: getTokenValue(token.Text, token.Tag),
			Tag:   token.Tag,
		}
	}

	return ParsedSentence{
		Text:   text,
		Tokens: tokens,
	}
}

func getTokenValue(tText string, tTag string) string {
	if strings.Contains(tTag, "VB") {
		return stemmer.Stem(tText)
	} else {
		return tText
	}
}
