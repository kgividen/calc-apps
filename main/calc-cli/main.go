package main

import (
	"log"
	"os"

	"github.com/kgividen/calc-apps/handlers"
	"github.com/kgividen/calc-lib"
)

func main() {
	handler := handlers.NewHandler(os.Stdout, &calc.Addition{})
	err := handler.Handle(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

}
