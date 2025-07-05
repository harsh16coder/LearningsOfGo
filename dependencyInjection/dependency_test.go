package dependencyinjection

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	t.Run("Greeting", func(t *testing.T) {
		buffer := bytes.Buffer{}
		Greet(&buffer, "Harsh")
		got := buffer.String()
		want := "Hello Harsh"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
