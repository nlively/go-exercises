/*
### 4. **Pipeline Pattern**
  - **Task:** Create a pipeline where data flows through multiple stages, each stage represented by a goroutine. For example, the first stage reads data from a channel, the second stage processes it, and the third stage outputs the final result.
  - **Goal:** Practice structuring programs with multiple stages of processing in a pipeline.
*/
package main

import (
	"fmt"
	"math/rand"
	"strings"
	"sync"

	"github.com/go-faker/faker/v4"
)

const (
	NUM_RECORDS = 1000
)

type Person struct {
	name    string
	age     int
	address string
}

func (p *Person) Print() {
	fmt.Printf("Name: %s, Age: %d, Address: %s\n", p.name, p.age, p.address)
}

func generate(ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(ch)

	for i := 0; i < NUM_RECORDS; i++ {
		name := faker.Name()
		ch <- name
	}
}

func to_upper(in_ch <-chan string, out_ch chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out_ch)

	for name := range in_ch {
		out_ch <- strings.ToUpper(name)
	}
}

func bundle(in_ch <-chan string, out_ch chan<- Person, wg *sync.WaitGroup) {
	defer wg.Done()
	defer close(out_ch)

	for name := range in_ch {
		person := Person{name: name, age: rand.Intn(99), address: faker.GetRealAddress().City}
		out_ch <- person
	}
}

func output(ch <-chan Person, wg *sync.WaitGroup) {
	defer wg.Done()

	for person := range ch {
		person.Print()
	}
}

func main() {
	var wg sync.WaitGroup

	names_ch := make(chan string)
	upper_ch := make(chan string)
	persons_ch := make(chan Person)

	wg.Add(4)
	go generate(names_ch, &wg)
	go to_upper(names_ch, upper_ch, &wg)
	go bundle(upper_ch, persons_ch, &wg)
	go output(persons_ch, &wg)

	wg.Wait()
}
