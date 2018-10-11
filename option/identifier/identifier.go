package identifier

import (
	OPT "../base"
	"bytes"
	"fmt"
	"io"
)

type Option struct {
	Identifier []byte
}

func (o Option) String() string {
	return fmt.Sprintf(`%+v`, o.Identifier)
}

func (o Option) Write(w io.Writer) error {
	n, err := w.Write(o.Identifier)
	if err != nil {
		return err
	}

	l := len(o.Identifier)
	if n != l {
		return fmt.Errorf(`length mismatch: %v != %v`, l, n)
	}

	return nil
}

func (o Option) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {
	b := make([]byte, r.Size())
	l, err := r.Read(b)
	if err != nil {
		return base, err
	}

	if int64(l) != r.Size() {
		return base, fmt.Errorf(`%v != %v`, l, r.Size())
	}

	base.V = Option{Identifier: b}
	return base, nil
}
