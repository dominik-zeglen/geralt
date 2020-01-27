package main

import "github.com/dominik-zeglen/geralt/api"

func main() {
	// r := repl.Repl{}
	// r.Init()
	// r.Start()

	api := api.API{}
	api.Init()
	api.Start()
}
