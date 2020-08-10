package cursor

type Appender interface {
	Append(c *Cursor) error
}

type Scanner interface {
	Scan(c *Cursor) error
}
