package testutil

func Data() (TestStruct, []byte) {
	return TestStruct{
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
			true,
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
			1,
		}
}

type TestStruct struct {
	Uint    uint
	Byte    byte
	Uint16  uint16
	Uint32  uint32
	Uint64  uint64
	Int     int
	Int8    int8
	Int16   int16
	Int32   int32
	Int64   int64
	Float32 float32
	Float64 float64
	Bytes   []byte
	String  string
	Bool    bool
}
