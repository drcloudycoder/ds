package linkedlist

import (
  "fmt"
)


// Represents structure of single item of list
type Node struct {
  data int
  next *Node
}

// Represents structure of the List
type List struct {
  head *Node
  size int
}

/*
  Create an empty list
*/
func New() *List {
  return &List{size:0}
}

/*
  Get the list's current length
*/
func (list *List) Size() int {
  return list.size
}

/*
  Get the head of the list
*/
func (list *List) Head() *Node {
  return list.head
}

/*
  Append an element to the end of list
  Time Complexity: O(n)
*/
func (list *List) AppendToEnd(data int) {
  // 1. Create a new Node
  newNode := &Node{data: data, next: nil}

  // 2a. If list contains no elements, set new node as head of list
  // 2b. If list contains any element, traverse till last and append new node
  if list.size == 0 {
    list.head = newNode
  } else if list.size > 0 {
    current := list.head
    for current.next != nil {
      current = current.next
    }
    current.next = newNode
  }

  // 3. Increment the list size
  list.size++
}

/*
  Prepend an element to the beginning of the list
  Time Complexity: O(1)
*/
func (list *List) PrependToBeginning(data int) {
  // 1. Create a new node
  newNode := &Node{data: data, next: nil}

  // 2. Point the new node's next to current head
  newNode.next = list.Head()

  // 3. Update the list's head as the new node
  list.head = newNode

  // 4. Increment the list size
  list.size++
}

/*
  Insert an element after a specific element in the list
*/
func (list *List) InsertAfter(data int, reference int) {
  // 1. Create a new node
  newNode := &Node{data: data, next: nil}

  // 2. Add node to list if the list is empty and return
  if list.Size() == 0 {
      list.head = newNode
      list.size++
      return
  }

  // 3. Get the head of the list as current iterator
  current := list.Head()

  // 4a. Traverse the list to find reference node
  // 4b. If reference node is found, insert the new node and return
  for current != nil {
    if current.data == reference {
      newNode.next = current.next
      current.next = newNode
      list.size++
      return
    }
    current = current.next
  }

  // 5. Provide message to user if the reference node was not found
  fmt.Println("Could not insert the node as did not find the given reference node")
}

/*
  Delete an element from the end of the list
  Time Complexity: O(n)
*/
func (list *List) DeleteFromEnd() {
  // 1. Provide message to user if the list is empty and return
  if list.Size() == 0 {
    fmt.Println("Nothing to delete, the list is empty")
    return
  }

  // 2. Get the head of the list as current iterator
  current := list.Head()

  // 3. Traverse the list until the second last element is reached
  for current.next.next != nil {
    current = current.next
  }

  // 4. Update next pointer of second last element such that last element is removed
  current.next = nil

  // 5. Decrement the list size
  list.size--

}

/*
  Delete an element from beginning of the list
  Time Complexity: O(1)
*/
func (list *List) DeleteFromBeginning() {
  // 1. Provide message to user if the list is empty and return
  if list.Size() == 0 {
    fmt.Println("Nothing to delete, the list is empty")
    return
  }

  // 2. Get the current head and save it in temp
  oldHead := list.Head()

  // 3. Update the list's head to next element in list
  list.head = oldHead.next

  // 4. Remove the link from the old head
  oldHead.next = nil

  // 5. Decrement the list size
  list.size--
}

/*
  Delete a specific element from the list
  Time Complexity: O(n)
*/
func (list *List) DeleteElement(data int) {
  // 1. Provide message to user if the list is empty and return
  if list.Size() == 0 {
    fmt.Println("Nothing to delete, the list is empty")
    return
  }

  // 2. Get the current head of the list
  current := list.Head()

  // 3. Update the head if current head is the requested element and return
  if current.data == data {
    list.head = current.next
    current.next = nil
    list.size--
    return
  }

  // 4. Traverse the list, remove the requested element and return
  for current.next != nil {
      if current.next.data == data {
        tmp := current.next.next
        current.next.next = nil
        current.next = tmp
        list.size--
        return
      }
      current = current.next
  }

  // 5. Provide a message to user if the requested element is not found in list
  fmt.Println("Could not delete since the element requested does not exist in list")
}

/*
  Find position of an element
  Time Complexity: O(n)
*/
func (list *List) FindPositionofElement(data int) int {
  // 1. Get the current head of list
  current := list.Head()

  // 2. Set position integer to 1
  position := 1

  // 3. Loop over the list until requested element is found and keep track of position
  for current != nil {
    if current.data == data {
      return position
    }
    current = current.next
    position++
  }

  // 4. Return -1 if the requested element is not found
  return -1

}
