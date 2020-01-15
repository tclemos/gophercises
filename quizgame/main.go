package main

import (
	"flag"

	"github.com/tclemos/gophercises/quizgame/game"
)

func main() {

	csvFilePath := flag.String("csv", "./problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	//loader := &game.DefaultLoader{}
	loader := game.NewCSVLoader(*csvFilePath)
	game := game.New(loader)
	game.Start()
}
