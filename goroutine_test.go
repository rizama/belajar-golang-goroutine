package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func TestCreateGoroutine(t *testing.T) {

	// fungsi ini di running menggunakan goroutine, sehingga tidak menunggu sampai beres proses fungsi tersebut dan akan pindah ke proses lain dibawahnya
	go RunHelloWorld()

	fmt.Println("Hello Sam")

	// Sleep digunakan sebagai waktu jeda untuk menunggu proses goroutine diatas selesai
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int) {
	fmt.Println("Display ", number)
}

func TestManyGoroutine(t *testing.T) {
	for i := 0; i < 100000; i++ {
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}
