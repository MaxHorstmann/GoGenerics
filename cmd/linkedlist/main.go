package main

import (
	"fmt"
)

type Node struct {
	value int
	next  *Node
}

type LinkedList struct {
	head *Node
}

func (l *LinkedList) Insert(val int) {
	newNode := Node{value: val}
	if l.head == nil {
		l.head = &newNode
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = &newNode
}

func (l *LinkedList) PrintAll() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func main() {
	var ll *LinkedList = new(LinkedList)
	ll.Insert(123)
	ll.Insert(456)
	ll.Insert(789)
	ll.PrintAll()

}
