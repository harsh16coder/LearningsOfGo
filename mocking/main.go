package main

import (
	"fmt"
	"io"
	"os"
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

func Countdown(out io.Writer, sleeper *configurable) {
	for i := 3; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprintf(out, "Go!")
}

// configurable sleeper

type configurable struct {
	duration         time.Duration
	sleep            func(time.Duration)
	totalTimeElapsed int
}

type mySleeper struct {
	timeslept time.Duration
}

func (s *mySleeper) sleep(duration time.Duration) {
	s.timeslept = duration
}

func (s *configurable) Sleep() {
	s.totalTimeElapsed += int(s.duration.Seconds())
	s.sleep(s.duration)
}

func main() {
	sleeper := &configurable{1 * time.Second, time.Sleep, 0}
	Countdown(os.Stdout, sleeper)
	fmt.Println("time slept %v seconds", sleeper.totalTimeElapsed)
}
