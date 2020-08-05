package errcursor

import "github.com/tdakkota/cursor"

type Cursor struct {
	cur *cursor.Cursor
	err error
}

func NewCursor(cur *cursor.Cursor) *Cursor {
	return &Cursor{cur: cur}
}

func (c *Cursor) Cursor() *cursor.Cursor {
	return c.cur
}

func (c *Cursor) Error() error {
	return c.err
}
