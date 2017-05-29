package linkedlist

import (
  "fmt"
  "testing"
)

func TestAppendToEnd(t *testing.T) {
  list := New()
  fmt.Println("Size - ", list.Size())
  list.AppendToEnd(2)
  list.AppendToEnd(3)
  list.AppendToEnd(5)
  fmt.Println("Size - ", list.Size())
  current := list.Head()
  for current != nil {
    fmt.Println("Element: ", current.data)
    current = current.next
  }
}