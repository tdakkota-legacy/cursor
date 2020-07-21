package cursor

import "math"

func (c *Cursor) WriteByte(b byte) error {
	c.need(1)

	c.buf[c.cursor] = b
	c.cursor++
	return nil
}

func (c *Cursor) WriteUint16(b uint16) error {
	c.need(2)

	c.order.PutUint16(c.buf[c.cursor:], b)
	c.cursor += 2
	return nil
}

func (c *Cursor) WriteUint32(b uint32) error {
	c.need(4)

	c.order.PutUint32(c.buf[c.cursor:], b)
	c.cursor += 4
	return nil
}

func (c *Cursor) WriteUint64(b uint64) error {
	c.need(8)

	c.order.PutUint64(c.buf[c.cursor:], b)
	c.cursor += 8
	return nil
}

func (c *Cursor) WriteInt8(b int8) error {
	return c.WriteByte(byte(b))
}

func (c *Cursor) WriteInt16(b int16) error {
	return c.WriteUint16(uint16(b))
}

func (c *Cursor) WriteInt32(b int32) error {
	return c.WriteUint32(uint32(b))
}

func (c *Cursor) WriteInt64(b int64) error {
	return c.WriteUint64(uint64(b))
}

func (c *Cursor) WriteFloat32(b float32) error {
	return c.WriteUint32(math.Float32bits(b))
}

func (c *Cursor) WriteFloat64(b float64) error {
	return c.WriteUint64(math.Float64bits(b))
}

func (c *Cursor) WriteBytes(s []byte, bits int64) (err error) {
	length := len(s)
	switch bits {
	case 8:
		if length > math.MaxUint8 {
			err = ErrStringTooLong
		}

		err = c.WriteByte(byte(length))
	case 16:
		if length > math.MaxUint16 {
			err = ErrStringTooLong
		}

		err = c.WriteUint16(uint16(length))
	case 32:
		if length > math.MaxUint32 {
			err = ErrStringTooLong
		}

		err = c.WriteUint32(uint32(length))
	case 64:
		if uint64(length) > math.MaxUint64 {
			err = ErrStringTooLong
		}
		err = c.WriteUint64(uint64(length))
	default:
		err = ErrInvalidBits
	}

	if err != nil {
		return err
	}
	c.need(length)

	c.cursor += copy(c.buf[c.cursor:], s)
	return nil
}

func (c *Cursor) WriteString(s string, bits int64) (err error) {
	return c.WriteBytes(s2b(s), bits)
}
