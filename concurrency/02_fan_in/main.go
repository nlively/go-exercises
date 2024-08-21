/*
Basic Goroutines and Channels

Implement a "fan-in" function that takes multiple channels as input and multiplexes the messages onto a single channel. Use this function to merge the outputs of two different goroutines generating sequences of numbers.
*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func upstream_sender(ch chan<- string, stream_num int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("Starting stream", stream_num)
	max := rand.Intn(100)
	fmt.Printf("[stream %d]: %d messages\n", stream_num, max)

	for i := 0; i < max; i++ {
		fmt.Printf("[stream %d]: sending message %d\n", stream_num, i)
		ch <- fmt.Sprintf("Stream %d.%d", stream_num, i)
	}

	fmt.Println("Done with stream", stream_num)
}

func aggregator(agg_ch chan<- string, upstream_ch <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for msg := range upstream_ch {
		agg_ch <- msg
	}
}

func final_receiver(ch <-chan string, done chan<- bool) {
	for msg := range ch {
		fmt.Println("received message", msg)
	}
	done <- true
}

func main() {
	var wg sync.WaitGroup
	var agg_wg sync.WaitGroup

	ch1 := make(chan string)
	ch2 := make(chan string)
	ch3 := make(chan string)
	ch4 := make(chan string)

	agg_ch := make(chan string)
	done := make(chan bool)

	wg.Add(4) // Adjust this to match the number of upstream_senders
	go upstream_sender(ch1, 1, &wg)
	go upstream_sender(ch2, 2, &wg)
	go upstream_sender(ch3, 3, &wg)
	go upstream_sender(ch4, 4, &wg)

	// Close the channels after all upstream_senders are done
	go func() {
		wg.Wait()
		close(ch1)
		close(ch2)
		close(ch3)
		close(ch4)
	}()

	// Aggregators
	agg_wg.Add(4)
	go aggregator(agg_ch, ch1, &agg_wg)
	go aggregator(agg_ch, ch2, &agg_wg)
	go aggregator(agg_ch, ch3, &agg_wg)
	go aggregator(agg_ch, ch4, &agg_wg)

	// Close aggregator channel after all messages are processed
	go func() {
		agg_wg.Wait()
		close(agg_ch)
	}()

	// Final receiver
	go final_receiver(agg_ch, done)

	<-done // Wait for final_receiver to finish
	fmt.Println("Finished")
}
