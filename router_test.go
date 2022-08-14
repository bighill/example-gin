package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestRoutes(t *testing.T) {
	t.Run("get root route", func(t *testing.T) {
		response := getResponse("/")
		assertStatus(t, response.Code, 200)
	})

	t.Run("get ping route", func(t *testing.T) {
		response := getResponse("/ping")
		assertStatus(t, response.Code, 200)
		assert.Equal(t, "pong", response.Body.String())
	})

	t.Run("get form route", func(t *testing.T) {
		key, value := "apple_message", "test_value"
		response := getResponse("/form?apple_field=" + value)

		got := response.Body.String()
		want := `{"` + key + `":"` + value + `"}`

		if got != want {
			t.Errorf("did not get expected response, got %s, want %s", got, want)
		}

		assertStatus(t, response.Code, 200)
	})
}

func getResponse(path string) *httptest.ResponseRecorder {
	router := initRouter()
	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", path, nil)
	router.ServeHTTP(response, request)
	return response
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("did not get correct status, got %d, want %d", got, want)
	}
}
