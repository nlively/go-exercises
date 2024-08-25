package main

import (
	"fmt"
	"time"
)

func main() {
	tck := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case x := <-tck.C:
				fmt.Println("Tick at", x)
			default:
				fmt.Println("n/a")
			}
			// Regulate the for loop's speed.  It runs at 5x the speed of the ticker
			// Otherwise this go routine will operate as fast as the computer lets it,
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(5 * time.Second)
	tck.Stop()
	time.Sleep(2 * time.Second)
	done <- true
	fmt.Println("Ticker stopped")
}
