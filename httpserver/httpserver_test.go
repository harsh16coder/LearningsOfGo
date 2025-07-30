package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

type SubServerRequester struct {
	scores map[string]string
	wins   []string
}

func (s *SubServerRequester) PlayerScore(name string) string {
	score := s.scores[name]
	return score
}

func (s *SubServerRequester) RecordWins(name string) {
	s.wins = append(s.wins, name)
}

func TestServer(t *testing.T) {
	store := SubServerRequester{
		map[string]string{
			"adam": "20",
			"joe":  "50",
		}, nil,
	}
	server := &Player{&store}
	t.Run("adam server", func(t *testing.T) {
		request := assertRequest(t, "adam")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusOK)
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
		assertStatus(t, response.Code, http.StatusOK)
		got := response.Body.String()
		want := "50"
		if got != want {
			t.Errorf("got this %v want this %v", got, want)
		}
	})

	t.Run("fenil", func(t *testing.T) {
		request := assertRequest(t, "fenil")
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		assertStatus(t, response.Code, http.StatusNotFound)
	})
}

// server_test.go
func TestStoreWins(t *testing.T) {
	store := SubServerRequester{
		map[string]string{}, nil,
	}
	server := &Player{&store}

	t.Run("it returns accepted on POST", func(t *testing.T) {
		request := assertPostRequest(t, "harsh")
		response := httptest.NewRecorder()

		server.ServeHTTP(response, request)

		assertStatus(t, response.Code, http.StatusAccepted)
		if len(store.wins) != 1 {
			t.Errorf("got this %v want this %v", len(store.wins), 1)
		}
	})
}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("Got this %v want this %v", got, want)
	}
}

func assertRequest(t *testing.T, playername string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/players/%s", playername), nil)
	return request
}

func assertPostRequest(t *testing.T, playername string) *http.Request {
	t.Helper()
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", playername), nil)
	return request
}
