package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	PlayerScore(name string) string
	RecordWins(name string)
}

type Player struct {
	store PlayerStore
}

func (p *Player) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case "POST":
		p.postScore(w, player)
	case "GET":
		p.showScore(w, player)
	}
}

func (p *Player) postScore(w http.ResponseWriter, player string) {
	p.store.RecordWins(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *Player) showScore(w http.ResponseWriter, player string) {
	score := p.store.PlayerScore(player)
	if score == "" {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprintf(w, "%s", p.store.PlayerScore(player))
}
