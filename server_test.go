package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGETBye(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, "/hello", nil)
	res := httptest.NewRecorder()

	byeHandler := ByeHandler{env: "test"}

	byeHandler.ServeHTTP(res, req)

	t.Run("returns status 200", func(t *testing.T) {
		result := res.Code
		expected := 200

		if result != expected {
			t.Errorf("got %q, expected %q", result, expected)
		}
	})
}
