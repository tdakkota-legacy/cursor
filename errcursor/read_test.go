package errcursor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
