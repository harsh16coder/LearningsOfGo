package main

import (
	"log"
	"net/http"
)

type InMemoryPlayerStore struct{}

func (i *InMemoryPlayerStore) PlayerScore(name string) string {
	return "123"
}

func (i *InMemoryPlayerStore) RecordWins(name string) {
}

func main() {
	server := &Player{&InMemoryPlayerStore{}}
	log.Fatal(http.ListenAndServe(":5000", server))
}
