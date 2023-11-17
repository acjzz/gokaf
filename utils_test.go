package gokaf

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Increment", func(t *testing.T) {
		counter := &Counter{}
		counter.Increment()

		expected := 1
		if counter.Value() != expected {
			t.Errorf("Expected counter value %d, but got %d", expected, counter.Value())
		}
	})

	t.Run("Decrement", func(t *testing.T) {
		counter := &Counter{}
		counter.Decrement()

		expected := -1
		if counter.Value() != expected {
			t.Errorf("Expected counter value %d, but got %d", expected, counter.Value())
		}
	})

	t.Run("ConcurrentIncrementDecrement", func(t *testing.T) {
		counter := &Counter{}
		var wg sync.WaitGroup
		numIterations := 1000

		// Simulate concurrent increment and decrement
		for i := 0; i < numIterations; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				counter.Increment()
				counter.Decrement()
			}()
		}

		// Wait for all goroutines to finish
		wg.Wait()

		expected := 0
		if counter.Value() != expected {
			t.Errorf("Expected counter value %d, but got %d", expected, counter.Value())
		}
	})
}
