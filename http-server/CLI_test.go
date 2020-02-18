package poker_test

import (
	"strings"
	"testing"

	poker "github.com/thatguynikita/learn-go-with-tests/http-server"
)

func TestCLI(t *testing.T) {
	t.Run("record nikita win from user input", func(t *testing.T) {
		in := strings.NewReader("Nikita wins")
		playerStore := &poker.StubPlayerStore{}
		poker.NewCLI(playerStore, in)
		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWins(t, playerStore, "Nikita")
	})

	t.Run("record cleo win from user input", func(t *testing.T) {
		in := strings.NewReader("Cleo wins")
		playerStore := &poker.StubPlayerStore{}

		cli := poker.NewCLI(playerStore, in)
		cli.PlayPoker()

		poker.AssertPlayerWins(t, playerStore, "Cleo")
	})
}
