package main

import (
	"fmt"
	"testing"
	"timecounter/timescale"
)

const LIMIT int64 = 1_000_000_000

func loop(limitIteration int64) {
	var index int64 = 0
	for index < limitIteration {
		index++
	}
}

func TestTimeCounter(t *testing.T) {
	fmt.Println("Package -> main")

	TimeCounter(timescale.MILLISECONDS)

	Start()
	loop(LIMIT)
	End()

	PrintTime()
}