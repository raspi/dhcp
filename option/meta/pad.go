package meta

import (
	OPT "../base"
	"bytes"
	"io"
)

type PadOption struct {
	End byte
}

func (o PadOption) String() string {
	return "<p>"
}

func (o PadOption) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o PadOption) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {

	base.V = PadOption{End: 0}
	return base, nil
}
