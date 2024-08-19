package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func rand_list(id int, wg *sync.WaitGroup) {
	defer wg.Done()

	max := rand.Intn(29) + 1

	fmt.Printf("[thread %d]: %d items\n", id, max)
	for i := 1; i <= max; i++ {
		fmt.Printf("[thread %d]: %d\n", id, i)
	}
}

func main() {
	var wg sync.WaitGroup

	threads := 8

	wg.Add(threads)
	for i := 1; i <= threads; i++ {
		go rand_list(i, &wg)
	}
	wg.Wait()

	fmt.Println("Done")
}
