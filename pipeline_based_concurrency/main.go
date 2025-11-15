package main

import (
	"fmt"
	"time"
)

// Generator stage: Produces numbers
func generate(done <-chan struct{}, nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			case out <- n:
			case <-done:
				return
			}
		}
	}()
	return out
}

// Filter stage: Filters out even numbers
func filterOdd(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			if n%2 != 0 { // Keep only odd numbers
				select {
				case out <- n:
				case <-done:
					return
				}
			}
		}
	}()
	return out
}

// Square stage: Squares numbers
func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			squared := n * n
			select {
			case out <- squared:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {
	// A done channel for graceful shutdown of all goroutines
	done := make(chan struct{})
	defer close(done) // Ensure done is closed when main exits

	// Stage 1: Generate numbers
	numbers := generate(done, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	// Stage 2: Filter out even numbers
	oddNumbers := filterOdd(done, numbers)

	// Stage 3: Square the odd numbers
	squaredOddNumbers := square(done, oddNumbers)

	// Final stage: Consume and print results
	fmt.Println("Pipeline results:")
	for result := range squaredOddNumbers {
		fmt.Printf("Result: %d\n", result)
		time.Sleep(time.Millisecond * 10) // Simulate final processing
	}

	fmt.Println("Pipeline finished.")
}
