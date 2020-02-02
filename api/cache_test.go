package api

import (
	"testing"
	"time"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/patrickmn/go-cache"
)

func TestFlowPersistance(t *testing.T) {
	userID := "test"
	user := handlers.User{}
	user.FlowState = flow.NewFlow()

	api := API{}
	api.users = cache.New(2*time.Minute, 4*time.Minute)

	api.users.Set(userID, &user, cache.NoExpiration)

	userFromCtx := api.getUser(userID)
	if userFromCtx.FlowState.Current() != flow.Default.String() {
		t.Errorf(
			"Mismatching states, expected: %s, got: %s",
			flow.Default.String(),
			userFromCtx.FlowState.Current(),
		)
	}

	userFromCtx.FlowState.Event(flow.BotNameSet.String())

	userFromCtxAfterFlowChange := api.getUser(userID)

	if userFromCtx.FlowState.Current() != userFromCtxAfterFlowChange.FlowState.Current() {
		t.Error("Flow does not persist between requests")
	}
}
