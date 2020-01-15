package game

import (
	"bytes"
	"encoding/csv"
	"io"
	"io/ioutil"

	"github.com/rs/zerolog/log"
)

// Loader knows how to load questions
type Loader interface {
	load() questions
}

// DefaultLoader provide some hardcoded questions
type DefaultLoader struct{}

// load some hardcoded questions
func (l *DefaultLoader) load() questions {

	qq := newQuestions()

	qq.add(question{
		description: "5+5",
		answer:      "10",
	})

	qq.add(question{
		description: "5*5",
		answer:      "25",
	})

	qq.add(question{
		description: "5-5",
		answer:      "0",
	})

	return *qq
}

// CSVLoader loads questions from a csv file
type CSVLoader struct {
	path string
}

// NewCSVLoader creates and initializes an instance of CSVLoader
func NewCSVLoader(path string) *CSVLoader {
	return &CSVLoader{
		path: path,
	}
}

// Load questions from the csv file
func (l *CSVLoader) load() questions {
	buf, err := ioutil.ReadFile(l.path)
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to load the csv file")
	}

	qq := newQuestions()

	reader := csv.NewReader(bytes.NewReader(buf))
	columns, err := reader.Read()
	for err == nil {

		qq.add(question{
			description: columns[0],
			answer:      columns[1],
		})

		columns, err = reader.Read()
	}

	if err != io.EOF {
		log.Fatal().Err(err).Msg("Failed to read the csv file content")
	}

	return *qq
}
