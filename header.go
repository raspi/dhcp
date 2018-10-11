package main

import (
	"./hardware"
	"./option"
	"./option/enums"
	"bytes"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"time"
)

type OpCode uint8

const (
	Unknown OpCode = iota
	BootRequest
	BootReply
)

type Flags struct {
	Broadcast bool // bit  1 #1
	NotUsed1  bool // bit  2 #1
	NotUsed2  bool // bit  3 #1
	NotUsed3  bool // bit  4 #1
	NotUsed4  bool // bit  5 #1
	NotUsed5  bool // bit  6 #1
	NotUsed6  bool // bit  7 #1
	NotUsed7  bool // bit  8 #1
	NotUsed8  bool // bit  9 #2
	NotUsed9  bool // bit 10 #2
	NotUsed10 bool // bit 11 #2
	NotUsed11 bool // bit 12 #2
	NotUsed12 bool // bit 13 #2
	NotUsed13 bool // bit 14 #2
	NotUsed14 bool // bit 15 #2
	NotUsed15 bool // bit 16 #2
}

type Header struct {
	Op                OpCode
	Hardware          hardware.HardwareType
	Hops              uint8
	TransactionID     uint32
	Seconds           time.Duration
	Flags             Flags
	ClientAddress     net.IP
	YourAddress       net.IP
	ServerAddress     net.IP
	RelayAgentAddress net.IP
	HardwareAddress   net.HardwareAddr
	ServerHostName    string
	File              string
}

func byteToIP(b []byte) (net.IP, error) {
	l := len(b)
	if l != 4 {
		return nil, fmt.Errorf(`len: %v`, l)
	}
	return net.IPv4(b[0], b[1], b[2], b[3]), nil
}

func IPToByte(ip net.IP) ([4]byte) {
	b := make([]byte, 4)
	copy(b, ip.To4())

	return [4]byte{b[0], b[1], b[2], b[3]}
}

func readPacket(r *bytes.Reader) (h Header, opts []option.OptionRaw, err error) {
	var s HeaderRaw

	err = binary.Read(r, binary.LittleEndian, &s)
	if err != nil {
		return h, nil, err
	}

	err = s.Validate()
	if err != nil {
		return h, nil, err
	}

	h.Op = s.Op
	h.TransactionID = s.TransactionID
	h.Hops = s.Hops
	h.Seconds = time.Duration(s.Seconds) * time.Second
	h.Hardware = s.HardwareType
	h.HardwareAddress = s.HardwareAddress[0:s.HardwareAddressLength]

	h.ClientAddress, err = byteToIP(s.ClientAddress[:])
	if err != nil {
		return h, nil, err
	}

	h.YourAddress, err = byteToIP(s.YourAddress[:])
	if err != nil {
		return h, nil, err
	}

	h.ServerAddress, err = byteToIP(s.ServerAddress[:])
	if err != nil {
		return h, nil, err
	}

	h.RelayAgentAddress, err = byteToIP(s.RelayAgentAddress[:])
	if err != nil {
		return h, nil, err
	}

	for _, v := range s.File {
		if v != 0 {
			h.File += string(v)
		}
	}

	for _, v := range s.ServerHostName {
		if v != 0 {
			h.ServerHostName += string(v)
		}
	}

	for {

		// Read option code 0-255
		t_, err := r.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}

			return h, nil, err
		}

		t := enums.OpCode(t_)

		if t == enums.TypePad || t == enums.TypeEnd {
			opts = append(opts, option.OptionRaw{
				Type:   t,
				Length: 0,
				Value:  bytes.NewReader([]byte{}),
			})
			continue
		}

		// Read length 0-255
		l, err := r.ReadByte()
		if err != nil {
			return h, nil, err
		}

		// Read value which is length bytes long
		v := make([]byte, l)
		rl, err := r.Read(v)
		if err != nil {
			return h, nil, err
		}

		if rl != int(l) {
			return h, nil, fmt.Errorf(`len mismatch with option %v`, t)
		}

		opts = append(opts, option.OptionRaw{
			Type:   t,
			Length: l,
			Value:  bytes.NewReader(v),
		})
	}

	return h, opts, nil

}

func (h Header) Write() (raw HeaderRaw, err error) {
	raw.Op = h.Op
	raw.HardwareType = h.Hardware
	//raw.HardwareAddress = h.HardwareAddress
	raw.Hops = h.Hops
	raw.TransactionID = h.TransactionID
	raw.Seconds = uint16(h.Seconds.Seconds())
	//raw.Flags = h.Flags

	raw.ClientAddress = IPToByte(h.ClientAddress)
	raw.YourAddress = IPToByte(h.YourAddress)
	raw.ServerAddress = IPToByte(h.ServerAddress)
	raw.RelayAgentAddress = IPToByte(h.RelayAgentAddress)

	tmp := make([]byte, 64)
	copy(tmp[:], h.ServerHostName)

	tmp = make([]byte, 128)
	copy(tmp[:], h.File)

	return raw, nil
}
