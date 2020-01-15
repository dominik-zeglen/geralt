package parser

import (
	"strings"

	"github.com/caneroj1/stemmer"
	"gopkg.in/jdkato/prose.v2"
)

type ParsedToken struct {
	Text string
	Stem string
	Tag  string
}
type ParsedSentence struct {
	Text   string
	Tokens []ParsedToken
}

func init() {
	prose.WithExtraction(false)
	prose.WithSegmentation(false)
}

func Transform(text string) ParsedSentence {
	doc, _ := prose.NewDocument(strings.ToLower(text))

	tokens := []ParsedToken{}

	for _, token := range doc.Tokens() {
		tokens = append(tokens, ParsedToken{
			Text: token.Text,
			Stem: stemmer.Stem(token.Text),
			Tag:  token.Tag,
		})
	}

	return ParsedSentence{
		Text:   text,
		Tokens: tokens,
	}
}
