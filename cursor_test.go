package cursor

import (
	"encoding/binary"
	"testing"

	"github.com/tdakkota/cursor/testutil"

	"github.com/stretchr/testify/require"
)

func testData() (testutil.TestStruct, []byte) {
	return testutil.Data()
}

func Append(t testutil.TestStruct, c *Cursor) (err error) {
	err = c.WriteUint(t.Uint)
	if err != nil {
		return err
	}

	err = c.WriteByte(t.Byte)
	if err != nil {
		return err
	}

	err = c.WriteUint16(t.Uint16)
	if err != nil {
		return err
	}

	err = c.WriteUint32(t.Uint32)
	if err != nil {
		return err
	}

	err = c.WriteUint64(t.Uint64)
	if err != nil {
		return err
	}

	err = c.WriteInt(t.Int)
	if err != nil {
		return err
	}

	err = c.WriteInt8(t.Int8)
	if err != nil {
		return err
	}

	err = c.WriteInt16(t.Int16)
	if err != nil {
		return err
	}

	err = c.WriteInt32(t.Int32)
	if err != nil {
		return err
	}

	err = c.WriteInt64(t.Int64)
	if err != nil {
		return err
	}

	err = c.WriteFloat32(t.Float32)
	if err != nil {
		return err
	}

	err = c.WriteFloat64(t.Float64)
	if err != nil {
		return err
	}

	_, err = c.WriteBytes(t.Bytes)
	if err != nil {
		return err
	}

	_, err = c.WriteString(t.String)
	if err != nil {
		return err
	}

	err = c.WriteBool(t.Bool)
	if err != nil {
		return err
	}

	return nil
}

func Scan(t *testutil.TestStruct, c *Cursor) (err error) {
	t.Uint, err = c.ReadUint()
	if err != nil {
		return err
	}

	t.Byte, err = c.ReadByte()
	if err != nil {
		return err
	}

	t.Uint16, err = c.ReadUint16()
	if err != nil {
		return err
	}

	t.Uint32, err = c.ReadUint32()
	if err != nil {
		return err
	}

	t.Uint64, err = c.ReadUint64()
	if err != nil {
		return err
	}

	t.Int, err = c.ReadInt()
	if err != nil {
		return err
	}

	t.Int8, err = c.ReadInt8()
	if err != nil {
		return err
	}

	t.Int16, err = c.ReadInt16()
	if err != nil {
		return err
	}

	t.Int32, err = c.ReadInt32()
	if err != nil {
		return err
	}

	t.Int64, err = c.ReadInt64()
	if err != nil {
		return err
	}

	t.Float32, err = c.ReadFloat32()
	if err != nil {
		return err
	}

	t.Float64, err = c.ReadFloat64()
	if err != nil {
		return err
	}

	t.Bytes, err = c.ReadBytes()
	if err != nil {
		return err
	}

	t.String, err = c.ReadString()
	if err != nil {
		return err
	}

	t.Bool, err = c.ReadBool()
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

		err := Append(s, cur)
		require.NoError(t, err)
		require.Equal(t, data, cur.Buffer())
	})

	t.Run("unmarshal", func(t *testing.T) {
		cur := NewCursor(data)

		require.Equal(t, len(data), cur.Len())

		s2 := testutil.TestStruct{}
		err := Scan(&s2, cur)
		require.NoError(t, err)
		require.Equal(t, s, s2)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		cur := NewCursor(nil)
		cur.Reset()
		err := Append(s, cur)
		require.NoError(t, err)
		require.Equal(t, data, cur.Buffer())

		cur.Move(0)
		require.Zero(t, cur.Index())
		s2 := testutil.TestStruct{}
		err = Scan(&s2, cur)
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
