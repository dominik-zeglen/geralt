package main

import (
	"fmt"

	"github.com/dominik-zeglen/geralt/core/flow"
	"github.com/looplab/fsm"
)

func main() {
	fmt.Println(fsm.Visualize(flow.NewFlow()))
}
