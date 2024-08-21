/*

# Basic Goroutines and Channels

Write a program that starts two goroutines. The first goroutine should generate a sequence of numbers (e.g., 1 to 10) and send them to a channel. The second goroutine should receive these numbers from the channel, square them, and print the result.

*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	buffer_size := 500
	ch := make(chan int, buffer_size)

	fmt.Println("Starting program")

	wg.Add(3)

	go func() {
		defer wg.Done()
		fmt.Println("Go func 1")
		for i := 0; i < buffer_size; i++ {
			num := rand.Intn(99)
			fmt.Println("Appending to channel", num)
			ch <- num
		}
		close(ch)
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Go func 2")
		for num := range ch {
			sq := num * num
			fmt.Println("Received number", num, " and squared it to ", sq)
		}
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Go func 3")
	}()

	wg.Wait()

	fmt.Println("Ending program")

}
