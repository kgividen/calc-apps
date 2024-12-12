package handlers

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("got: [%v], want: [%v]", actual, expected)
	}
}

func TestHTTPServer_Add(t *testing.T) {
	assertHTTP(t, http.MethodPost, "/add?a=1&b=2", http.StatusMethodNotAllowed, "text/plain; charset=utf-8", "Method Not Allowed\n")
	assertHTTP(t, http.MethodGet, "/add?a=NaN&b=2", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "a was invalid\n")
	assertHTTP(t, http.MethodGet, "/add?a=1&b=2", http.StatusOK, "text/plain; charset=utf-8", "3")
	assertHTTP(t, http.MethodGet, "/add?a=1&b=NaN", http.StatusUnprocessableEntity, "text/plain; charset=utf-8", "b was invalid\n")
}

func assertHTTP(t *testing.T, method, target string, expectedStatus int, expectedContentType, expectedResponse string) {
	t.Run(fmt.Sprintf("%s %s", method, target), func(t *testing.T) {
		request := httptest.NewRequest(method, target, nil)
		response := httptest.NewRecorder()

		NewRouter(nil).ServeHTTP(response, request)

		assertEqual(t, expectedStatus, response.Code)
		assertEqual(t, expectedResponse, response.Body.String())
		assertEqual(t, expectedContentType, response.Header().Get("Content-Type"))

		t.Log(response.Code)
		t.Log(response.Header())
		t.Log(response.Body)
	})

}
