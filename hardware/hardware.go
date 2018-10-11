package hardware

type HardwareType uint8

const (
	Unknown HardwareType = iota
	Ethernet
	ExperimentalEthernet
	AmateurRadioAX25
	TokenRing
	Chaos
	IEEE802
	ARCNET
	Hyperchannel
	Lanstar
	AutonetShortAddress
	LocalTalk
	LocalNet
)

func (t HardwareType) String() string {
	switch t {
	default:
		return "??"
	case Unknown:
		return "unknown"
	case Ethernet:
		return "ethernet"

	}
}
