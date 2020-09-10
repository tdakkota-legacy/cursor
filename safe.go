// +build appengine unsafe

package cursor

func b2s(b []byte) string {
	return string(b)
}

func s2b(s string) (b []byte) {
	return []byte(s)
}
