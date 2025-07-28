package generics

import "testing"

func TestStack(t *testing.T) {
	t.Run("test integers", func(t *testing.T) {
		stackOfInt := NewStack[int]()
		AssertTrue(t, stackOfInt.IsEmpty())
		stackOfInt.Push(12)
		AssertFalse(t, stackOfInt.IsEmpty())
		stackOfInt.Push(20)
		stackOfInt.Push(50)
		poppedEle, _ := stackOfInt.Pop()
		AsserEqual(t, 50, poppedEle)
	})
	t.Run("test strings", func(t *testing.T) {
		stackOfInt := NewStack[string]()
		AssertTrue(t, stackOfInt.IsEmpty())
		stackOfInt.Push("Harsh")
		AssertFalse(t, stackOfInt.IsEmpty())
		stackOfInt.Push("Prachiti")
		stackOfInt.Push("Love")
		poppedEle, _ := stackOfInt.Pop()
		AsserEqual(t, "Love", poppedEle)
	})
}

func AssertTrue(t *testing.T, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func AssertFalse(t *testing.T, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}

func AsserEqual[T comparable](t *testing.T, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got this %v want this %v ", got, want)
	}
}
