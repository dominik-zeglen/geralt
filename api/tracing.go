package api

import (
	"net/http"

	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
)

func (api *API) withTracing(
	next http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		span := opentracing.GlobalTracer().StartSpan("api-handler")
		defer span.Finish()
		ctx := opentracing.ContextWithSpan(r.Context(), span)
		span.LogFields(
			log.String("url", r.URL.Path),
		)

		next(w, r.WithContext(ctx))
	}
}
