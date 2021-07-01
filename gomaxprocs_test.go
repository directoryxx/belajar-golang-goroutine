package belajar_golang_goroutine

import (
	"fmt"
	"runtime"
	"testing"
)

func TestGoMaxProcs(t *testing.T) {
	totalCpu := runtime.NumCPU()
	fmt.Println("Total CPU",totalCpu)

	totalThreads := runtime.GOMAXPROCS(-1)
	fmt.Println("Total Threads",totalThreads)

	totalGoRoutine := runtime.NumGoroutine()
	fmt.Println("TOtal goroutine",totalGoRoutine)
}
