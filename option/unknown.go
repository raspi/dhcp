package option

import (
	"./base"
	"bytes"
	"fmt"
	"io"
)

// Unknown
type OptUnknown struct {
	Value []byte
}

func (o OptUnknown) Read(r *bytes.Reader) (base base.OptionBase, err error) {

	if r.Size() == 0 {
		base.V = OptUnknown{
			Value: []byte{},
		}

		return base, nil
	}

	raw := make([]byte, r.Size())

	l, err := r.Read(raw)

	if err != nil {
		return base, err
	}

	if int64(l) != r.Size() {
		return base, fmt.Errorf(`len mismatch %v != %v`, r.Size(), l)
	}

	base.V = OptUnknown{
		Value: raw,
	}

	return base, nil
}

func (o OptUnknown) Write(w io.Writer) error {
	//w.Write([]byte{byte(o.Code), uint8(len(o.Value))})
	w.Write(o.Value)
	return nil
}
