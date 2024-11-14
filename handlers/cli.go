package handlers

import (
	"errors"
	"fmt"
	"io"
	"strconv"
)

type Calculator interface {
	Calculate(a, b int) int
}

type Handler struct {
	stdout     io.Writer
	calculator Calculator
}

func NewHandler(stdout io.Writer, calculator Calculator) *Handler {
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
		return errInvalidArg
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
