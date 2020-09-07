package cursor

// AppendSize reallocates new buffer if len(b)+length > cap(b)
// Otherwise returns b.
func AppendSize(b []byte, length int) (r []byte) {
	if cap(b)-len(b) < length {
		r = make([]byte, len(b)+length)
		copy(r, b)
		return r
	}

	return b
}
