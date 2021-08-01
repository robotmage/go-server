package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpperCaseHandler(t *testing.T) {
	t.Run("Test response", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "", nil)
		response := httptest.NewRecorder()
		handler(response, request)
		got := response.Body.String()
		want := responseString()

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})
}
