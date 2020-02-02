package main

import "github.com/dominik-zeglen/geralt/api"

func main() {
	api := api.API{}
	api.Init()

	api.Start()
}
