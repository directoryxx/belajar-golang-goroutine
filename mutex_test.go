package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


// Untuk menghindari race condition menggunakan mutex
// Konsep mutex adalah lock dan unlock
// jadi ketika sudah di lock hanya satu goroutine yang diperbolehkan
// nb : lebih lambat dari goroutine biasa
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex
	for i := 1; i <= 1000; i++ {
		go func() {
			for j:= 1;j<=100;j++ {
				mutex.Lock()
				x++
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Counter = ",x)
}

