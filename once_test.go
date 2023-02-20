package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce() {
	counter++
}

func TestOnlyOnce(t *testing.T) {
	var once sync.Once
	var group sync.WaitGroup
	var mutex sync.Mutex

	for i := 0; i < 100; i++ {
		group.Add(1)
		go func() {
			mutex.Lock()
			once.Do(OnlyOnce)
			// OnlyOnce()
			mutex.Unlock()

			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter:", counter)
}
