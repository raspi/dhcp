package ipv4

import (
	OPT "../base"
	"bytes"
	"fmt"
	"io"
	"net"
)

type OptIPv4Address struct {
	Address net.IP
}

func (o OptIPv4Address) String() string {
	return fmt.Sprintf(`%v`, o.Address)
}


func (o OptIPv4Address) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o OptIPv4Address) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {

	addr := make([]byte, 4)
	l, err := r.Read(addr)
	if err != nil {
		return base, err
	}

	if l != 4 {
		return base, fmt.Errorf(`len: %v`, l)
	}

	base.V = OptIPv4Address{Address: net.IPv4(addr[0], addr[1], addr[2], addr[3])}
	return base, nil
}
