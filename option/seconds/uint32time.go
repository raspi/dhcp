package seconds

import (
	OPT "../base"
	"bytes"
	"encoding/binary"
	"io"
	"time"
)

type Option struct {
	Time time.Duration
}

func (o Option) String() string {
	return o.Time.String()
}

func (o Option) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o Option) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {

	var secs uint32
	err = binary.Read(r, binary.BigEndian, &secs)

	if err != nil {
		return base, err
	}

	base.V = Option{Time: time.Duration(secs) * time.Second}
	return base, nil
}
