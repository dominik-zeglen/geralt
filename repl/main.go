package repl

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/dominik-zeglen/geralt/core"
)

func Start() {
	geralt := core.Core{}
	geralt.Init()

	for true {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		ctx := context.TODO()

		fmt.Printf("%s\n", geralt.Reply(ctx, text))
	}
}
