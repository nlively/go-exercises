package main

import (
	"fmt"

	"github.com/nlively/go-exercises/data_structures/linked_lists/common"
)

func main() {
	list := &common.LinkedList{}

	list.Dump()

	list.Prepend(3)
	list.Prepend(18)
	seventythree := list.Prepend(73)
	nine := list.Prepend(9)

	list.Append(4)

	list.InsertAfter(23, seventythree)

	list.InsertBefore(1, nine)

	fmt.Println("Expected: 1, 9, 73, 23, 18, 3, 4")

	list.Dump()

	list.DeleteValue(73)
	list.DeleteValue(9999)

	list.Dump()
	list.DumpValues()

	list.Reverse()
	list.DumpValues()
}
