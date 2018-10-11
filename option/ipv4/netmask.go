package ipv4

import (
	OPT "../base"
	"bytes"
	"fmt"
	"io"
	"net"
)

type OptNetmask struct {
	Mask net.IPMask
}

func (o OptNetmask) String() string {
	cidr, _ := o.Mask.Size()

	return fmt.Sprintf(`CIDR:%v Mask:%v.%v.%v.%v (0x%v)`, cidr, o.Mask[0], o.Mask[1], o.Mask[2], o.Mask[3], o.Mask)
}

func (o OptNetmask) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o OptNetmask) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {
	addr := make([]byte, 4)

	l, err := r.Read(addr)
	if err != nil {
		return base, err
	}

	if l != 4 {
		return base, fmt.Errorf(`len: %v`, l)
	}

	base.V = OptNetmask{Mask: net.IPv4Mask(addr[0], addr[1], addr[2], addr[3])}
	return base, nil
}
