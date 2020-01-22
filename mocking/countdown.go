package main

import (
	"time"
	"io"
	"os"
	"fmt"
)

const (
	finalWord = "Go!"
	countdownStart = 3
)

// Sleeper interface to support all sorts of sleepers (primarily for mocking)
type Sleeper interface {
    Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep	func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

// Countdown counts from 3 to 1 per line
func Countdown(w io.Writer, s Sleeper) {
	for i := countdownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(w, i)
	}
	s.Sleep()
	fmt.Fprint(w, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}