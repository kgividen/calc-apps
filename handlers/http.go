package handlers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/kgividen/calc-lib"
)

func NewRouter(logger *log.Logger) http.Handler {
	router := http.NewServeMux()
	//router.Handle("/add", NewHTTPHander(logger, &calc.Addition{}))
	router.Handle("GET /add", NewHTTPHander(logger, &calc.Addition{}))
	return router
}

type HTTPHandler struct {
	logger     *log.Logger
	calculator Calculator
}

func NewHTTPHander(logger *log.Logger, calculator Calculator) http.Handler {
	return &HTTPHandler{logger: logger, calculator: calculator}
}

func (this *HTTPHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	query := request.URL.Query()
	a, err := strconv.Atoi(query.Get("a"))
	if err != nil {
		http.Error(response, "a was invalid", http.StatusUnprocessableEntity)
		//response.Header().Set("Content-Type", "text/plain; charset=utf-8")
		//response.WriteHeader(http.StatusUnprocessableEntity)
		//fmt.Fprint(response, "a was invalid")
		return
	}
	b, err := strconv.Atoi(query.Get("b"))
	if err != nil {
		http.Error(response, "b was invalid", http.StatusUnprocessableEntity)
		return
	}
	c := this.calculator.Calculate(a, b)
	fmt.Fprint(response, c)
}
