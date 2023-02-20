package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var x int32 = 0
	group := sync.WaitGroup{}

	for i := 0; i < 1000; i++ {
		group.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				// x = x + 1
				atomic.AddInt32(&x, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Nilai Akhir", x)
}
