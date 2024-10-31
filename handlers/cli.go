package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/kgividen/calc-lib"
)

type Handler struct {
	stdout     io.Writer
	calculator *calc.Addition
}

func NewHandler(stdout io.Writer, calculator *calc.Addition) *Handler {
	return &Handler{
		stdout:     stdout,
		calculator: calculator,
	}
}

func (this *Handler) Handle(args []string) error {
	if len(args) != 2 {
		return errWrongNumberofThings
	}
	a, err := strconv.Atoi(args[0])
	if err != nil {
		return errInvalidArg
	}

	b, err := strconv.Atoi(args[1])

	if err != nil {
		return err
	}

	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.stdout, result)

	if err != nil {
		return err
	}

	return nil
}

var errWrongNumberofThings = errors.New("Wrong number of things.")
var errInvalidArg = errors.New("Invalid Arg.")
