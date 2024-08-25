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

	go func() {
		for i := range ticker.C {
			select {
			case ch <- struct{}{}:
				fmt.Println("Token sent at", i)
			default:
				fmt.Println("Rate limit exceeded")
			}
		}
	}()

	go func() {
		for i := 0; i < 100; i++ {
			fmt.Printf("Attempting operation %d at %v\n", i, time.Now())
			<-ch
			fmt.Printf("Completed operation %d at %v\n", i, time.Now())
		}
	}()

	time.Sleep(30 * time.Second)
}
