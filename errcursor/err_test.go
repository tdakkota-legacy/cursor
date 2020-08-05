package errcursor

import (
	"github.com/stretchr/testify/require"
	"github.com/tdakkota/cursor"
	"testing"
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

func TestRead(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		_, errcur := create(nil) // empty buffer

		b := errcur.ReadBool() // buffer is too small, errcur.Error() is not nil
		require.False(t, b)    // zero value is false
		require.Error(t, errcur.Error())

		// should nothing do here
		errcur.ReadUint()
		errcur.ReadByte()
		errcur.ReadUint8()
		errcur.ReadUint16()
		errcur.ReadUint32()
		errcur.ReadUint64()
		errcur.ReadInt()
		errcur.ReadInt8()
		errcur.ReadInt16()
		errcur.ReadInt32()
		errcur.ReadInt64()
		errcur.ReadFloat32()
		errcur.ReadFloat64()
		errcur.ReadBool()
		errcur.ReadBytesBits(8)
		errcur.ReadStringBits(8)
		errcur.ReadBytes()
		errcur.ReadString()
	})

	t.Run("no-error", func(t *testing.T) {
		_, errcur := create([]byte{1})

		b := errcur.ReadBool()
		require.True(t, b)
		require.NoError(t, errcur.Error())
	})
}

func TestWrite(t *testing.T) {
	t.Run("error", func(t *testing.T) {
		_, errcur := create(nil) // empty buffer

		errcur.WriteStringBits("", 7) // invalid bit size

		// should nothing do here
		errcur.WriteUint(0)
		errcur.WriteByte(0)
		errcur.WriteUint8(0)
		errcur.WriteUint16(0)
		errcur.WriteUint32(0)
		errcur.WriteUint64(0)
		errcur.WriteInt(0)
		errcur.WriteInt8(0)
		errcur.WriteInt16(0)
		errcur.WriteInt32(0)
		errcur.WriteInt64(0)
		errcur.WriteFloat32(0)
		errcur.WriteFloat64(0)
		errcur.WriteBool(false)
		errcur.WriteBytesBits([]byte{}, 8)
		errcur.WriteStringBits("", 8)
		errcur.WriteBytes([]byte{})
		errcur.WriteString("")

		require.Error(t, errcur.Error())
	})

	t.Run("no-error", func(t *testing.T) {
		_, errcur := create([]byte{1})

		b := errcur.ReadBool()
		require.True(t, b)
		require.NoError(t, errcur.Error())
	})
}
