package base

import (
	"../enums"
	"bytes"
	"io"
)

type OptionInterface interface {
	Read(r *bytes.Reader) (base OptionBase, err error)
	Write(w io.Writer) error
}

type OptionBase struct {
	Type enums.OpCode
	V    OptionInterface // Dynamic type value
}
