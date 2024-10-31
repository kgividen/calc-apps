package handlers

import (
	"bytes"
	"errors"
	"testing"

	"github.com/kgividen/calc-lib"
)

func assertError(t *testing.T, err error, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Errorf("expected %v, got %v, a", err, target)
	}
}

func TestHandler_WrongNUmberOfArguments(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle(nil)
	assertError(t, err, errWrongNUmberofThings)
}

func TestHandler_InvalidFirstArgument(t *testing.T) {
	handler := NewHandler(nil, nil)
	err := handler.Handle([]string{"Invalid", "3"})
	assertError(t, err, errWrongNUmberofThings)
}

func TestHandler_HappyPath(t *testing.T) {
	writer := &bytes.Buffer{}
	handler := NewHandler(writer, &calc.Addition{})
	err := handler.Handle([]string{"3", "4"})
	assertError(t, err, nil)
}

//func TestHandler_InvalidSecondArgument(t *testing.T) {
//	handler := NewHandler(nil, nil)
//	err := handler.Handle([]string{"1", "invalid"})
//	if !errors.Is(err, errInvalidArg) {
//		t.Error("wrong error")
//	}
//}

func TestHandler_OutputWriter(t *testing.T) {
	badError := errors.New("badError")
	writer := &ErringWriter{err: badError}
	handler := NewHandler(writer, nil)
	err := handler.Handle([]string{"3", "4"})
	assertError(t, err, badError)

}

type ErringWriter struct {
	err error
}

func (this *ErringWriter) Write(p []byte) (n int, err error) {
	return 0, this.err
}
