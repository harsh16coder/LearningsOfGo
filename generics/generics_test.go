package generics

import "testing"

func TestGenrics(t *testing.T) {
	t.Run("compare Integers", func(t *testing.T) {
		stacklistint := NewStack[int]()
		assertTrue(t, stacklistint.IsEmpty())
		stacklistint.Push(10)
		assertFalse(t, stacklistint.IsEmpty())
		stacklistint.Push(43)
		value, _ := stacklistint.Pop()
		assertEqual(t, value, 43)
	})

	t.Run("compare strings", func(t *testing.T) {
		stacklistint := NewStack[string]()
		assertTrue(t, stacklistint.IsEmpty())
		stacklistint.Push("harsh")
		assertFalse(t, stacklistint.IsEmpty())
		stacklistint.Push("Prachiti")
		value, _ := stacklistint.Pop()
		assertEqual(t, value, "Prachiti")
	})
}

func assertTrue(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v and want true", got)
	}
}

func assertFalse(t testing.TB, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v and want false", got)
	}
}

func assertEqual[T comparable](t testing.TB, got, want T) {
	if got != want {
		t.Errorf("got %v and want %v", got, want)
	}
}
