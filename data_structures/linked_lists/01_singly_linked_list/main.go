package main

import "fmt"

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	first *Node
}

func (l *LinkedList) prepend(value int) *Node {
	node := &Node{value: value, next: l.first}
	l.first = node

	return node
}

func (l *LinkedList) findLast() *Node {
	if l.first == nil {
		return nil
	}

	item := l.first
	for {
		if item.next == nil {
			break
		}
		item = item.next
	}

	return item
}

func (l *LinkedList) append(value int) *Node {
	node := &Node{value: value, next: nil}

	last := l.findLast()
	last.next = node

	return node
}

func (l *LinkedList) insertAfter(value int, after *Node) *Node {
	node := &Node{value: value, next: after.next}

	after.next = node

	return node
}

func (l *LinkedList) insertBefore(value int, before *Node) *Node {
	node := &Node{value: value, next: before}
	item := l.first

	if before == l.first {
		l.first = node
	} else {

		for {
			if item.next == before {
				item.next = node
				break
			} else if item.next == nil {
				return nil
			}
		}
	}

	return node
}

func (l *LinkedList) dump() {
	fmt.Printf("LIST DUMP: ")

	if l.first == nil {
		fmt.Printf("[empty]\n")
		return
	}
	fmt.Printf("\n")

	for item := l.first; item != nil; item = item.next {
		if item.next == nil {
			fmt.Printf("Last item: %d\n", item.value)
		} else {
			fmt.Printf("item %d. next %d\n", item.value, item.next.value)
		}
	}
}

func (l *LinkedList) values() []int {
	values := make([]int, 0)
	for item := l.first; item != nil; item = item.next {
		values = append(values, item.value)
	}
	return values
}

func (l *LinkedList) deleteValue(value int) {
	fmt.Printf("Deleting value %d\n", value)

	var prev *Node

	for item := l.first; item != nil; item = item.next {
		if item.value == value {
			if prev == nil {
				l.first = item.next
			} else {
				prev.next = item.next
			}
			return
		}
		prev = item
	}

	fmt.Printf("Value %d not found\n", value)
}

func (l *LinkedList) reverse() {
	var prev *Node
	item := l.first

	for {
		next := item.next
		item.next = prev
		prev = item
		if next == nil {
			l.first = prev
			break
		}
		item = next
	}
}

func (l *LinkedList) dumpValues() {
	fmt.Printf("values(): %v\n", l.values())
}

func main() {
	list := &LinkedList{}

	list.dump()

	list.prepend(3)
	list.prepend(18)
	seventythree := list.prepend(73)
	nine := list.prepend(9)

	list.append(4)

	list.insertAfter(23, seventythree)

	list.insertBefore(1, nine)

	fmt.Println("Expected: 1, 9, 73, 23, 18, 3, 4")

	list.dump()

	list.deleteValue(73)
	list.deleteValue(9999)

	list.dump()
	list.dumpValues()

	list.reverse()
	list.dumpValues()
}
