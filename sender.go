package main

import (
	"fmt"
	"sync"
)

func sender(id int, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 0; i < 5; i++ {
		fmt.Printf("Goroutine %d sending %d\n", id, i)
		ch <- i
		//time.Sleep(time.Millisecond * 500) // Simulate work
	}
}

func receiver(ch <-chan int, agg_ch chan<- int, done chan<- bool, mu *sync.Mutex) {
	mu.Lock()
	for i := range ch {
		fmt.Printf("Receiver received: %d\n", i)
		fmt.Printf("Sending to aggregate: %d\n", i)
		agg_ch <- i
	}
	mu.Unlock()
	done <- true
}

func agg_receiver(ch <-chan int, done chan<- bool) {
	for i := range ch {
		fmt.Printf("Aggregate receiver received: %d\n", i)
	}
	done <- true
}

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	ch3 := make(chan int)
	agg_ch := make(chan int)
	done := make(chan bool)
	agg_done := make(chan bool)
	var wg sync.WaitGroup
	var mu sync.Mutex

	// Start multiple sender goroutines
	//for i := 1; i <= 3; i++ {
	wg.Add(1)
	go sender(1, ch1, &wg)
	wg.Add(1)
	go sender(2, ch2, &wg)
	wg.Add(1)
	go sender(3, ch3, &wg)
	//}

	// Start the receiver goroutine
	go receiver(ch1, agg_ch, done, &mu)
	go receiver(ch2, agg_ch, done, &mu)
	go receiver(ch3, agg_ch, done, &mu)

	// Wait for all sender goroutines to finish
	wg.Wait()
	close(ch1) // Close the channel so the receiver konws to stop
	close(ch2)
	close(ch3)

	// Wait for the receiver to finish
	<-done

	go agg_receiver(agg_ch, agg_done)

	<-agg_done
	fmt.Println("All goroutines finished")
}
