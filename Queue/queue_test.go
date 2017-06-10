package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	fmt.Println("---------- Testing Queue ----------")
	Enqueue(1)
	Enqueue(3)
	Enqueue(4)
	fmt.Println("Peeking: ", Peek())

	for !IsEmpty() {
		fmt.Println(Dequeue())
	}
	Dequeue()
}
