package common

import "fmt"

type Node struct {
	value int
	next  *Node
}

func (n *Node) Value() int {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) SetNext(n2 *Node) {
	n.next = n2
}

type LinkedList struct {
	first *Node
}

func NewLinkedList(firstValue int) *LinkedList {
	return &LinkedList{first: &Node{value: firstValue}}
}

func (l *LinkedList) First() *Node {
	return l.first
}

func (l *LinkedList) Prepend(value int) *Node {
	node := &Node{value: value, next: l.first}
	l.first = node

	return node
}

func (l *LinkedList) FindLast() *Node {
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

func (l *LinkedList) Append(value int) *Node {
	node := &Node{value: value, next: nil}

	last := l.FindLast()
	last.next = node

	return node
}

func (l *LinkedList) InsertAfter(value int, after *Node) *Node {
	node := &Node{value: value, next: after.next}

	after.next = node

	return node
}

func (l *LinkedList) InsertBefore(value int, before *Node) *Node {
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

func (l *LinkedList) Dump() {
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

func (l *LinkedList) Values() []int {
	values := make([]int, 0)
	for item := l.first; item != nil; item = item.next {
		values = append(values, item.value)
	}
	return values
}

func (l *LinkedList) DeleteValue(value int) {
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

func (l *LinkedList) Reverse() {
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

func (l *LinkedList) DumpValues() {
	fmt.Printf("values(): %v\n", l.Values())
}
