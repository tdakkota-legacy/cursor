package cursor

func AppendSize(b []byte, length int) (r []byte) {
	if cap(b)-len(b) < length {
		r = make([]byte, len(b)+length)
		copy(r, b)
		return r
	}

	return b
}
