package quest

import (
	"net"
	"regexp"
	"testing"
)

func TestCreateRandomCIDRMask(t *testing.T) {
	var i int
	for i = 0; i < 10000; i++ {
		r := createRandomCIDRMask()
		if r < 16 || r > 30 {
			t.Errorf("expect = 8 to 30, got = %d", r)
		}
	}
}

func TestCreateRandom8BitIPAddress(t *testing.T) {
	var i int
	for i = 0; i < 100000; i++ {
		r := createRandom8BitIPAddress()
		if r < 0 || r > 255 {
			t.Errorf("expect = 0 to 255, got = %d", r)
		}
	}
}

func TestCreateRandomIPNetwork(t *testing.T) {
	reg := regexp.MustCompile(`^[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}\.[\d]{1,3}/\d\d$`)
	var i int
	for i = 0; i < 100000; i++ {
		ipNet := createRandomIPNetwork()
		if !reg.Match([]byte(ipNet.String())) {
			t.Errorf("got: %v", ipNet.String())
		}
		if ipNet.IP[0] >= 224 {
			t.Errorf("got: %v, its class D or higher.", ipNet.String())
		}
		if ipNet.IP[0] == 169 && ipNet.IP[1] == 254 {
			t.Errorf("got: %v, its Link local.", ipNet.String())
		}
	}
}

func TestCreateRandomIPHost(t *testing.T) {
	{
		_, ipn, _ := net.ParseCIDR("192.168.250.0/30")
		a := []string{"192.168.250.1", "192.168.250.2"}

		for i := 0; i < 100; i++ {
			ipAddress := createRandomIPAddress(ipn)
			ip := ipAddress.IP.String()
			if ip != a[0] && ip != a[1] {
				t.Errorf("expect: %v, got: %s", a, ip)
			}
			if ipAddress.Network.IP.String() != ipn.IP.String() {
				t.Errorf("expect: %s, got: %s", ipn.IP, ipAddress.Network.IP)
			}
			if ipn.IP.String() != "192.168.250.0" {
				t.Errorf("broken original IP address.")
			}
		}
	}

	{
		_, ipn, _ := net.ParseCIDR("192.168.250.248/29")
		b := []string{
			"192.168.250.249",
			"192.168.250.250",
			"192.168.250.251",
			"192.168.250.252",
			"192.168.250.253",
			"192.168.250.254",
		}

		for i := 0; i < 100; i++ {
			ipAddress := createRandomIPAddress(ipn)
			ip := ipAddress.IP.String()

			if ip != b[0] && ip != b[1] && ip != b[2] && ip != b[3] && ip != b[4] && ip != b[5] {
				t.Errorf("expect = %v, got = %s", b, ip)
			}
			if ipAddress.Network.IP.String() != ipn.IP.String() {
				t.Errorf("expect: %s, got: %s", ipn.IP, ipAddress.Network.IP)
			}
			if ipn.IP.String() != "192.168.250.248" {
				t.Errorf("broken original IP address.")
			}
		}
	}
}
