package msgtype

import (
	OPT "../base"
	"bytes"
	"fmt"
	"io"
)

type MessageType uint8

const (
	Unknown MessageType = iota
	Discover
	Offer
	Request
	Decline
	Acknowledge
	NotAcknowledge
	Release
	Inform
)

type Option struct {
	Type MessageType
}

func (o Option) String() string {
	switch o.Type {
	default:
		return "???"
	case Unknown:
		return "unknown"
	case Discover:
		return "discover"
	case Offer:
		return "offer"
	case Request:
		return "request"
	case Decline:
		return "decline"
	case Acknowledge:
		return "acknowledge"
	case NotAcknowledge:
		return "notacknowledge"
	case Release:
		return "release"
	case Inform:
		return "inform"
	}
}

func (o Option) Write(w io.Writer) error {
	//w.Write([]byte{TYPE_IP, 4})
	//w.Write(o.Address[12:16])
	return nil
}

func (o Option) Read(r *bytes.Reader) (base OPT.OptionBase, err error) {

	msgtype := make([]byte, 1)

	l, err := r.Read(msgtype)
	if err != nil {
		return base, err
	}

	if l != 1 {
		return base, fmt.Errorf(`len: %v`, l)
	}

	base.V = Option{Type: MessageType(msgtype[0])}
	return base, nil
}
