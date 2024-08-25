/*
### 5. **Rate Limiter**
   - **Task:** Implement a rate limiter using channels. The rate limiter should allow only a certain number of operations per second. You can test this by creating a goroutine that tries to perform an operation at a higher rate and observe how the rate limiter handles it.
   - **Goal:** Learn how to control the rate of operations in a concurrent program.
*/

package main

import (
	"fmt"
	"time"
)

func main() {

	ticker := time.NewTicker(time.Second / 5) // 5 operations per second
	defer ticker.Stop()

	ch := make(chan struct{}, 1)
	done := make(chan bool)

	// This is how tickers work in Go.  A goroutine responds to each "tick"
	// at the ticker's specified interval as long as the ticker is running.
	go func() {
		// The ticker regulates how fast this loop runs
		for i := range ticker.C {
			select {
			case <-done:
				fmt.Println("received DONE signal")
				return
			case ch <- struct{}{}:
				fmt.Println("Token sent at", i)
			default:
				fmt.Println("Rate limit exceeded")
			}
		}
	}()

	// A separate thread that runs up to 100 times, and whose speed is regulated
	// by the speed at which items are fed into the channel
	go func() {
		for i := 0; i < 15; i++ {
			fmt.Printf("Attempting operation %d at %v\n", i, time.Now())
			// Block until channel has data
			<-ch
			fmt.Printf("Completed operation %d at %v\n", i, time.Now())
			// Pause before iterating through the loop
			time.Sleep(1000 * time.Millisecond)
		}
		done <- true
	}()

	time.Sleep(60 * time.Second)
}
