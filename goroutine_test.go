package belajar_golang_goroutine

import (
	"fmt"
	"strconv"
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

func DisplayNumber(i int)  {
	fmt.Println(i)
}

func TestManyGoroutine(t *testing.T)  {
	for i:=0; i< 10000;i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}

func TestCreateChannel(t *testing.T)  {
	// buat channel
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Angga"
		fmt.Println("Selesai Mengirim data")
	}()

	data := <- channel
	fmt.Println(data)

	time.Sleep(5* time.Second)
}

func TestChannelAsParameter(t *testing.T)  {
	channel := make(chan string)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)
	close(channel)
}

func GiveMeResponse(channel chan string)  {
	time.Sleep(2 * time.Second)
	channel <- "Angga Wijaya"
}

func OnlyIn(channel chan<- string)  {
	time.Sleep(2 * time.Second)
	channel <- "Angga"
}

func OnlyOut(channel <-chan string)  {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T)  {
	channel := make(chan string)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(3 * time.Second)
	close(channel)
}

func TestBufferedChannel(t *testing.T) {
	// membuat bufferedchannel
	// bufferedchannel mengubah jumlah data yang default 1 menjadi n
	channel := make(chan string,3)
	defer close(channel)

	channel <- "angga"
	channel <- "test"

	// mengambil data dari channel
	fmt.Println(<- channel)
	fmt.Println(<- channel)
	//fmt.Println(<- channel)

	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	// memasukkan data dan menampilkan data melalui loop
	channel := make(chan string)
	
	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke-"+strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel{
		fmt.Println("Receive Data",data)
	}

	fmt.Println("Selesai")
}

