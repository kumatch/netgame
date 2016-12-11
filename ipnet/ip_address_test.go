package ipnet

import (
	"testing"
)

func TestNewIPAddressByCIDR(t *testing.T) {
	{
		address, err := NewIPAddressByCIDR("10.20.30.40/8")
		if err != nil {
			t.Errorf("raise error: %v", err)
		}

		if address.String() != "10.20.30.40/8" {
			t.Errorf("invalid IP address string. expect: 10.20.30.40/8, got: %s", address.String())
		}
		if address.IP.String() != "10.20.30.40" {
			t.Errorf("invalid IP address. expect: 10.20.30.40, got: %s", address.IP)
		}
		if address.Network.String() != "10.0.0.0/8" {
			t.Errorf("invalid IP network. expect: 10.0.0.0/8, got: %s", address.Network)
		}
		if address.Network.IP.String() != "10.0.0.0" {
			t.Errorf("invalid IP network address. expect: 10.0.0.0, got: %s", address.Network.IP)
		}
		mask, _ := address.Network.Mask.Size()
		if mask != 8 {
			t.Errorf("invalid IP network mask number. expect: 8, got: %d", mask)
		}
	}

	{
		invalidCIDR := []string{
			"foo",
			"10.20.30.40",
			"10.20.30.256/24",
			"10.20.30.40/33",
		}
		for _, cidr := range invalidCIDR {
			_, err := NewIPAddressByCIDR(cidr)
			if err == nil {
				t.Errorf("not raise error: cidr=%s", cidr)
			}
		}
	}
}

func TestSameAddress(t *testing.T) {
	a, _ := NewIPAddressByCIDR("192.168.0.100/24")
	b, _ := NewIPAddressByCIDR("192.168.0.100/24")
	c, _ := NewIPAddressByCIDR("192.168.0.101/24")

	if !a.IsSame(b) {
		t.Errorf("same ip address IsSame() is false: %s, %s", a, b)
	}
	if a.IsSame(c) {
		t.Errorf("not same ip address IsSame() is true: %s, %s", a, c)
	}
}

func TestSameNetworkAddress(t *testing.T) {
	{
		a, _ := NewIPAddressByCIDR("192.168.0.0/24")
		b, _ := NewIPAddressByCIDR("192.168.0.0/24")
		c, _ := NewIPAddressByCIDR("192.168.1.0/24")

		if !a.IsSameNetwork(b) {
			t.Errorf("same network address IsSameNetwork() is false: %s, %s", a, b)
		}
		if a.IsSameNetwork(c) {
			t.Errorf("not same network address IsSameNetwork() is true: %s, %s", a, c)
		}
	}

	{
		a, _ := NewIPAddressByCIDR("192.168.0.100/24")
		b, _ := NewIPAddressByCIDR("192.168.0.101/24")
		c, _ := NewIPAddressByCIDR("192.168.0.100/25")

		if !a.IsSameNetwork(b) {
			t.Errorf("same network address IsSameNetwork() is false: %s, %s", a, b)
		}
		if a.IsSameNetwork(c) {
			t.Errorf("not same network address IsSameNetwork() is true: %s, %s", a, c)
		}
		if b.IsSameNetwork(c) {
			t.Errorf("not same network address IsSameNetwork() is true: %s, %s", b, c)
		}
	}
}

func TestBroadcastAddress(t *testing.T) {
	data := map[string]string{
		"10.20.30.0/8":  "10.255.255.255/8",
		"10.20.30.0/9":  "10.127.255.255/9",
		"10.20.30.0/15": "10.21.255.255/15",
		"10.20.30.0/16": "10.20.255.255/16",
		"10.20.30.0/17": "10.20.127.255/17",
		"10.20.30.0/24": "10.20.30.255/24",
		"10.20.30.0/30": "10.20.30.3/30",
	}

	for cidr, expect := range data {
		ipAddress, err := NewIPAddressByCIDR(cidr)
		if err != nil {
			t.Errorf("raise error: %v", err)
		}

		broadcast := ipAddress.getBroadcastAddress()
		if expect != broadcast.String() {
			t.Errorf("expect = %s, got = %s", expect, broadcast)
		}
		if !broadcast.IsBroadcastAddress() {
			t.Errorf("%s is not broadcast address", broadcast)
		}
		if broadcast.IsNetworkAddress() {
			t.Errorf("%s is network address, its invalid.", broadcast)
		}
	}
}

func TestNetworkAddress(t *testing.T) {
	networkAddress, _ := NewIPAddressByCIDR("10.0.0.0/8")

	if !networkAddress.IsNetworkAddress() {
		t.Errorf("%s is not network address", networkAddress)
	}
	if networkAddress.IsBroadcastAddress() {
		t.Errorf("%s is broadcast address, its invalid", networkAddress)
	}
}
