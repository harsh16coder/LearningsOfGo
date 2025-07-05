package main

import (
	"fmt"
	"io"
	"time"
)

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

const write = "write"
const sleep = "sleep"

type SpyCountDownOperation struct {
	Calls []string
}

func (s *SpyCountDownOperation) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpyCountDownOperation) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

func Countdown(out io.Writer, sleeper *SpyCountDownOperation) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, "Go!")
}

func main() {
	// Countdown(os.Stdout,)
}
