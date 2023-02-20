package belajar_golang_goroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup, counter int) {
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello ", counter)
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T) {
	group := &sync.WaitGroup{}

	for i := 1; i <= 100; i++ {
		go RunAsynchronous(group, i)
	}

	group.Wait()
	fmt.Println("Complete")
}
