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

func TestPrependToBeginning(t *testing.T) {
  list := New()
  fmt.Println("Size - ", list.Size())
  list.AppendToEnd(2)
  list.AppendToEnd(3)
  list.AppendToEnd(5)
  list.PrependToBeginning(6)
  fmt.Println("Size - ", list.Size())
  current := list.Head()
  for current != nil {
    fmt.Println("Element: ", current.data)
    current = current.next
  }
}

func TestInsertAfter(t *testing.T) {
  list := New()
  fmt.Println("Size - ", list.Size())
  list.AppendToEnd(2)
  list.AppendToEnd(3)
  list.AppendToEnd(5)
  list.PrependToBeginning(6)
  list.InsertAfter(10,3)
  fmt.Println("Size - ", list.Size())
  current := list.Head()
  for current != nil {
    fmt.Println("Element: ", current.data)
    current = current.next
  }
}

func TestDeleteFromEnd(t *testing.T) {
  list := New()
  fmt.Println("Size - ", list.Size())
  list.AppendToEnd(2)
  list.AppendToEnd(3)
  list.AppendToEnd(5)
  list.AppendToEnd(6)
  list.DeleteFromEnd()
  list.AppendToEnd(7)
  fmt.Println("Size - ", list.Size())
  current := list.Head()
  for current != nil {
    fmt.Println("Element: ", current.data)
    current = current.next
  }
}

func TestDeleteFromBeginning(t *testing.T) {
  list := New()
  fmt.Println("Size - ", list.Size())
  list.AppendToEnd(2)
  list.AppendToEnd(3)
  list.AppendToEnd(5)
  list.AppendToEnd(6)
  list.DeleteFromBeginning()
  fmt.Println("Size - ", list.Size())
  current := list.Head()
  for current != nil {
    fmt.Println("Element: ", current.data)
    current = current.next
  }
}

func TestDeleteElement(t *testing.T) {
  list := New()
  fmt.Println("Size - ", list.Size())
  list.AppendToEnd(5)
  list.AppendToEnd(6)
  list.AppendToEnd(7)
  list.AppendToEnd(8)
  list.DeleteElement(7)
  fmt.Println("Size - ", list.Size())
  current := list.Head()
  for current != nil {
    fmt.Println("Element: ", current.data)
    current = current.next
  }
}
