package option

import (
	"./enums"
	"bytes"
)

type OptionRaw struct {
	Type   enums.OpCode
	Length uint8
	Value  *bytes.Reader
}
