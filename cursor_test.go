package cursor

import (
	"encoding/binary"
	"github.com/stretchr/testify/require"
	"testing"
)

func testData() (testStruct, []byte) {
	return testStruct{
			1,
			2,
			3,
			4,
			5,
			6,
			7,
			8,
			0,
			0,
			[]byte{'x', 'y'},
			"abc",
		}, []byte{
			1,
			2, 0,
			3, 0, 0, 0,
			4, 0, 0, 0, 0, 0, 0, 0,
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
	byte
	uint16
	uint32
	uint64
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

	err = c.WriteBytes(t.bytes, 8)
	if err != nil {
		return err
	}

	err = c.WriteString(t.string, 8)
	if err != nil {
		return err
	}

	return nil
}

func (t *testStruct) Read(c *Cursor) (err error) {
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

	t.bytes, err = c.ReadBytes(8)
	if err != nil {
		return err
	}

	t.string, err = c.ReadString(8)
	if err != nil {
		return err
	}

	return nil
}

func TestCursor(t *testing.T) {
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

		s2 := testStruct{}
		err := s2.Read(cur)
		require.NoError(t, err)
		require.Equal(t, s, s2)
	})

	t.Run("marshal-unmarshal", func(t *testing.T) {
		cur := NewCursor(nil)
		err := s.Append(cur)
		require.NoError(t, err)
		require.Equal(t, data, cur.Buffer())

		cur.Reset()
		s2 := testStruct{}
		err = s2.Read(cur)
		require.NoError(t, err)
		require.Equal(t, s, s2)
	})
}
