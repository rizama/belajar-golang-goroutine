package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// membuat channel
	channel := make(chan string)

	// defer: apapun yang terjadi, selalu jalan kan kode dibawah
	defer close(channel)

	// menggunakan goroutine dan kirim/memasukan data ke channel
	// channel ini sebagai wadah untuk return value dari fungsi goroutine
	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Sam Pratama"
		fmt.Println("Selesai mengirim data ke channel")
	}()

	// mengambil value dari channel
	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// ------------------------------------------------------------------------------------

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Sam Pratama"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

// fungsi ini hanya untuk mengirimkan data ke channel
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Rizky Sam Pratama"
}

// fungsi ini hanya untuk mengambil data dari channel
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestChannelAsParameterWithInOut(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}
