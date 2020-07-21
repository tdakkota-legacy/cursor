package cursor

import "math"

func (c *Cursor) ReadByte() (b byte, err error) {
	err = c.should(1)
	if err != nil {
		return
	}

	b = c.buf[c.cursor]
	c.cursor++
	return
}

func (c *Cursor) ReadUint16() (b uint16, err error) {
	err = c.should(2)
	if err != nil {
		return
	}

	b = c.order.Uint16(c.buf[c.cursor:])
	c.cursor += 2
	return
}

func (c *Cursor) ReadUint32() (b uint32, err error) {
	err = c.should(4)
	if err != nil {
		return
	}

	b = c.order.Uint32(c.buf[c.cursor:])
	c.cursor += 4
	return
}

func (c *Cursor) ReadUint64() (b uint64, err error) {
	err = c.should(8)
	if err != nil {
		return
	}

	b = c.order.Uint64(c.buf[c.cursor:])
	c.cursor += 8
	return
}

func (c *Cursor) ReadInt8() (b int8, err error) {
	r, err := c.ReadByte()
	if err != nil {
		return
	}

	return int8(r), nil
}

func (c *Cursor) ReadInt16() (b int16, err error) {
	r, err := c.ReadUint16()
	if err != nil {
		return
	}

	return int16(r), nil
}

func (c *Cursor) ReadInt32() (b int32, err error) {
	r, err := c.ReadUint32()
	if err != nil {
		return
	}

	return int32(r), nil
}

func (c *Cursor) ReadInt64() (b int64, err error) {
	r, err := c.ReadUint64()
	if err != nil {
		return
	}

	return int64(r), nil
}

func (c *Cursor) ReadFloat32() (b float32, err error) {
	r, err := c.ReadUint32()
	if err != nil {
		return
	}

	return math.Float32frombits(r), nil
}

func (c *Cursor) ReadFloat64() (b float64, err error) {
	r, err := c.ReadUint64()
	if err != nil {
		return
	}

	return math.Float64frombits(r), nil
}

func (c *Cursor) ReadBytes(bits int64) (s []byte, err error) {
	length := uint64(0)

	switch bits {
	case 8:
		l, err := c.ReadByte()
		if err != nil {
			return nil, err
		}
		length = uint64(l)
	case 16:
		l, err := c.ReadUint16()
		if err != nil {
			return nil, err
		}
		length = uint64(l)
	case 32:
		l, err := c.ReadUint32()
		if err != nil {
			return nil, err
		}
		length = uint64(l)
	case 64:
		l, err := c.ReadUint64()
		if err != nil {
			return nil, err
		}
		length = uint64(l)
	default:
		err = ErrInvalidBits
	}

	if err != nil {
		return nil, err
	}

	s = make([]byte, length)
	c.cursor += copy(s, c.buf[c.cursor:int(length)+c.cursor])
	return
}

func (c *Cursor) ReadString(bits int64) (s string, err error) {
	b, err := c.ReadBytes(bits)
	if err != nil {
		return "", err
	}

	return b2s(b), nil
}
