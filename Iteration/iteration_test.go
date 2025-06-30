package iteration

import "testing"

func TestIteation(t *testing.T) {
	repeated := iteration("a")
	expected := "aaaaa"
	if expected != repeated {
		t.Errorf("Expected this %q got this %q", expected, repeated)
	}
}

func BenchmarkIteartion(b *testing.B) {
	for b.Loop() {
		iterationloop("a")
	}
}
