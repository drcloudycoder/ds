package stack

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	fmt.Println("---------- Testing Stack ----------")
	Push(1)
	Push(2)
	Push(3)
	fmt.Println("Peeking: ", Peek())
	for !IsEmpty() {
		fmt.Println(Pop())
	}
	Pop()
}
