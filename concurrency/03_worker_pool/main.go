/*
### 3. **Worker Pool**

**Task:**
Create a worker pool where a fixed number of goroutines (workers)
process tasks from a shared task queue. Each worker should process
tasks concurrently, and the main goroutine should wait for all
workers to finish before exiting.

**Goal:**
Understand how to manage a pool of workers to efficiently process tasks concurrently.
*/

/*
	Approach:

	Generate 20,000 fake names

	Define a func that generates names and feeds them into a channel.

	Spawn X instances of that function via goroutines, where X is defined as a constant.

	Create a waitgroup and add X to it.

	Stream the channel output into a file, where each line is a random name with the ID of the worker that handled it
*/

package main

import (
	"fmt"
	"math/rand"
	"os"
	"sync"

	"github.com/go-faker/faker/v4"
)

const (
	NUM_FAKE_NAMES = 200000
	NUM_WORKERS    = 7
)

func generate_names(worker_id int, ch chan<- string, num_names int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Spawning worker %d\n", worker_id)

	for i := 0; i < num_names; i++ {
		// Generate fake name and send it to channel
		name := fmt.Sprintf("[worker %d] - [iteration: %d] - %s, age %d\n", worker_id, i, faker.Name(), rand.Intn(99))
		ch <- name
	}

	fmt.Printf("Finished worker %d, which was responsible for %d names\n", worker_id, num_names)
}

func main() {
	var wg sync.WaitGroup

	ch := make(chan string, 250)

	remaining := NUM_FAKE_NAMES
	base_batch_size := NUM_FAKE_NAMES / NUM_WORKERS
	extra := NUM_FAKE_NAMES % NUM_WORKERS

	for i := 0; i < NUM_WORKERS; i++ {
		num_names := base_batch_size
		if i < extra {
			num_names++ // Distribute the extra names across workers
		}

		remaining -= num_names

		wg.Add(1)
		go generate_names(i+1, ch, num_names, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	file, err := os.OpenFile("names.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	for msg := range ch {
		_, err := file.WriteString(msg)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Printf("Finished generating %d names over %d worker(s)\n", NUM_FAKE_NAMES, NUM_WORKERS)
}
