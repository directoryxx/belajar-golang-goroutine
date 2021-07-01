package belajar_golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	// Ketika pool diambil datanya maka datanya akan hilang
	// maka jika data masih diperlukan harus dimasukkan kembali
	pool := sync.Pool{
		// prevent ketika data dipool kosong
		New: func() interface{}{
			return "new"
		},
	}

	pool.Put("angga")
	pool.Put("w")

	for i := 0; i < 10; i++ {
		go func() {
			// mengambil data
			data := pool.Get()
			// menampilkan data
			fmt.Println(data)
			// memasukkan kembali
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
}
