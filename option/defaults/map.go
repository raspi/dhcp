package defaults

import (
	"../base"
	e "../enums"
	"../identifier"
	"../ipv4"
	"../meta"
	"../msgtype"
	"../opcodes"
	"../seconds"
	"../str"
	"bytes"
)

type TypeMap map[e.OpCode]func(*bytes.Reader) (base.OptionBase, error)

func GetDefaults() TypeMap {
	m := make(TypeMap)

	m[                 e.TypePad] = new(meta.PadOption).Read        // 0
	m[          e.TypeSubnetMask] = new(ipv4.OptNetmask).Read       // 1
	m[             e.TypeGateway] = new(ipv4.OptIPv4Address).Read   // 3
	m[    e.TypeDomainNameServer] = new(ipv4.OptIPv4Addresses).Read // 6
	m[  e.TypeRequestedIPAddress] = new(ipv4.OptIPv4Address).Read   // 50
	m[  e.TypeIPAddressLeaseTime] = new(seconds.Option).Read        // 51
	m[         e.TypeMessageType] = new(msgtype.Option).Read        // 53
	m[    e.TypeServerIdentifier] = new(ipv4.OptIPv4Address).Read   // 54
	m[e.TypeParameterRequestList] = new(opcodes.Option).Read        // 55
	m[         e.TypeRenewalTime] = new(seconds.Option).Read        // 58
	m[       e.TypeRebindingTime] = new(seconds.Option).Read        // 59
	m[    e.TypeClientIdentifier] = new(identifier.Option).Read     // 61
	m[      e.TypeTFTPServerName] = new(str.Option).Read            // 66
	m[                 e.TypeEnd] = new(meta.EndOption).Read        // 255

	return m
}
