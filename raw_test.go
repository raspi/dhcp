package main

import (
	"./hardware"
	"bytes"
	"fmt"
	"log"
	"net"
	"testing"
)

var (
	// Raw bytes where client is asking for IP address
	ClientDiscover = []byte{
		0x01, 0x01, 0x06, 0x00, 0x00, 0x00, 0x3d, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x0b, 0x82, 0x01, 0xfc, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82,
		0x53, 0x63, 0x35, 0x01, 0x01, 0x3d, 0x07, 0x01, 0x00, 0x0b, 0x82, 0x01, 0xfc, 0x42,
		0x32, 0x04, 0x00, 0x00, 0x00, 0x00, 0x37, 0x04, 0x01, 0x03, 0x06, 0x2a, 0xff, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00,}

	// Raw bytes where server is offering IP address
	ServerOffer = []byte{
		0x02, 0x01, 0x06, 0x00, 0x00, 0x00, 0x3d, 0x1d, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xc0, 0xa8, 0x00, 0x0a, 0xc0, 0xa8, 0x00, 0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x82, 0x01,
		0xfc, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82, 0x53, 0x63,
		0x35, 0x01, 0x02, 0x01, 0x04, 0xff, 0xff, 0xff, 0x00, 0x3a, 0x04, 0x00, 0x00, 0x07, 0x08, 0x3b,
		0x04, 0x00, 0x00, 0x0c, 0x4e, 0x33, 0x04, 0x00, 0x00, 0x0e, 0x10, 0x36, 0x04, 0xc0, 0xa8, 0x00,
		0x01, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	// Raw bytes where client is renewing allocated IP address
	ClientRequest = []byte{
		0x01, 0x01, 0x06, 0x00, 0x00, 0x00, 0x3d, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x82, 0x01,
		0xfc, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82, 0x53, 0x63,
		0x35, 0x01, 0x03, 0x3d, 0x07, 0x01, 0x00, 0x0b, 0x82, 0x01, 0xfc, 0x42, 0x32, 0x04, 0xc0, 0xa8,
		0x00, 0x0a, 0x36, 0x04, 0xc0, 0xa8, 0x00, 0x01, 0x37, 0x04, 0x01, 0x03, 0x06, 0x2a, 0xff, 0x00,
	}

	// Raw bytes where server is renewing allocated IP address
	ServerAcknowledge = []byte{
		0x02, 0x01, 0x06, 0x00, 0x00, 0x00, 0x3d, 0x1e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xc0, 0xa8, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x0b, 0x82, 0x01,
		0xfc, 0x42, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82, 0x53, 0x63,
		0x35, 0x01, 0x05, 0x3a, 0x04, 0x00, 0x00, 0x07, 0x08, 0x3b, 0x04, 0x00, 0x00, 0x0c, 0x4e, 0x33,
		0x04, 0x00, 0x00, 0x0e, 0x10, 0x36, 0x04, 0xc0, 0xa8, 0x00, 0x01, 0x01, 0x04, 0xff, 0xff, 0xff,
		0x00, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	ServerAuthentication = []byte{
		0x02, 0x01, 0x06, 0x01, 0x77, 0x71, 0xcf, 0x85, 0x00, 0x0a, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x0a, 0x0a, 0x08, 0xeb, 0xac, 0x16, 0xb2, 0xea, 0x0a, 0x0a, 0x08, 0xf0, 0x00, 0x0e, 0x86, 0x11,
		0xc0, 0x75, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82, 0x53, 0x63,
		0x35, 0x01, 0x02, 0x01, 0x04, 0xff, 0xff, 0xff, 0x00, 0x36, 0x04, 0xac, 0x16, 0xb2, 0xea, 0x33,
		0x04, 0x00, 0x00, 0xa8, 0xc0, 0x03, 0x04, 0x0a, 0x0a, 0x08, 0xfe, 0x06, 0x08, 0x8f, 0xd1, 0x04,
		0x01, 0x8f, 0xd1, 0x05, 0x01, 0x42, 0x0e, 0x31, 0x37, 0x32, 0x2e, 0x32, 0x32, 0x2e, 0x31, 0x37,
		0x38, 0x2e, 0x32, 0x33, 0x34, 0x78, 0x05, 0x01, 0xac, 0x16, 0xb2, 0xea, 0x3d, 0x10, 0x00, 0x6e,
		0x61, 0x74, 0x68, 0x61, 0x6e, 0x31, 0x63, 0x6c, 0x69, 0x65, 0x6e, 0x74, 0x69, 0x64, 0x5a, 0x1f,
		0x01, 0x01, 0x00, 0xc8, 0x78, 0xc4, 0x52, 0x56, 0x40, 0x20, 0x81, 0x31, 0x32, 0x33, 0x34, 0x8f,
		0xe0, 0xcc, 0xe2, 0xee, 0x85, 0x96, 0xab, 0xb2, 0x58, 0x17, 0xc4, 0x80, 0xb2, 0xfd, 0x30, 0x52,
		0x16, 0x01, 0x14, 0x20, 0x50, 0x4f, 0x4e, 0x20, 0x31, 0x2f, 0x31, 0x2f, 0x30, 0x37, 0x2f, 0x30,
		0x31, 0x3a, 0x31, 0x2e, 0x30, 0x2e, 0x31, 0xff,
	}

	ClientBootpOverload = []byte{
		0x01, 0x01, 0x06, 0x00, 0xac, 0x2e, 0xff, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x6c, 0x82,
		0xdc, 0x4e, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x38, 0x14, 0x73, 0x6e,
		0x61, 0x6d, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x6f, 0x76, 0x65, 0x72, 0x6c, 0x6f,
		0x61, 0x64, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x38, 0x18, 0x66, 0x69,
		0x6c, 0x65, 0x20, 0x6e, 0x61, 0x6d, 0x65, 0x20, 0x66, 0x69, 0x65, 0x6c, 0x64, 0x20, 0x6f, 0x76,
		0x65, 0x72, 0x6c, 0x6f, 0x61, 0x64, 0xff, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x63, 0x82, 0x53, 0x63,
		0x35, 0x01, 0x01, 0x39, 0x02, 0x02, 0x4e, 0x37, 0x04, 0x01, 0x1c, 0x03, 0x2b, 0x33, 0x04, 0x00,
		0x00, 0x0e, 0x10, 0x34, 0x01, 0x03, 0x38, 0x07, 0x50, 0x61, 0x64, 0x64, 0x69, 0x6e, 0x67, 0x00,
		0x3d, 0x07, 0x01, 0x00, 0x00, 0x6c, 0x82, 0xdc, 0x4e, 0xff,
	}
)

func flagsCompare(actual Flags, expected Flags) error {
	if actual.Broadcast != expected.Broadcast {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed1 != expected.NotUsed1 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed2 != expected.NotUsed2 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed3 != expected.NotUsed3 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed4 != expected.NotUsed4 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed5 != expected.NotUsed5 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed6 != expected.NotUsed6 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed7 != expected.NotUsed7 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed8 != expected.NotUsed8 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed9 != expected.NotUsed9 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed10 != expected.NotUsed10 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed11 != expected.NotUsed11 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed12 != expected.NotUsed12 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed13 != expected.NotUsed13 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed14 != expected.NotUsed14 {
		return fmt.Errorf(`flag`)
	}

	if actual.NotUsed15 != expected.NotUsed15 {
		return fmt.Errorf(`flag`)
	}

	return nil

}

func TestDiscover(t *testing.T) {
	hdr, opts, err := readPacket(bytes.NewReader(ClientDiscover))

	if err != nil {
		log.Printf(`%#v`, err)
		panic(err)
	}

	if hdr.Op != BootRequest {
		t.Fatalf(`not boot request`)
	}

	if hdr.Hardware != hardware.Ethernet {
		t.Fatalf(`type not ethernet`)
	}

	if hdr.Hops != 0 {
		t.Fatalf(`hops not 0`)
	}

	expectedMac := net.HardwareAddr{0x0, 0xb, 0x82, 0x1, 0xfc, 0x42}
	for i := 0; i < 6; i++ {
		expected := expectedMac[i]
		actual := hdr.HardwareAddress[i]

		if actual != expected {
			t.Fatalf(`hwaddr`)
		}

	}

	expectedFlags := Flags{
		Broadcast: false,
		NotUsed1:  false,
		NotUsed2:  false,
		NotUsed3:  false,
		NotUsed4:  false,
		NotUsed5:  false,
		NotUsed6:  false,
		NotUsed7:  false,
		NotUsed8:  false,
		NotUsed9:  false,
		NotUsed10: false,
		NotUsed11: false,
		NotUsed12: false,
		NotUsed13: false,
		NotUsed14: false,
		NotUsed15: false,
	}

	flagcmperr := flagsCompare(hdr.Flags, expectedFlags)

	if flagcmperr != nil {
		t.Fatalf(`flag %v`, flagcmperr)
	}

	log.Printf(`%#v`, hdr)
	log.Printf(`%v`, opts)
}
