package errcursor

import (
	"testing"

	"github.com/stretchr/testify/require"
)

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
