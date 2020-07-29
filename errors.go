package cursor

import (
	"errors"
)

var ErrStringTooLong = errors.New("string too long")
var ErrInvalidLength = errors.New("invalid length")
var ErrInvalidBits = errors.New("invalid bits size")
var ErrUnknownIntSize = errors.New("unknown int size")
