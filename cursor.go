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

func (c *Cursor) Buffer() []byte {
	return c.buf
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
