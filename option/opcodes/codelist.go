package opcodes

import (
	OPT "../base"
	"../enums"
	"bytes"
	"io"
	"strings"
)

type Option struct {
	Codes []enums.OpCode
}

func (o Option) String() string {
	var out []string

	for _, c := range o.Codes {
		out = append(out, c.String())
	}

	return strings.Join(out, `, `)

}

func (o Option) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o Option) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {

	l := Option{}

	for {
		b, err := r.ReadByte()

		if err != nil {
			if err == io.EOF {
				break
			}

			return base, err
		}

		l.Codes = append(l.Codes, enums.OpCode(b))
	}

	base.V = l
	return base, nil
}
