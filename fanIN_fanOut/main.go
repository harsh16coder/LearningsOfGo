package main

import (
	"fmt"
	"sync"
	"time"
)

func worker(id int, in <-chan int, out chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()
	for n := range in {
		squared := n * n
		fmt.Printf("Worker %d: processing %d -> %d\n", id, n, squared)
		time.Sleep(time.Millisecond * 50)
		out <- squared
	}
	fmt.Printf("Worker %d finished: ", id)
}

func main() {
	const (
		numWorkers = 3
		numJobs    = 20
	)
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	var wg sync.WaitGroup
	// start worker fan out
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}
	// send jobs to jobs channel
	for i := 1; i <= numJobs; i++ {
		jobs <- i
	}
	close(jobs) // No more jobs to send

	// Wait for all workers to finish their current jobs
	// This also ensures all results are sent to the 'results' channel
	wg.Wait()
	close(results) // Important: Close results channel AFTER all workers are done
	// to signal that no more results will be produced for the fan-in collector.

	// Fan-in: Collect results
	fmt.Println("\nCollecting results:")
	for r := range results {
		fmt.Printf("Collected result: %d\n", r)
	}

	fmt.Println("All done!")
}
