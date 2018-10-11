package ipv4

import (
	OPT "../base"
	"bytes"
	"fmt"
	"io"
	"net"
)

// Multiple IPv4 addresses
type OptIPv4Addresses struct {
	Addresses []net.IP
}

func (o OptIPv4Addresses) String() string {
	return fmt.Sprintf(`%+v`, o.Addresses)
}

func (o OptIPv4Addresses) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o OptIPv4Addresses) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {

	addrs := OptIPv4Addresses{}

	for {
		addr := make([]byte, 4)
		l, err := r.Read(addr)
		if err != nil {
			if err == io.EOF {
				break
			}
			return base, err
		}
		if l != 4 {
			return base, fmt.Errorf(`len: %v`, l)
		}

		addrs.Addresses = append(addrs.Addresses, net.IPv4(addr[0], addr[1], addr[2], addr[3]))

	}

	base.V = addrs
	return base, nil
}
