package main

import (
	"./hardware"
	"fmt"
)

// Dynamic Host Configuration Protocol (DHCP)
// rfc2131
//
// https://tools.ietf.org/html/rfc2131#section-2
// https://tools.ietf.org/html/rfc2132
type HeaderRaw struct {
	// op 1 byte
	// Message op code / message type.
	// 1 = BOOTREQUEST, 2 = BOOTREPLY
	Op OpCode

	// Hardware address type.
	// htype 1 byte, rfc1060 page 46
	// See ARP section in "Assigned Numbers" RFC; e.g.,
	// 1 = 10mb ethernet.
	HardwareType hardware.HardwareType

	// Hardware address length.
	// hlen 1 byte
	// (e.g.  '6' for 10mb ethernet).
	// mac = 00:11:22:33:44:55 = 6
	HardwareAddressLength uint8

	// hops 1 byte
	// Client sets to zero, optionally used by relay agents
	Hops uint8

	// xid 4 bytes
	// Transaction ID, a random number chosen by the client,
	// used by the client and server to associate messages
	// and responses between a client and a server.
	TransactionID uint32

	// secs 2 bytes
	// Filled in by client, seconds elapsed since client
	// began address acquisition or renewal process.
	Seconds uint16

	// flags 2 bytes
	//
	//                     1 1 1 1 1 1
	// 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	// |B|             MBZ             |
	// +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
	//
	// B:  BROADCAST flag
	// MBZ:  MUST BE ZERO(reserved for future use)
	Flags uint16

	// ciaddr 4 bytes
	// Client IP address; only filled in if client is in
	// BOUND, RENEW or REBINDING state and can respond
	// to ARP requests.
	ClientAddress [4]byte

	// yiaddr 4 bytes
	// 'your' (client) IP address.
	YourAddress [4]byte

	// siaddr 4 bytes
	// IP address of next server to use in bootstrap;
	// returned in DHCPOFFER, DHCPACK by server.
	ServerAddress [4]byte

	// giaddr 4 bytes
	// Relay agent IP address, used in booting via a relay agent.
	RelayAgentAddress [4]byte

	// chaddr 16 bytes
	HardwareAddress [16]byte

	// sname 64 bytes
	// Optional server host name, null terminated string.
	ServerHostName [64]byte

	// file 128 bytes
	// Boot file name, null terminated string; "generic" name or
	// null in DHCPDISCOVER, fully qualified directory-path name in DHCPOFFER.
	File [128]byte

	// magic cookie 4 bytes
	// https://tools.ietf.org/html/rfc2132#section-2
	MagicCookie [4]byte // must be 99 130 83 99
}

func (s HeaderRaw) Validate() error {
	if !(s.MagicCookie[0] == 99 && s.MagicCookie[1] == 130 && s.MagicCookie[2] == 83 && s.MagicCookie[3] == 99) {
		return fmt.Errorf(`invalid magic cookie`)
	}

	if !(s.Op == BootReply || s.Op == BootRequest) {
		return fmt.Errorf(`invalid op code: %v`, s.Op)
	}

	return nil

}
