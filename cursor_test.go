package cursor

import (
	"encoding/binary"
	"testing"

	"github.com/stretchr/testify/require"
)

func testData() (testStruct, []byte) {
	return testStruct{
			10,
			1,
			2,
			3,
			4,
			11,
			5,
			6,
			7,
			8,
			0,
			0,
			[]byte{'x', 'y'},
			"abc",
		}, []byte{
			10, 0, 0, 0, 0, 0, 0, 0,
			1,
			2, 0,
			3, 0, 0, 0,
			4, 0, 0, 0, 0, 0, 0, 0,
			11, 0, 0, 0, 0, 0, 0, 0,
			5,
			6, 0,
			7, 0, 0, 0,
			8, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0,
			2, 'x', 'y',
			3, 'a', 'b', 'c',
		}
}

type testStruct struct {
	uint
	byte
	uint16
	uint32
	uint64
	int
	int8
	int16
	int32
	int64
	float32
	float64
	bytes []byte
	string
}

func (t testStruct) Append(c *Cursor) (err error) {
	err = c.WriteUint(t.uint)
	if err != nil {
		return err
	}

	err = c.WriteByte(t.byte)
	if err != nil {
		return err
	}

	err = c.WriteUint16(t.uint16)
	if err != nil {
		return err
	}

	err = c.WriteUint32(t.uint32)
	if err != nil {
		return err
	}

	err = c.WriteUint64(t.uint64)
	if err != nil {
		return err
	}

	err = c.WriteInt(t.int)
	if err != nil {
		return err
	}

	err = c.WriteInt8(t.int8)
	if err != nil {
		return err
	}

	err = c.WriteInt16(t.int16)
	if err != nil {
		return err
	}

	err = c.WriteInt32(t.int32)
	if err != nil {
		return err
	}

	err = c.WriteInt64(t.int64)
	if err != nil {
		return err
	}

	err = c.WriteFloat32(t.float32)
	if err != nil {
		return err
	}

	err = c.WriteFloat64(t.float64)
	if err != nil {
		return err
	}

	_, err = c.WriteBytes(t.bytes)
	if err != nil {
		return err
	}

	_, err = c.WriteString(t.string)
	if err != nil {
		return err
	}

	return nil
}

func (t *testStruct) Read(c *Cursor) (err error) {
	t.uint, err = c.ReadUint()
	if err != nil {
		return err
	}

	t.byte, err = c.ReadByte()
	if err != nil {
		return err
	}

	t.uint16, err = c.ReadUint16()
	if err != nil {
		return err
	}

	t.uint32, err = c.ReadUint32()
	if err != nil {
		return err
	}

	t.uint64, err = c.ReadUint64()
	if err != nil {
		return err
	}

	t.int, err = c.ReadInt()
	if err != nil {
		return err
	}

	t.int8, err = c.ReadInt8()
	if err != nil {
		return err
	}

	t.int16, err = c.ReadInt16()
	if err != nil {
		return err
	}

	t.int32, err = c.ReadInt32()
	if err != nil {
		return err
	}

	t.int64, err = c.ReadInt64()
	if err != nil {
		return err
	}

	t.float32, err = c.ReadFloat32()
	if err != nil {
		return err
	}

	t.float64, err = c.ReadFloat64()
	if err != nil {
		return err
	}

	t.bytes, err = c.ReadBytes()
	if err != nil {
		return err
	}

	t.string, err = c.ReadString()
	if err != nil {
		return err
	}

	return nil
}

func TestMarshaling(t *testing.T) {
	s, data := testData()

	t.Run("marshal", func(t *testing.T) {
		cur := NewCursor(nil)
		cur.Order(binary.LittleEndian)

		err := s.Append(cur)
		require.NoError(t, err)
		require.Equal(t, data, cur.Buffer())
	})

	t.Run("unmarshal", func(t *testing.T) {
		cur := NewCursor(data)

		require.Equal(t, len(data), cur.Len())

		s2 := testStruct{}
		err := s2.Read(cur)
		require.NoError(t, err)
		require.Equal(t, s, s2)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		cur := NewCursor(nil)
		cur.Reset()
		err := s.Append(cur)
		require.NoError(t, err)
		require.Equal(t, data, cur.Buffer())

		cur.Move(0)
		require.Zero(t, cur.Index())
		s2 := testStruct{}
		err = s2.Read(cur)
		require.NoError(t, err)
		require.Equal(t, s, s2)
	})
}

func TestCursor(t *testing.T) {
	cursor := NewCursor(make([]byte, 10, 20))

	r := require.New(t)
	r.Len(cursor.Buffer(), 10)
	r.Equal(10, cursor.Len())
	r.Equal(20, cursor.Cap())
	cursor.LengthBitSize(16)
	r.Equal(16, cursor.defaultBitSize)

	cursor.SetBuffer(nil)
	r.Nil(cursor.buf)
}

func TestInvalidBits(t *testing.T) {
	t.Run("read", func(t *testing.T) {
		_, err := NewCursor(make([]byte, 10)).ReadStringBits(7)
		require.Error(t, err)
	})

	t.Run("write", func(t *testing.T) {
		err := NewCursor(nil).WriteStringBits("", 7)
		require.Error(t, err)
	})
}

func TestShould(t *testing.T) {
	t.Run("enough", func(t *testing.T) {
		cursor := NewCursor(make([]byte, 10))
		err := cursor.should(10)
		require.NoError(t, err)
	})

	t.Run("not-enough", func(t *testing.T) {
		cursor := NewCursor(nil)
		err := cursor.should(10)
		require.Error(t, err)
	})
}

func TestCursor_Sub(t *testing.T) {
	t.Run("ok", func(t *testing.T) {
		data := []byte{0, 1, 2, 3, 4, 5, 6}

		cur := NewCursor(data)
		cur2, ok := cur.Sub(1, 4)
		require.True(t, ok)
		require.Equal(t, []byte{1, 2, 3}, cur2.Buffer())
	})

	t.Run("ok", func(t *testing.T) {
		data := []byte{0}

		cur := NewCursor(data)
		cur2, ok := cur.Sub(0, 1)
		require.True(t, ok)
		require.Equal(t, []byte{0}, cur2.Buffer())
	})

	t.Run("not-ok", func(t *testing.T) {
		data := []byte{0}

		cur := NewCursor(data)
		_, ok := cur.Sub(1, 4)
		require.False(t, ok)
	})

	t.Run("not-ok", func(t *testing.T) {
		data := []byte{0, 1, 2}

		cur := NewCursor(data)
		_, ok := cur.Sub(1, 4)
		require.False(t, ok)
	})
}

func TestCursor_WriteCursor(t *testing.T) {
	cursor := NewCursor(nil)
	err := cursor.WriteByte('b')
	require.NoError(t, err)

	cursor2 := NewCursor(nil)
	err = cursor2.WriteByte('a')
	require.NoError(t, err)

	err = cursor2.WriteCursor(cursor)
	require.NoError(t, err)

	require.Equal(t, []byte{'a', 'b'}, cursor2.Buffer())
}
