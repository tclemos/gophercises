package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
)

// Game define the game rules and controle the game execution
type Game struct {
	questions questions
}

// New creates a new game
func New(ql Loader) *Game {

	qq := ql.load()

	return &Game{
		questions: qq,
	}
}

// Start the game
func (g *Game) Start() {

	count := 0
	correct := 0

	r := bufio.NewReader(os.Stdin)

	current := g.questions.head
	for current != nil {
		count++
		fmt.Printf("%s = ", current.question.description)
		a, err := r.ReadString('\n')
		if err != nil {
			log.Error().Err(err).Msg("Can not read the answer! Please, try it again.")
		}

		a = a[:len(a)-1]
		if a == current.question.answer {
			correct++
		}

		current = current.next
	}

	fmt.Printf("You made %d out of %d\n", correct, count)
	if correct > 0 {
		fmt.Println("Congratulations!")
	} else {
		fmt.Println("Better luck next time!")
	}
}
