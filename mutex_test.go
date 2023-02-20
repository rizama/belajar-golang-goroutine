package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

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

type UserBalance struct {
	Mutex   sync.Mutex
	Name    string
	Balance int
}

func (user *UserBalance) Lock() {
	user.Mutex.Lock()
}

func (user *UserBalance) Unlock() {
	user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int) {
	user.Balance = user.Balance + amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int) {
	user1.Lock()
	fmt.Println("Lock user1:", user1.Name)
	user1.Change(-amount)

	time.Sleep(1 * time.Second)

	user2.Lock()
	fmt.Println("Lock user2:", user2.Name)
	user2.Change(amount)

	time.Sleep(1 * time.Second)

	user1.Unlock()
	user2.Unlock()
}

// Kondisi dimana goroutine saling menunggu untuk mengunci proses yang sudah di Lock() oleh proses lain
// proses dibawah adalah saling transfer pada saat yang bersamaan, dimana user akan di Lock ketika melakukan transfer
// ketika user1 ingin Lock() user2, tapi disisi lain user2 sudah di Lock() oleh goroutine lain,
// maka hasilnya saling menunggu Unlock() dan akhirnya menghasilkan deadlock
func TestDeadlock(t *testing.T) {
	user1 := UserBalance{
		Name:    "Sam",
		Balance: 1000000,
	}

	user2 := UserBalance{
		Name:    "Bimbo",
		Balance: 1000000,
	}

	go Transfer(&user1, &user2, 100000)
	go Transfer(&user2, &user1, 200000)

	// time.Sleep(5 * time.Second)

	fmt.Println("User ", user1.Name, " Balance ", user1.Balance)
	fmt.Println("User ", user2.Name, " Balance ", user2.Balance)

}
