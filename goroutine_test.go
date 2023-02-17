package belajargolanggoroutine

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
