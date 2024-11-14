package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/kgividen/calc-apps/handlers"
	"github.com/kgividen/calc-lib"
)

func main() {
	var operation string
	flag.StringVar(&operation, "op", "+", "Operation type")
	flag.Parse()
	fmt.Println(operation)
	handler := handlers.NewHandler(os.Stdout, calculators[operation])
	err := handler.Handle(flag.Args())
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
