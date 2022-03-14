package main

import (
	"fmt"
)

type List[T any] interface {
	insert(val T)
	PrintAll()
}

type Node[T any] struct {
	value T
	next  *Node[T]
}

type LinkedList[T any] struct {
	head *Node[T]
}

func (l *LinkedList[U]) insert(val U) {
	var newNode = new(Node[U])
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

func (l *LinkedList[T]) PrintAll() {
	curr := l.head
	for curr != nil {
		fmt.Println(curr.value)
		curr = curr.next
	}
}

func main() {
	var l List[int]
	l = new(LinkedList[int])
	l.insert(123)
	l.insert(456)
	l.insert(789)
	l.PrintAll()
}
