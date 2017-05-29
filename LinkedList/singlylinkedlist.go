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
*/
func (list *List) AppendToEnd(data int) {

}

/*
  Prepend an element to the beginning of the list
*/
func (list *List) PrependToBeginning(data int) {

}

/*
  Insert an element at specific position in the list
*/
func (list *List) InsertAt(data int, position int) {

}

/*
  Delete an element from the end of the list
*/
func (list *List) DeleteFromEnd() {

}

/*
  Delete an element from beginning of the list
*/
func (list *List) DeleteFromBeginning() {

}

/*
  Delete a specific element from the list
*/
func (list *List) DeleteElement(data int) {

}

/*
  Find an element
*/
func (list *List) FindPositionofElement(data int) {

}
