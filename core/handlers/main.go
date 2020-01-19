package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/parser"
)

type handlerName string

type ReplyHandler func(
	ctx context.Context,
	sentence parser.ParsedSentence,
) string
