package main

import (
	"fmt"

	"github.com/nlively/go-exercises/data_structures/linked_lists/common"
)

func detectCycle(l *common.LinkedList) *common.Node {
	fmt.Println("Running cycle detection on linked list")
	tortoise := l.First()
	hare := l.First()

	for hare != nil && hare.Next() != nil {
		tortoise = tortoise.Next()
		hare = hare.Next().Next()
		if hare == tortoise {
			// Cycle detected
			fmt.Printf("Cycle detected at value %d!\n", hare.Value())
			break
		}
	}

	if hare == nil || hare.Next() == nil {
		// Hare reached the end. No cycle detected
		fmt.Println("No cycle detected.")
		return nil
	}

	tortoise = l.First()
	for tortoise != hare {
		fmt.Printf("Cycle discovery loop. tortoise=%d, hare=%d\n", tortoise.Value(), hare.Value())
		tortoise = tortoise.Next()
		hare = hare.Next()
	}

	fmt.Printf("Cycle begins at value %d\n", hare.Value())
	return hare
}

func main() {
	ll := common.NewLinkedList(37)

	ll.Append(99)
	ll.Append(12)
	ll.Append(23)
	node1 := ll.Append(42)
	ll.Append(7)
	ll.Append(16)
	node2 := ll.Append(8)
	ll.Append(55)
	ll.Append(2)
	ll.Append(1)
	ll.Append(5)
	ll.Append(9)
	ll.Append(55) // The presence of the same number twice is intentional

	ll.Dump()

	// Detect cycle: shouldn't be one
	detectCycle(ll)

	// Create a cycle
	node2.SetNext(node1)

	// Detect cycle: should be one
	detectCycle(ll)
}
