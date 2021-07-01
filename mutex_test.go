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

type BankAccount struct {
	// RWMutex untuk melakukan lock dan unlocking tetapi lebih fokus ke bagian read dan bagian write
	// karena mutex biasa tidak mengakomodir 2 lock write dan read (lebih ribet)
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int)  {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance",account.GetBalance())
}

