package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SubServerRequest struct {
	scores map[string]string
}

func TestServer(t *testing.T) {
	store := SubServerRequest{
		map[string]string{
			"adam": "20",
			"joe":  "50",
		},
	}
	server := &Player{&store}
	t.Run("adam server", func(t *testing.T) {
		request := assertRequest(t, "adam")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "20"
		if got != want {
			t.Errorf("got this %v want this %v", got, want)
		}
	})
	t.Run("joe", func(t *testing.T) {
		request := assertRequest(t, "joe")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		got := response.Body.String()
		want := "50"
		if got != want {
			t.Errorf("got this %v want this %v", got, want)
		}
	})
}

func assertRequest(t *testing.T, playername string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", playername), nil)
	return request
}

func (s *SubServerRequest) PlayerScore(name string) string {
	score := s.scores[name]
	return score
}
