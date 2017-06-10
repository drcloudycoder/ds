package stack

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

var top *Node

func New(data int) *Node {
	return &Node{data: data}
}

func IsEmpty() bool {
	return top == nil
}

func Peek() int {
	return top.data
}

func Push(data int) {
	new := New(data)
	new.next = top
	top = new
}

func Pop() int {
	if top == nil {
		fmt.Println("There is nothing to pop. Stack is empty.")
		return -1
	}
	data := top.data
	top = top.next
	return data
}
