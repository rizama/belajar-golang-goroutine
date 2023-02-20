package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	var pool = sync.Pool{
		New: func() interface{} {
			return "Default"
		},
	}
	var group sync.WaitGroup

	pool.Put("rizky")
	pool.Put("sam")
	pool.Put("pratama")

	for i := 0; i < 10; i++ {
		group.Add(1)
		go func() {
			// get data dari pool
			data := pool.Get()

			// gunakan data
			fmt.Println(data)

			time.Sleep(1 * time.Second)
			// kembalikan kembali ke pool, jika tidak dikembalikan maka data yang tadi akan hilang dari pool
			pool.Put(data)

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Completed")
}
