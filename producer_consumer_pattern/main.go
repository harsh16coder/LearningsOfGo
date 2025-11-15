package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func linesProducer(filepath string, lines chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("Error occured,%v", err)
		close(lines)
		file.Close()
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file; %v", err)
	}
	for scanner.Scan() {
		lines <- scanner.Text()
	}
	close(lines)
}

func linesConsumer(id int, lines <-chan string, wg *sync.WaitGroup, processedCount *int64) {
	defer wg.Done()
	for line := range lines {
		time.Sleep(time.Millisecond * 10)
		num, err := strconv.Atoi(strings.TrimSpace(line))
		if err == nil {
			fmt.Printf("Consumer %d processed: %d Count: %d (squared: %d)\n", id, num, *processedCount, num*num)
		} else {
			fmt.Printf("Consumer %d skipped non-integer line: %s\n", id, line)
		}
		// Safely increment processed count using a mutex or atomic operation
		// For simplicity, we'll use atomic.AddInt64 in main
	}
	fmt.Printf("Consumer %d finished.\n", id)
}

func main() {
	const (
		numConsumers = 5
		bufferSize   = 100
		filepath     = "data.txt"
	)
	createDummyFile(filepath, 1000)
	linesChannel := make(chan string, bufferSize)
	var wg sync.WaitGroup
	var processedCount int64
	//start the producer
	wg.Add(1)
	go linesProducer(filepath, linesChannel, &wg)

	// start consumer
	for i := 0; i < numConsumers; i++ {
		wg.Add(1)
		go linesConsumer(i, linesChannel, &wg, &processedCount)
	}
	wg.Wait()
	fmt.Printf("All producers and consumers finished.\n")
}

// Helper function to create a dummy file
func createDummyFile(filePath string, numLines int) {
	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i := 0; i < numLines; i++ {
		fmt.Fprintf(writer, "%d\n", i)
	}
	writer.Flush()
	fmt.Printf("Created dummy file: %s with %d lines.\n", filePath, numLines)
}
