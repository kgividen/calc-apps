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
	writer     io.Writer
	calculator Calculator
}

func NewHandler(writer io.Writer, calculator Calculator) *Handler {
	return &Handler{
		writer:     writer,
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

	if this.calculator == nil {
		return errUnsupportedOperation
	}

	result := this.calculator.Calculate(a, b)
	_, err = fmt.Fprint(this.writer, result)

	if err != nil {
		return err
	}

	return nil
}

var errUnsupportedOperation = errors.New("Unsupported operation.")
var errWrongNumberofThings = errors.New("Wrong number of things.")
var errInvalidArg = errors.New("Invalid Arg.")
