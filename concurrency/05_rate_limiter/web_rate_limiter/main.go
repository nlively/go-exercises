package main

import (
	"fmt"
	"math/rand"
	"time"
)

const MAX_PER_TICK = 5

func producer(ch chan<- string, done_ch <-chan bool) {
	for i := 0; ; i++ {
		select {
		case <-done_ch:
			fmt.Println("Finished producing requests")
			return
		default:
			ch <- fmt.Sprintf("R%d", i)
			time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		}
	}
}

func server(ch <-chan string, ticker *time.Ticker, done_ch chan<- bool) {
	processed_count := 0
	rejected_count := 0
	i := 0
	for req := range ch {
		select {
		case <-ticker.C:
			fmt.Println("[tick]")
			i = 0
		default:
			i++
		}
		if i < MAX_PER_TICK {
			fmt.Printf("Processing request %s\n", req)
			processed_count++
		} else {
			fmt.Printf("rate limit exceeded! throwing away %s\n", req)
			rejected_count++
		}
	}

	fmt.Println("------------------")
	fmt.Println("SUMMARY:")
	fmt.Printf("Processed: %d\n Rejected: %d\n", processed_count, rejected_count)
	done_ch <- true
}

func main() {
	request_pipeline := make(chan string, 20)
	done_producing := make(chan bool)
	done_processing := make(chan bool)
	ticker := time.NewTicker(1 * time.Second)

	go producer(request_pipeline, done_producing)
	go server(request_pipeline, ticker, done_processing)

	// Run the simulation for a set period of time
	time.Sleep(10 * time.Second)
	done_producing <- true
	ticker.Stop()
	close(request_pipeline)
	<-done_processing
}
