package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var locker = sync.Mutex{}
var group = sync.WaitGroup{}
var cond = sync.NewCond(&locker)

func WaitCondition(value int) {
	defer group.Done()

	// Locking dengan kondisi Wait(), menunggu lanjut atau tidak sesudah di lock
	cond.L.Lock()

	// diberi arahan untuk menunggu signal apakah langsung lanjut atau tidak
	cond.Wait()
	fmt.Println("Done ", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T) {
	for i := 0; i < 10; i++ {
		group.Add(1)
		go WaitCondition(i)
	}

	// mengirim signal 1 per 1, untuk melanjutkan cond.Wait() diatas, jika tidak mengirimkan signal, maka akan deadlock
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)

			// memberi signal lanjut setiap 1 detik setiap yang di lock
			cond.Signal()
		}
	}()

	// mengirim boradcast secara keseluruhan, untuk melanjutkan cond.Wait() diatas, jika tidak mengirimkan signal, maka akan deadlock
	// sehingga semua cond.Wait() tidak perlu menunggu signal lain, karena sudah diberi arahan untuk lanjut pada semua lock
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		time.Sleep(1 * time.Second)
	// 		cond.Broadcast()
	// 	}
	// }()

	group.Wait()
}
