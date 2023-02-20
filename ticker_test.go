package belajar_golang_goroutines

import (
	"fmt"
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Second)

	// menghentikan ticker
	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

// func TestTick(t *testing.T) {
// 	channelTick := time.Tick(1 * time.Second)

// 	for time := range channelTick {
// 		fmt.Println(time)
// 	}
// }
