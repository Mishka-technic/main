package main

import (
	"fmt"
	"runtime"
	"strings"
	"time"
)

const (
	iterationsNum = 4
)

func main() {
	for i := 0; i < runtime.NumCPU(); i++ {
		go func(th int) {
			for j := 0; j < iterationsNum; j++ {
				fmt.Printf(formatWork(th, j))
				time.Sleep(time.Millisecond)
			}
		}(i)
	}
	// fmt.Scanln()
}

func formatWork(in, j int) string {
	return fmt.Sprintln(strings.Repeat("  ", in), "█",
		strings.Repeat("  ", runtime.NumCPU()-in),
		"th", in,
		"iter", j, strings.Repeat("■", j))
}
