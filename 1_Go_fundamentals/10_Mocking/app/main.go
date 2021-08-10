package main

import (
	"os"
	"time"

	. "github.com/qinchenfeng/HelloLearnGoWithTest/1_Go_fundamentals/10_Mocking"
)

func main() {
	sleeper := &ConfigurableSleeper{Duration: 1 * time.Second, SleepFunc: time.Sleep}
	Countdown(os.Stdout, sleeper)
}
