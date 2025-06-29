package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Harsh", "")
		want := "Hello Harsh"
		if got != want {
			assertCorrectMessage(t, got, want)
		}
	})
	t.Run("Blank name passed", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello World"
		assertCorrectMessage(t, got, want)
	})
	t.Run("In spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola Elodie"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}
}
