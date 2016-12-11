package ipnet

import (
	"net"
	"strconv"
)

type IPAddresss struct {
	IP      net.IP
	Network *net.IPNet
}

func NewIPAddress(ip net.IP, ipn *net.IPNet) *IPAddresss {
	return &IPAddresss{
		IP:      ip,
		Network: ipn,
	}
}

func NewIPAddressByCIDR(cidr string) (*IPAddresss, error) {
	ip, ipn, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	return &IPAddresss{
		IP:      ip,
		Network: ipn,
	}, nil
}

func (a *IPAddresss) String() string {
	mask, _ := a.Network.Mask.Size()
	return a.IP.String() + "/" + strconv.Itoa(mask)
}

func (a *IPAddresss) getBroadcastAddress() *IPAddresss {
	ip := make(net.IP, net.IPv4len)
	copy(ip, a.Network.IP)

	for i, v := range a.Network.Mask {
		ip[i] |= 255 ^ v
	}

	return NewIPAddress(ip, a.Network)
}

func (a *IPAddresss) IsSame(b *IPAddresss) bool {
	if a == nil || b == nil {
		return false
	}

	return a.String() == b.String()
}

func (a *IPAddresss) IsSameNetwork(b *IPAddresss) bool {
	if a.Network == nil || b.Network == nil {
		return false
	}

	return a.Network.String() == b.Network.String()
}

func (a *IPAddresss) IsBroadcastAddress() bool {
	return a.String() == a.getBroadcastAddress().String()
}

func (a *IPAddresss) IsNetworkAddress() bool {
	return a.String() == a.Network.String()
}
