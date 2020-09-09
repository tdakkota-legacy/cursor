package cursor

import (
	"math"
	"unsafe"
)

// WriteUint writes given uint to buffer.
func (c *Cursor) WriteUint(b uint) error {
	switch unsafe.Sizeof(uint(0)) {
	case 1:
		return c.WriteUint8(uint8(b))
	case 2:
		return c.WriteUint16(uint16(b))
	case 4:
		return c.WriteUint32(uint32(b))
	case 8:
		return c.WriteUint64(uint64(b))
	}

	return ErrUnknownIntSize
}

// WriteByte writes given byte to buffer.
func (c *Cursor) WriteByte(b byte) error {
	return c.WriteUint8(b)
}

// WriteUint8 writes given uint8 to buffer.
func (c *Cursor) WriteUint8(b uint8) error {
	c.need(1)

	c.buf[c.cursor] = b
	c.cursor++
	return nil
}

// WriteUint16 writes given uint16 to buffer.
func (c *Cursor) WriteUint16(b uint16) error {
	c.need(2)

	c.order.PutUint16(c.buf[c.cursor:], b)
	c.cursor += 2
	return nil
}

// WriteUint32 writes given uint32 to buffer.
func (c *Cursor) WriteUint32(b uint32) error {
	c.need(4)

	c.order.PutUint32(c.buf[c.cursor:], b)
	c.cursor += 4
	return nil
}

// WriteUint64 writes given uint64 to buffer.
func (c *Cursor) WriteUint64(b uint64) error {
	c.need(8)

	c.order.PutUint64(c.buf[c.cursor:], b)
	c.cursor += 8
	return nil
}

// WriteInt writes given int to buffer.
func (c *Cursor) WriteInt(b int) error {
	switch unsafe.Sizeof(int(0)) {
	case 1:
		return c.WriteInt8(int8(b))
	case 2:
		return c.WriteInt16(int16(b))
	case 4:
		return c.WriteInt32(int32(b))
	case 8:
		return c.WriteInt64(int64(b))
	}

	return ErrUnknownIntSize
}

// WriteInt8 writes given int8 to buffer.
func (c *Cursor) WriteInt8(b int8) error {
	return c.WriteByte(byte(b))
}

// WriteInt16 writes given int16 to buffer.
func (c *Cursor) WriteInt16(b int16) error {
	return c.WriteUint16(uint16(b))
}

// WriteInt32 writes given int32 to buffer.
func (c *Cursor) WriteInt32(b int32) error {
	return c.WriteUint32(uint32(b))
}

// WriteInt64 writes given int64 to buffer.
func (c *Cursor) WriteInt64(b int64) error {
	return c.WriteUint64(uint64(b))
}

// WriteFloat32 writes given float32 to buffer.
func (c *Cursor) WriteFloat32(b float32) error {
	return c.WriteUint32(math.Float32bits(b))
}

// WriteFloat64 writes given float64 to buffer.
func (c *Cursor) WriteFloat64(b float64) error {
	return c.WriteUint64(math.Float64bits(b))
}

// WriteBool writes given bool to buffer.
func (c *Cursor) WriteBool(b bool) error {
	if b {
		return c.WriteInt8(1)
	} else {
		return c.WriteInt8(0)
	}
}

// WriteBytes writes given slice and length to buffer.
func (c *Cursor) WriteBytes(s []byte) (int, error) {
	err := c.WriteBytesBits(s, int64(c.defaultBitSize))
	if err != nil {
		return 0, err
	}
	return len(s), nil
}

// WriteBytes writes given string and length to buffer.
func (c *Cursor) WriteString(s string) (int, error) {
	err := c.WriteStringBits(s, int64(c.defaultBitSize))
	if err != nil {
		return 0, err
	}
	return len(s), nil
}

func (c *Cursor) WriteBytesBits(s []byte, bits int64) (err error) {
	length := len(s)
	switch bits {
	case 8:
		if length > math.MaxUint8 {
			return ErrStringTooLong
		}

		err = c.WriteByte(byte(length))
	case 16:
		if length > math.MaxUint16 {
			return ErrStringTooLong
		}

		err = c.WriteUint16(uint16(length))
	case 32:
		if length > math.MaxUint32 {
			return ErrStringTooLong
		}

		err = c.WriteUint32(uint32(length))
	case 64:
		err = c.WriteUint64(uint64(length))
	default:
		return ErrInvalidBits
	}

	if err != nil {
		return err
	}
	c.need(length)

	c.cursor += copy(c.buf[c.cursor:], s)
	return nil
}

func (c *Cursor) WriteStringBits(s string, bits int64) (err error) {
	return c.WriteBytesBits(s2b(s), bits)
}

func (c *Cursor) Append(b []byte) (err error) {
	c.need(len(b))
	c.cursor += copy(c.buf[c.cursor:], b)
	return nil
}

func (c *Cursor) WriteCursor(cur *Cursor) (err error) {
	return c.Append(cur.Buffer())
}
