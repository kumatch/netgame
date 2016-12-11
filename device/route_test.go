package device

import (
	"net"
	"testing"

	"github.com/kumatch/netgame/ipnet"
)

func TestConnectedRoute(t *testing.T) {
	ng, _ := ipnet.NewIPAddressByCIDR("192.168.0.1/24")
	netIF := &NetInterface{
		status:    false,
		l3Address: ng,
	}
	route := newConnectedRoute(netIF)

	if route.Enabled() {
		t.Error("status false but route enabled")
	}

	ip1, n1, _ := net.ParseCIDR("192.168.0.100/24")
	if !route.SameNetwork(n1) {
		t.Errorf("%s and %s,  but not same network", route.netIF, n1)
	}
	if !route.ContainIP(ip1) {
		t.Errorf("%s and %s,  but not contain", route.netIF, ip1)
	}

	ip2, n2, _ := net.ParseCIDR("192.168.1.100/24")
	if route.SameNetwork(n2) {
		t.Errorf("%s and %s,  but same network", route.netIF, n2)
	}
	if route.ContainIP(ip2) {
		t.Errorf("%s and %s,  but contain", route.netIF, ip2)
	}

	ip3, n3, _ := net.ParseCIDR("192.168.0.100/16")
	if route.SameNetwork(n3) {
		t.Errorf("%s and %s,  but same network", route.netIF, n3)
	}
	if !route.ContainIP(ip3) {
		t.Errorf("%s and %s,  but not contain", route.netIF, ip3)
	}
}
