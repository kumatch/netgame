package device

import (
	"net"
	"testing"

	"github.com/kumatch/netgame/ipnet"
)

func TestSendPacket(t *testing.T) {
	ipAddress, _ := ipnet.NewIPAddressByCIDR("192.168.0.100/24")
	receiver := &NetInterface{
		status:    true,
		l3Address: ipAddress,
	}

	{
		sender := &NetInterface{
			status: false,
		}
		if sender.SendPacket(receiver) {
			t.Errorf("send packet on down interface")
		}
	}
	{
		sender := &NetInterface{
			status: true,
		}
		if sender.SendPacket(receiver) {
			t.Errorf("send packet on has no ip address interface")
		}
	}
	{
		ng, _ := ipnet.NewIPAddressByCIDR("192.168.0.255/24")
		sender := &NetInterface{
			status:    true,
			l3Address: ng,
		}
		if sender.SendPacket(receiver) {
			t.Errorf("send packet on broadcast address interface")
		}
	}
	{
		ng, _ := ipnet.NewIPAddressByCIDR("192.168.0.0/24")
		sender := &NetInterface{
			status:    true,
			l3Address: ng,
		}
		if sender.SendPacket(receiver) {
			t.Errorf("send packet on network address interface")
		}
	}
	{
		ng, _ := ipnet.NewIPAddressByCIDR("192.168.0.1/24")
		sender := &NetInterface{
			status:    true,
			l3Address: ng,
		}
		if !sender.SendPacket(receiver) {
			t.Errorf("cannot send packet")
		}
	}
}

func TestReceivePacket(t *testing.T) {
	ip := net.ParseIP("192.168.0.100")

	{
		netIF := &NetInterface{
			status: false,
		}
		if netIF.ReceivePacket(ip) {
			t.Errorf("Receive packet on down interface.")
		}
	}
	{
		netIF := &NetInterface{
			status: true,
		}
		if netIF.ReceivePacket(ip) {
			t.Errorf("Receive packet on has no IP address interface.")
		}
	}
	{
		ng, _ := ipnet.NewIPAddressByCIDR("192.168.0.100/24")
		netIF := &NetInterface{
			status:    true,
			l3Address: ng,
		}
		if netIF.ReceivePacket(ip) {
			t.Errorf("Receive packet on same IP address.")
		}
	}
	{
		ng, _ := ipnet.NewIPAddressByCIDR("192.168.0.1/24")
		netIF := &NetInterface{
			status:    true,
			l3Address: ng,
		}
		netIF.status = true
		if !netIF.ReceivePacket(ip) {
			t.Errorf("Cannot receive packet")
		}
	}
	{
		ng, _ := ipnet.NewIPAddressByCIDR("192.168.1.1/24")
		netIF := &NetInterface{
			status:    true,
			l3Address: ng,
		}
		if netIF.ReceivePacket(ip) {
			t.Errorf("Receive packet on different IP network.")
		}
	}
	{
		ng, _ := ipnet.NewIPAddressByCIDR("192.168.0.1/16")
		netIF := &NetInterface{
			status:    true,
			l3Address: ng,
		}
		if !netIF.ReceivePacket(ip) {
			t.Errorf("Cannot receive packet, different IP network but can take.")
		}
	}
}

func TestIfShow(t *testing.T) {
	mac := [6]uint8{10, 20, 128, 250, 0, 32}
	i := &NetInterface{
		name:      "eth1",
		status:    false,
		ifType:    InterfaceTypes[0],
		l2Address: mac,
	}
	ipAdddress, err := ipnet.NewIPAddressByCIDR("192.168.24.10/24")
	if err != nil {
		t.Error(err)
	}

	i.SetIPAddress(ipAdddress)
	v := i.Show()
	t.Log(v)
}
