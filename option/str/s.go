package str

import (
	OPT "../base"
	"bytes"
	"fmt"
	"io"
)

type Option struct {
	Value string
}

func (o Option) String() string {
	return o.Value
}

func (o Option) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o Option) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {

	b := make([]byte, r.Size())

	n, err := r.Read(b)

	if err != nil {
		return base, err
	}

	if int64(n) != r.Size() {
		return base, fmt.Errorf(`%v != %v`, n, r.Size())
	}

	base.V = Option{Value: string(b)}
	return base, nil
}
