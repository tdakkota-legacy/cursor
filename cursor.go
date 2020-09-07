package cursor

import (
	"encoding/binary"
)

// Cursor is a byte-slice cursor.
type Cursor struct {
	cursor int
	buf    []byte

	defaultBitSize int
	order          binary.ByteOrder
}

// NewCursor creates new Cursor.
func NewCursor(buf []byte) *Cursor {
	return &Cursor{buf: buf, order: binary.LittleEndian, defaultBitSize: 8}
}

// Sets default bit size of length word.
func (c *Cursor) LengthBitSize(i int) {
	c.defaultBitSize = i
}

// Sets byte order.
func (c *Cursor) Order(order binary.ByteOrder) {
	c.order = order
}

// Increase buffer size.
func (c *Cursor) Grow(length int) {
	c.buf = AppendSize(c.buf, length)
}

// Sets buffer size to zero, moves cursor to zero index
func (c *Cursor) Reset() {
	c.buf = c.buf[:0]
	c.cursor = 0
}

// Move sets cursor position.
// Warning: user should reallocate buffer manually.
func (c *Cursor) Move(index int) {
	c.cursor = index
}

// Index returns current position of cursor.
func (c *Cursor) Index() int {
	return c.cursor
}

// Returns length of buffer.
func (c *Cursor) Len() int {
	return len(c.buf)
}

// Returns capacity of buffer.
func (c *Cursor) Cap() int {
	return cap(c.buf)
}

// Buffer returns current buffer.
func (c *Cursor) Buffer() []byte {
	return c.buf
}

// SetBuffer sets current buffer.
func (c *Cursor) SetBuffer(buf []byte) {
	c.buf = buf
}

func (c *Cursor) checkRange(i int) bool {
	return i >= 0 && i < len(c.buf)
}

func (c *Cursor) Sub(from, to int) (*Cursor, bool) {
	if from <= to && c.checkRange(from) && c.checkRange(to-1) {
		return NewCursor(c.buf[from:to]), true
	}
	return nil, false
}

func (c *Cursor) need(length int) {
	c.Grow(length)
}

func (c *Cursor) should(length int) error {
	if c.cursor+length > len(c.buf) {
		return ErrInvalidLength
	}

	return nil
}
