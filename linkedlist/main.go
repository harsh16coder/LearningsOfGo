package main

import (
	"fmt"
)

type Node struct {
	Data int
	Next *Node
}

// instantiate
func NewNode(data int) *Node {
	return &Node{Data: data, Next: nil}
}

// linked list structure

type LinkedList struct {
	Head *Node
	Tail *Node
}

// singly linkedlist
type SinglyLinkedList struct {
	Head *Node
}

func (s *SinglyLinkedList) Append(data int) {
	node := NewNode(data)
	if s.Head == nil {
		s.Head = node
		return
	}
	last := s.Head
	for last.Next != nil {
		last = last.Next
	}
	last.Next = node

}

func (s *SinglyLinkedList) Remove(data int) {
	if s.Head == nil {
		return
	}
	if s.Head.Data == data {
		s.Head = s.Head.Next
	}
	current := s.Head
	for current.Next != nil {
		if current.Next.Data == data {
			current.Next = current.Next.Next
			return
		}
		current = current.Next
	}
}

func (s *SinglyLinkedList) Display() {
	current := s.Head
	if current == nil {
		fmt.Println("List is empty")
		return
	}

	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}

type DoublyNode struct {
	Data int
	prev *DoublyNode
	next *DoublyNode
}

func NewDoublyNode(data int) *DoublyNode {
	return &DoublyNode{
		Data: data,
		prev: nil,
		next: nil,
	}
}

type DoublyLinkedList struct {
	Head *DoublyNode
}

func (d *DoublyLinkedList) AppendAtLast(data int) {
	node := NewDoublyNode(data)
	if d.Head == nil {
		d.Head = node
		return
	}
	last := d.Head
	for last.next != nil {
		last = last.next
	}
	last.next = node
	node.prev = last
}

func (d *DoublyLinkedList) AppendAtStart(data int) {
	node := NewDoublyNode(data)
	if d.Head == nil {
		d.Head = node
		return
	}
	node.next = d.Head
	d.Head.prev = node
	d.Head = node
}

func (d *DoublyLinkedList) DeleteAtEnd() {
	if d.Head == nil {
		fmt.Println("List is empty")
	}
	if d.Head.next == nil {
		d.Head = nil
		return
	}
	last := d.Head
	for last.next != nil {
		last = last.next
	}
	last.prev.next = nil
}

func (d *DoublyLinkedList) DeleteAtStart() {
	if d.Head == nil {
		fmt.Println("list is empty")
	}
	if d.Head.next == nil {
		d.Head = nil
		return
	}
	newHead := d.Head.next
	newHead.prev = nil
	d.Head = newHead
}

func (d *DoublyLinkedList) Traverse() {
	current := d.Head
	if current == nil {
		fmt.Println("List is empty")
		return
	}

	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.next
	}
	fmt.Println("nil")
}

func main() {
	// s := SinglyLinkedList{}
	// s.Append(10)
	// s.Append(20)
	// s.Append(30)
	// s.Display()
	// s.Remove(20)
	// s.Display()
	d := DoublyLinkedList{}
	d.AppendAtLast(10)
	d.AppendAtLast(20)
	d.AppendAtLast(30)
	d.Traverse()
	d.AppendAtStart(40)
	d.Traverse()
	d.AppendAtLast(100)
	d.Traverse()
	d.DeleteAtStart()
	d.DeleteAtEnd()
	d.Traverse()

}
