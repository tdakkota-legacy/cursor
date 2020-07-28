package cursor

import (
	"encoding/binary"
)

type Cursor struct {
	cursor int
	buf    []byte
	order  binary.ByteOrder
}

func NewCursor(buf []byte) *Cursor {
	return &Cursor{buf: buf, order: binary.LittleEndian}
}

func (c *Cursor) Order(order binary.ByteOrder) {
	c.order = order
}

func (c *Cursor) PreAllocate(length int) {
	c.buf = AppendSize(c.buf, length)
}

func (c *Cursor) Reset() {
	c.Move(0)
}

func (c *Cursor) Move(index int) {
	c.cursor = index
}

func (c *Cursor) Index() int {
	return c.cursor
}

func (c *Cursor) Len() int {
	return len(c.buf)
}

func (c *Cursor) Buffer() []byte {
	return c.buf
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
	c.PreAllocate(length)
}

func (c *Cursor) should(length int) error {
	if c.cursor+length > len(c.buf) {
		return ErrInvalidLength
	}

	return nil
}
