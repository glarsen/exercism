package linkedlist

import (
	"errors"
)

// Node is a node in the doubly-linked list
type Node struct {
	next, previous *Node
	Value          any
}

// Next gets the next node in a list
func (n *Node) Next() *Node {
	return n.next
}

// Prev gets the previous node in a list
func (n *Node) Prev() *Node {
	return n.previous
}

// List is a doubly-linked list.
type List struct {
	first *Node
	last  *Node
	size  int
}

// NewList creates a new linked list preserving the order
// of the provided elements (if any)
func NewList(elements ...interface{}) *List {
	list := &List{}

	for _, element := range elements {
		list.Push(element)
	}

	return list
}

// Push inserts a value at the back of the list
func (l *List) Push(v interface{}) {
	if l.size == 0 {
		l.first = &Node{Value: v}
		l.last = l.first
	} else {
		l.last.next = &Node{previous: l.last, Value: v}
		l.last = l.last.next
	}

	l.size++
}

// Pop removes the value from the back of the list
func (l *List) Pop() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("list is empty")
	}

	// Get the value from the last node in the list
	value := l.last.Value

	// Remove the last node from the list
	l.last = l.last.previous
	l.size--

	// If the list is now empty, clear the first/last pointers
	if l.size == 0 {
		l.first = nil
		l.last = nil
	} else {
		// Otherwise clear the tail's next pointer
		l.last.next = nil
	}

	return value, nil
}

// Shift removes a value from the front of the list
func (l *List) Shift() (interface{}, error) {
	if l.size == 0 {
		return nil, errors.New("list is empty")
	}

	// Get the first node in the list
	value := l.first.Value

	// Remove the first node from the list
	l.first = l.first.next
	l.size--

	// If the list is now empty, zero out the first/last pointers
	if l.size == 0 {
		l.first = nil
		l.last = nil
	} else {
		// Otherwise clear the tail's next pointer
		l.first.previous = nil
	}

	return value, nil
}

// Unshift inserts a value at the front of the list
func (l *List) Unshift(v interface{}) {
	if l.size == 0 {
		l.first = &Node{Value: v}
		l.last = l.first
	} else {
		l.first.previous = &Node{next: l.first, Value: v}
		l.first = l.first.previous
	}

	l.size++
}

// Reverse reverses the linked list in place
func (l *List) Reverse() {
	for e := l.first; e != nil; e = e.previous {
		e.next, e.previous = e.previous, e.next
	}
	l.first, l.last = l.last, l.first
}

// First returns a pointer to the first node (head)
func (l *List) First() *Node {
	return l.first
}

// Last returns a pointer to the last node (tail)
func (l *List) Last() *Node {
	return l.last
}

// Size returns the list size
func (l *List) Size() int {
	return l.size
}
