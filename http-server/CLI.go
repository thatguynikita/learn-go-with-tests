package poker

import (
	"bufio"
	"io"
	"strings"
)

// CLI implements store manipulation via cli
type CLI struct {
	store PlayerStore
	in    *bufio.Scanner
}

// NewCLI instantiates CLI struct
func NewCLI(store PlayerStore, in io.Reader) *CLI {
	return &CLI{
		store: store,
		in:    bufio.NewScanner(in),
	}
}

// PlayPoker plays poker
func (cli *CLI) PlayPoker() {
	userInput := cli.readLine()
	cli.store.RecordWin(extractWinner(userInput))
}

func extractWinner(userInput string) string {
	return strings.Replace(userInput, " wins", "", 1)
}

func (cli *CLI) readLine() string {
	cli.in.Scan()
	return cli.in.Text()
}
