package handlers

import (
	"context"

	"github.com/dominik-zeglen/geralt/models"
	"github.com/looplab/fsm"
)

type ContextKey string

const BotContextKey = ContextKey("bot")
const UserContextKey = ContextKey("user")

func GetBotFromContext(ctx context.Context) models.Bot {
	return ctx.Value(BotContextKey).(models.Bot)
}

type User struct {
	Data      models.User
	FlowState *fsm.FSM
}

func GetUserFromContext(ctx context.Context) *User {
	return ctx.Value(UserContextKey).(*User)
}
