package main

import "github.com/dominik-zeglen/geralt/repl"

func main() {
	r := repl.Repl{}
	r.Init()
	r.Start()
}
