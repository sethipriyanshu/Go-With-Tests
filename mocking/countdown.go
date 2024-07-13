package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalword = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type spySleeper struct {
	Calls int
}

func (s *spySleeper) Sleep() {
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleep Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleep.Sleep()
	}
	fmt.Fprint(out, finalword)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
