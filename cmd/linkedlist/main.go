package main

import (
	"fmt"
)

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func Insert[T any](l *LinkedList[T], val T) {
	var newNode = new(Node[T])
	newNode.value = val
	if l.head == nil {
		l.head = newNode
		return
	}
	curr := l.head
	for curr.next != nil {
		curr = curr.next
	}
	curr.next = newNode
}

func PrintAll[T any](l *LinkedList[T]) {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func main() {
	ll := new(LinkedList[string])
	Insert(ll, "123")
	Insert(ll, "456")
	Insert(ll, "789")
	PrintAll(ll)
}
