package sync

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	WordCount := 1000
	c := counterstruct{}
	var wg sync.WaitGroup
	wg.Add(WordCount)
	for i := 0; i < WordCount; i++ {
		go func() {
			c.Count()
			wg.Done()
		}()
	}
	wg.Wait()
	assertError(t, &c, WordCount)
}

func assertError(t *testing.T, got *counterstruct, wordcount int) {
	if got.Value() != wordcount {
		t.Errorf("got %d want %d", got.Value(), wordcount)
	}
}
