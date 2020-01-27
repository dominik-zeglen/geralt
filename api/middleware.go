package api

import (
	"net/http"
)

type Middleware func(http.HandlerFunc) http.HandlerFunc

func combineMiddlewares(
	handler http.HandlerFunc,
	middlewares []Middleware,
) http.HandlerFunc {

	if len(middlewares) < 1 {
		return handler
	}

	wrapped := handler

	for i := len(middlewares) - 1; i >= 0; i-- {
		wrapped = middlewares[i](wrapped)
	}

	return wrapped

}
