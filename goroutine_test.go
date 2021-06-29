package belajar_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(name string) {
	fmt.Println( "Hello ",name)
}

func TestCreateGoroutine(t *testing.T)  {
	go RunHelloWorld("Angga")
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}
