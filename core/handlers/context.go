package handlers

import (
	"github.com/dominik-zeglen/geralt/models"
	"github.com/looplab/fsm"
)

type ContextKey string

const BotContextKey = ContextKey("bot")
const UserContextKey = ContextKey("user")

type User struct {
	Data      models.User
	FlowState *fsm.FSM
}
