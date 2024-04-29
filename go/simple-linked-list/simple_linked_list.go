package singlelinkedlist

import (
	"errors"
	"slices"
)

type List struct {
	head *Element
	size int
}

type Element struct {
	next  *Element
	value int
}

func New(elements []int) *List {
	list := &List{}

	for _, element := range elements {
		list.Push(element)
	}

	return list
}

func (l *List) Size() int {
	return l.size
}

func (l *List) Push(v int) {
	l.head = &Element{value: v, next: l.head}
	l.size++
}

func (l *List) Pop() (int, error) {
	if l.size == 0 {
		return 0, errors.New("list is empty")
	}

	// Get the value from the last node in the list
	value := l.head.value
	l.head = l.head.next
	l.size--

	return value, nil
}

// Array generates an array of integers from the list's elements
func (l *List) Array() []int {
	values := make([]int, l.size)

	current := l.head
	for i := 0; i < l.size; i++ {
		values[l.size-i-1] = current.value
		current = current.next
	}

	return values
}

// Reverse returns a new list with elements in opposite order
func (l *List) Reverse() *List {
	list := &List{}

	values := l.Array()

	// Reverse the array of values
	slices.Reverse(values)

	// Push the reversed values onto the new list
	for _, value := range values {
		list.Push(value)
	}

	return list
}
