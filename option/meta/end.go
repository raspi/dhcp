package meta

import (
	OPT "../base"
	"bytes"
	"io"
)

type EndOption struct {
	End byte
}

func (o EndOption) String() string {
	return "<EOF>"
}

func (o EndOption) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o EndOption) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {

	base.V = EndOption{End: 0xff}
	return base, nil
}
