package quest

import (
	"github.com/kumatch/netgame/ipnet"
	"testing"
)

func TestNewCompanion(t *testing.T) {
	ip, _ := ipnet.NewIPAddressByCIDR("192.168.0.1/24")

	c := newCompanion(ip)
	if c.NetIF.GetName() != "eth1" {
		t.Errorf("want eth1, got %s", c.NetIF.GetName())
	}
	if c.ipAddress != ip {
		t.Errorf("want %s, got %s", ip, c.ipAddress)
	}
}
