package repl

import (
	"bufio"
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

		geralt.Reply(text)
		fmt.Print("\n")
	}
}
