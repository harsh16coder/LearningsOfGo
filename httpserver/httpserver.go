package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	PlayerScore(string) string
}

type Player struct {
	store PlayerStore
}

func (p *Player) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	fmt.Fprintf(w, "%s", p.store.PlayerScore(player))
}

func PlayerScore(name string) string {
	if name == "Pepper" {
		return "20"
	}

	if name == "Floyd" {
		return "10"
	}

	return ""
}
