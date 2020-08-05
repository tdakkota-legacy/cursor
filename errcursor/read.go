package errcursor

func (c *Cursor) ReadUint() (b uint) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadUint()
	return
}

func (c *Cursor) ReadByte() (b byte) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadByte()
	return
}

func (c *Cursor) ReadUint8() (b uint8) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadUint8()
	return
}

func (c *Cursor) ReadUint16() (b uint16) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadUint16()
	return
}

func (c *Cursor) ReadUint32() (b uint32) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadUint32()
	return
}

func (c *Cursor) ReadUint64() (b uint64) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadUint64()
	return
}

func (c *Cursor) ReadInt() (b int) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadInt()
	return
}

func (c *Cursor) ReadInt8() (b int8) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadInt8()
	return
}

func (c *Cursor) ReadInt16() (b int16) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadInt16()
	return
}

func (c *Cursor) ReadInt32() (b int32) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadInt32()
	return
}

func (c *Cursor) ReadInt64() (b int64) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadInt64()
	return
}

func (c *Cursor) ReadFloat32() (b float32) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadFloat32()
	return
}

func (c *Cursor) ReadFloat64() (b float64) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadFloat64()
	return
}

func (c *Cursor) ReadBool() (b bool) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadBool()
	return
}

func (c *Cursor) ReadBytesBits(bits int64) (b []byte) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadBytesBits(bits)
	return
}

func (c *Cursor) ReadStringBits(bits int64) (b string) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadStringBits(bits)
	return
}

func (c *Cursor) ReadBytes() (b []byte) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadBytes()
	return
}

func (c *Cursor) ReadString() (b string) {
	if c.err != nil {
		return
	}

	b, c.err = c.cur.ReadString()
	return
}
