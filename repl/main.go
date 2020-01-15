package repl

import (
	"bufio"
	"fmt"
	"os"

	"github.com/dominik-zeglen/geralt/core/intents"
)

func Start() {
	for true {
		fmt.Print("> ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')

		intents.Reply(text)
		fmt.Print("\n")
	}
}
