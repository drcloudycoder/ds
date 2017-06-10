package queue

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

var head *Node // Dequeue from head
var tail *Node // Enqueue at the tail

func IsEmpty() bool {
	return head == nil
}

func Peek() int {
	if head == nil {
		fmt.Println("Queue is empty.")
		return -1
	}
	return head.data
}

func Enqueue(data int) {
	new := &Node{data: data}
	if tail != nil {
		tail.next = new
	}
	tail = new

	if head == nil {
		head = new
	}
}

func Dequeue() int {
	if head == nil {
		fmt.Println("Queue is empty.")
		return -1
	}
	data := head.data
	head = head.next
	if head == nil {
		tail = nil
	}
	return data
}
