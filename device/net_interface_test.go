package device

import (
	"testing"

	"github.com/kumatch/netgame/ipnet"
)

func TestPing(t *testing.T) {
	ip, _ := ipnet.NewIPAddressByCIDR("192.168.0.101/24")

	ng1, _ := ipnet.NewIPAddressByCIDR("192.168.1.101/24")
	ng2, _ := ipnet.NewIPAddressByCIDR("192.168.0.102/25")
	ngIPs := []*ipnet.IPAddresss{ng1, ng2}

	{
		netIF := &NetInterface{}
		if netIF.Ping(ip) {
			t.Errorf("ping OK, but its invalid: ip=%s", ip)
		}
		for _, ng := range ngIPs {
			if netIF.Ping(ng) {
				t.Errorf("ping OK, but its invalid: ip=%s", ng)
			}
		}
	}

	{
		netIF := &NetInterface{}
        netIF.Up()
		if netIF.Ping(ip) {
			t.Errorf("ping OK, but its invalid: ip=%s", ip)
		}
		for _, ng := range ngIPs {
			if netIF.Ping(ng) {
				t.Errorf("ping OK, but its invalid: ip=%s", ng)
			}
		}
	}    

	{
    	setIP, _ := ipnet.NewIPAddressByCIDR("192.168.0.1/24")
        
		netIF := &NetInterface{}
        netIF.Up()
        netIF.SetIPAddress(setIP) 
		if !netIF.Ping(ip) {
			t.Errorf("ping NG, its invalid: ip=%s", ip)
		}
		for _, ng := range ngIPs {
			if netIF.Ping(ng) {
				t.Errorf("ping OK, but its invalid: ip=%s", ng)
			}
		}
	}

	{
    	setIP, _ := ipnet.NewIPAddressByCIDR("192.168.0.0/24")
        
		netIF := &NetInterface{}
        netIF.Up()
        netIF.SetIPAddress(setIP) 
		if netIF.Ping(ip) {
			t.Errorf("ping OK, but its invalid: ip=%s", ip)
		}
	}            

	{
    	setIP, _ := ipnet.NewIPAddressByCIDR("192.168.0.255/24")
        
		netIF := &NetInterface{}
        netIF.Up()
        netIF.SetIPAddress(setIP) 
		if netIF.Ping(ip) {
			t.Errorf("ping OK, but its invalid: ip=%s", ip)
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
	v, _ := i.Show()
	t.Log(v)
}
