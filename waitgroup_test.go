package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsync(group *sync.WaitGroup,i int)  {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello",i)
	time.Sleep(1 * time.Second)
}

// WaitGroup berfungsi menunggu goroutine sampai selesai
// waitgroup pengganti sleep pada sebelumnya
func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsync(group,i)
	}

	group.Wait()
	fmt.Println("Completed")
}
