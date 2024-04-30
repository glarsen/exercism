package circular

import (
	"errors"
)

// Buffer implements a circular buffer
type Buffer struct {
	allocated int
	elements  []byte

	read  int
	write int
}

// NewBuffer creates a Buffer with the specified size.
func NewBuffer(size int) *Buffer {
	return &Buffer{elements: make([]byte, size)}
}

// ReadByte reads the oldest byte in the buffer
func (b *Buffer) ReadByte() (byte, error) {
	if b.empty() {
		return 0, errors.New("buffer is empty")
	}

	// Get the oldest element in the buffer
	element := b.elements[b.read]

	// Increment the read cursor
	b.increment(&b.read)

	// Account for the de-allocation
	b.allocated--

	return element, nil
}

// WriteByte writes a byte into the buffer
func (b *Buffer) WriteByte(c byte) error {
	if b.full() {
		return errors.New("buffer is full")
	}

	// Write the byte into the buffer
	b.elements[b.write] = c

	// Increment the write cursor
	b.increment(&b.write)

	// Account for the allocation
	b.allocated++

	return nil
}

// Overwrite forces writes a byte into the buffer
func (b *Buffer) Overwrite(c byte) {
	// If the buffer is not full, write normally
	if err := b.WriteByte(c); err == nil {
		return
	}

	// Write the byte into the buffer
	b.elements[b.write] = c

	// Increment the read & write cursors
	b.increment(&b.read)
	b.increment(&b.write)
}

// Reset clears the buffer
func (b *Buffer) Reset() {
	b.allocated, b.read, b.write = 0, 0, 0
	b.elements = make([]byte, cap(b.elements))
}

func (b *Buffer) increment(p *int) {
	*p = (*p + 1) % b.size()
}

func (b *Buffer) empty() bool {
	return b.allocated == 0
}

func (b *Buffer) full() bool {
	return b.allocated == b.size()
}

func (b *Buffer) size() int {
	return cap(b.elements)
}
