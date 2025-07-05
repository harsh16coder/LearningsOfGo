package main

import (
	"bytes"
	"reflect"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Greet", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		Countdown(buffer, &SpyCountDownOperation{})
		got := buffer.String()
		want := `3
2
1
Go!`
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &SpyCountDownOperation{}
		Countdown(spySleepPrinter, spySleepPrinter)

		want := []string{
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
			t.Errorf("wanted calls %v got %v", want, spySleepPrinter.Calls)
		}
	})
}
