package game

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/rs/zerolog/log"
)

// Game define the game rules and controle the game execution
type Game struct {
	questions questions
	limit     int
}

// New creates a new game
func New(ql Loader, limit int) *Game {

	qq := ql.load()

	return &Game{
		questions: qq,
		limit:     limit,
	}
}

// Start the game
func (g *Game) Start() {

	timedout := make(chan struct{})
	score := make(chan int)
	finished := make(chan struct{})

	go g.setTimer(timedout)
	go g.play(score, finished)

	var finalScore int

	for {
		select {
		case s := <-score:
			finalScore = s
		case <-timedout:
			fmt.Println()
			fmt.Println("The game has ended due to a timeout")
			g.showScore(finalScore)
			return
		case <-finished:
			g.showScore(finalScore)
			return
		}
	}
}

// setTimer sets an interval to end the game due to timeout
func (g *Game) setTimer(timedout chan struct{}) {
	time.Sleep(time.Duration(g.limit) * time.Second)
	timedout <- struct{}{}
}

func (g *Game) play(score chan int, finished chan struct{}) {
	correct := 0
	r := bufio.NewReader(os.Stdin)

	current := g.questions.head
	for current != nil {
		fmt.Printf("%s = ", current.question.description)
		a, err := r.ReadString('\n')
		if err != nil {
			log.Error().Err(err).Msg("Can not read the answer! Please, try it again.")
		}

		a = a[:len(a)-1]
		if a == current.question.answer {
			correct++
			score <- correct
		}

		current = current.next
	}

	finished <- struct{}{}
}

func (g *Game) showScore(correct int) {
	fmt.Printf("You made %d out of %d\n", correct, g.questions.length)
	if correct > 0 {
		fmt.Println("Congratulations!")
	} else {
		fmt.Println("Better luck next time!")
	}
}
