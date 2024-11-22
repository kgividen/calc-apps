package main

import (
	"log"
	"os"

	"github.com/kgividen/calc-apps/handlers"
	"github.com/kgividen/calc-lib"
)

func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags)
	handler := handlers.NewCSVHandler(logger, os.Stdin, os.Stdout, calculators)
	err := handler.Handle()
	if err != nil {
		log.Fatal(err)
	}
}

var calculators = map[string]handlers.Calculator{
	"+": &calc.Addition{},
	"-": &calc.Subtraction{},
	"*": &calc.Multiplication{},
	"/": &calc.Division{},
}
