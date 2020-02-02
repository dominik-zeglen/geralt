package api

import (
	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/dominik-zeglen/geralt/core/handlers"
	"github.com/dominik-zeglen/geralt/models"
	"github.com/patrickmn/go-cache"
)

func (api *API) rememberUser(user models.User) *handlers.User {
	userData := &handlers.User{
		Data:      user,
		FlowState: flow.NewFlow(),
	}

	api.users.Set(user.ID.Hex(), userData, cache.DefaultExpiration)

	return userData
}

func (api *API) getUser(id string) *handlers.User {
	u, ok := api.users.Get(id)
	if !ok {
		return nil
	}

	user := u.(*handlers.User)

	return user
}
