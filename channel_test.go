package belajar_golang_goroutines

import (
	"fmt"
	"strconv"
	"sync"
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

func TestBufferedChannel(t *testing.T) {

	// membuat channel menggunakan buffer, seakan2 memasukan data ke buffer
	// 3 disini adalah kuota dari buffer, sehingga channel tersebut tidak akan menunggu consumer untuk menerima data dari pengirim selanjutnya
	// buffer seperti antrian
	channel := make(chan string, 3)
	defer close(channel)

	// channel <- "Rizky"
	// channel <- "Sam"
	// channel <- "Pratama"

	// baca channel pertama, pembacaan data sesuai urutan
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)
	// fmt.Println(<-channel)

	// Use go routine
	go func() {
		channel <- "Rizky"
		channel <- "Sam"
		channel <- "Pratama"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	time.Sleep(2 * time.Second)
	fmt.Println("Selesai")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	// Menerima data yang tidak diketahui berapa banyak yang diterima oleh si channel
	go func() {
		for i := 0; i < 100; i++ {
			channel <- "ke-" + strconv.Itoa(i)
		}

		// pastikan harus di close agar tahu batas channelnya dan agar tidak error
		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	// insert data ke channel
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	// select {
	// case data := <-channel1:
	// 	fmt.Println("data dari channel 1", data)
	// case data := <-channel2:
	// 	fmt.Println("data dari channel 1", data)
	// }

	// use looping
	// pastikan tentukan kapan looping harus berhenti, atau pastikan channel tersebut kapan kosongnya
	// karena kalau semua sudah dikonsumsi dan looping terus berjalan maka akan muncul error karena channel tidak memiliki data
	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu Data...")
		}

		if counter == 2 {
			break
		}
	}
}

func TestRaceCondition(t *testing.T) {
	x := 0

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				x = x + 1
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Nilai Akhir", x)
}

// Mutex adalah cara untuk membuat sebuah kode hanya dieksekusi oleh 1 proses saja
// sehingga ketika sedang dieksekusi, proses lain tidak bisa meng-eksekusi kode tersebut
func TestMutex(t *testing.T) {
	x := 0
	var mutex sync.Mutex

	for i := 0; i < 1000; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				// mengunci proses dibawah agar tidak diproses oleh goroutine lain
				mutex.Lock()
				x = x + 1
				// seteleh beres proses diatas, maka mutex harus di unlock, agar goroutine lain bisa memproses lagi
				mutex.Unlock()
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Nilai Akhir", x)
}

type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}

func (account *BankAccount) AddBalance(amount int) {
	account.RWMutex.Lock()
	account.Balance = account.Balance + amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
	account.RWMutex.RLock()
	balance := account.Balance
	account.RWMutex.RUnlock()
	return balance
}

func TestRWMutex(t *testing.T) {
	account := BankAccount{}

	// looping ini menjalankan proses (AddBalance, GetBalance) secara goroutine (asynchronous)
	// artinya proses tersebut dikerjakan oleh beberapa goroutine secara bersamaan
	// didalam AddBalance dan GetBalance sudah dipasang Mutex agar tidak terjadi Race condition
	for i := 0; i < 100; i++ {
		go func() {
			for j := 0; j < 100; j++ {
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}

	time.Sleep(5 * time.Second)
	fmt.Println("Total Balance:", account.GetBalance())
}

// Note: penggunakan Mutex itu disarankan ketika sebuah proses/struct dll di akses oleh beberapa goroutine sekaligus, agar tidak terjadi race condition
