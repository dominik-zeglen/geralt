package main

import (
	"github.com/dominik-zeglen/geralt/api"
	"github.com/dominik-zeglen/geralt/tracing"
	"github.com/opentracing/opentracing-go"
)

func main() {
	tracer, closer := tracing.InitJaeger()
	opentracing.SetGlobalTracer(tracer)
	defer closer.Close()

	api := api.API{}
	api.Init()

	api.Start()
}
