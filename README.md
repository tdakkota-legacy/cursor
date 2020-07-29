# Cursor

[![Go](https://github.com/tdakkota/cursor/workflows/Go/badge.svg)](https://github.com/tdakkota/cursor/actions)
[![Documentation](https://godoc.org/github.com/tdakkota/cursor?status.svg)](https://pkg.go.dev/github.com/tdakkota/cursor?tab=subdirectories)
[![codecov](https://codecov.io/gh/tdakkota/cursor/branch/master/graph/badge.svg)](https://codecov.io/gh/tdakkota/cursor)
[![license](https://img.shields.io/github/license/tdakkota/cursor.svg?maxAge=2592000)](https://github.com/tdakkota/cursor/blob/master/LICENSE)


Go byte slice cursor

## Usage 

```go
cur := cursor.NewCursor(nil)
cur.PreAllocate(1 + 2) // preallocate buffer
cur.WriteByte(1) // write one byte
cur.WriteInt16(2) // write two bytes — one int16
cur.WriteString("abc", 8) // reallocate buffer, write 1 byte for length and 3 for string
```

### Available methods
Write:
```go
func (c *Cursor) WriteUint(b uint) error
func (c *Cursor) WriteByte(b byte) error
func (c *Cursor) WriteUint16(b uint16) error
func (c *Cursor) WriteUint32(b uint32) error
func (c *Cursor) WriteUint64(b uint64) error
func (c *Cursor) WriteInt(b int) error
func (c *Cursor) WriteInt8(b int8) error
func (c *Cursor) WriteInt16(b int16) error
func (c *Cursor) WriteInt32(b int32) error
func (c *Cursor) WriteInt64(b int64) error
func (c *Cursor) WriteFloat32(b float32) error
func (c *Cursor) WriteFloat64(b float64) error
func (c *Cursor) WriteBytes(s []byte, bits int64) (err error)
func (c *Cursor) WriteString(s string, bits int64) (err error)
```


Read:
```go
func (c *Cursor) ReadUint() (b uint, err error)
func (c *Cursor) ReadByte() (b byte, err error)
func (c *Cursor) ReadUint16() (b uint16, err error)
func (c *Cursor) ReadUint32() (b uint32, err error)
func (c *Cursor) ReadUint64() (b uint64, err error)
func (c *Cursor) ReadInt() (b int, err error)
func (c *Cursor) ReadInt8() (b int8, err error)
func (c *Cursor) ReadInt16() (b int16, err error)
func (c *Cursor) ReadInt32() (b int32, err error)
func (c *Cursor) ReadInt64() (b int64, err error)
func (c *Cursor) ReadFloat32() (b float32, err error)
func (c *Cursor) ReadFloat64() (b float64, err error)
func (c *Cursor) ReadBytes(bits int64) (s []byte, err error)
func (c *Cursor) ReadString(bits int64) (s string, err error)
```

## Install
```
go get github.com/tdakkota/cursor
```