package errcursor

import (
	"testing"

	"github.com/tdakkota/cursor/testutil"

	"github.com/stretchr/testify/require"
	"github.com/tdakkota/cursor"
)

func create(buf []byte) (cur *cursor.Cursor, errcur *Cursor) {
	cur = cursor.NewCursor(buf)
	errcur = NewCursor(cur)

	return
}

func TestNewCursor(t *testing.T) {
	cur, errcur := create(nil)

	require.Equal(t, cur, errcur.Cursor())
	require.Nil(t, errcur.Error())
}

func testData() (testutil.TestStruct, []byte) {
	return testutil.Data()
}

func Append(t testutil.TestStruct, c *Cursor) (err error) {
	c.WriteUint(t.Uint)
	c.WriteByte(t.Byte)
	c.WriteUint16(t.Uint16)
	c.WriteUint32(t.Uint32)
	c.WriteUint64(t.Uint64)
	c.WriteInt(t.Int)
	c.WriteInt8(t.Int8)
	c.WriteInt16(t.Int16)
	c.WriteInt32(t.Int32)
	c.WriteInt64(t.Int64)
	c.WriteFloat32(t.Float32)
	c.WriteFloat64(t.Float64)
	_ = c.WriteBytes(t.Bytes)
	_ = c.WriteString(t.String)
	c.WriteBool(t.Bool)

	return c.Error()
}

func Scan(t *testutil.TestStruct, c *Cursor) (err error) {
	t.Uint = c.ReadUint()
	t.Byte = c.ReadByte()
	t.Uint16 = c.ReadUint16()
	t.Uint32 = c.ReadUint32()
	t.Uint64 = c.ReadUint64()
	t.Int = c.ReadInt()
	t.Int8 = c.ReadInt8()
	t.Int16 = c.ReadInt16()
	t.Int32 = c.ReadInt32()
	t.Int64 = c.ReadInt64()
	t.Float32 = c.ReadFloat32()
	t.Float64 = c.ReadFloat64()
	t.Bytes = c.ReadBytes()
	t.String = c.ReadString()
	t.Bool = c.ReadBool()

	return c.Error()
}

func TestMarshaling(t *testing.T) {
	s, data := testData()

	t.Run("marshal", func(t *testing.T) {
		cur := NewCursor(cursor.NewCursor(nil))

		err := Append(s, cur)
		require.NoError(t, err)
		require.Equal(t, data, cur.Cursor().Buffer())
	})

	t.Run("unmarshal", func(t *testing.T) {
		cur := NewCursor(cursor.NewCursor(data))

		require.Equal(t, len(data), cur.Cursor().Len())

		s2 := testutil.TestStruct{}
		err := Scan(&s2, cur)
		require.NoError(t, err)
		require.Equal(t, s, s2)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		cur := NewCursor(cursor.NewCursor(nil))
		err := Append(s, cur)
		require.NoError(t, err)
		require.Equal(t, data, cur.Cursor().Buffer())

		cur.Cursor().Move(0)
		require.Zero(t, cur.Cursor().Index())
		s2 := testutil.TestStruct{}
		err = Scan(&s2, cur)
		require.NoError(t, err)
		require.Equal(t, s, s2)
	})
}
