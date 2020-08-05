package errcursor

import "github.com/tdakkota/cursor"

func (c *Cursor) WriteUint(b uint) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteUint(b)

}

func (c *Cursor) WriteByte(b byte) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteByte(b)

}

func (c *Cursor) WriteUint8(b uint8) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteUint8(b)

}

func (c *Cursor) WriteUint16(b uint16) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteUint16(b)

}

func (c *Cursor) WriteUint32(b uint32) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteUint32(b)

}

func (c *Cursor) WriteUint64(b uint64) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteUint64(b)

}

func (c *Cursor) WriteInt(b int) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteInt(b)

}

func (c *Cursor) WriteInt8(b int8) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteInt8(b)

}

func (c *Cursor) WriteInt16(b int16) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteInt16(b)

}

func (c *Cursor) WriteInt32(b int32) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteInt32(b)
}

func (c *Cursor) WriteInt64(b int64) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteInt64(b)
}

func (c *Cursor) WriteFloat32(b float32) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteFloat32(b)
}

func (c *Cursor) WriteFloat64(b float64) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteFloat64(b)
}

func (c *Cursor) WriteBool(b bool) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteBool(b)
}

func (c *Cursor) WriteBytesBits(b []byte, bits int64) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteBytesBits(b, bits)
}

func (c *Cursor) WriteStringBits(b string, bits int64) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteStringBits(b, bits)
}

func (c *Cursor) WriteBytes(b []byte) (r int) {
	if c.err != nil {
		return 0
	}

	r, c.err = c.cur.WriteBytes(b)
	return
}

func (c *Cursor) WriteString(b string) (r int) {
	if c.err != nil {
		return
	}

	r, c.err = c.cur.WriteString(b)
	return
}

func (c *Cursor) Append(b []byte) {
	if c.err != nil {
		return
	}

	c.err = c.cur.Append(b)
}

func (c *Cursor) WriteCursor(cur *cursor.Cursor) {
	if c.err != nil {
		return
	}

	c.err = c.cur.WriteCursor(cur)
}
