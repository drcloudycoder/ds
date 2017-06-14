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

/*
  Peek() function returns the front/head value of the Queue
*/
func Peek() int {
	if head == nil {
		fmt.Println("Queue is empty.")
		return -1
	}
	return head.data
}

/*
  Enqueue() adds a new node at the end/tail of the Queue
  Time Complexity: O(1)
*/
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

/*
  Dequeue() removes the front/head node of the Queue
  Time Complexity: O(1)
*/
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
