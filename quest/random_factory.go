package quest

import (
	"fmt"
	"math"
	"math/rand"
	"net"
	"time"

	"github.com/kumatch/netgame/device"
	"github.com/kumatch/netgame/ipnet"
)

func createRandom8BitIPAddress() int {
	return 1 + rand.Intn(255)
}

func createRandomCIDRMask() int {
	return 20 + rand.Intn(11)
}

// createRandomIPNetwork TODO exclude multicast, unitcast...
func createRandomIPNetwork() *net.IPNet {
	a := make([]int, 4)
	for i := 0; i < 4; i++ {
		a[i] = createRandom8BitIPAddress()
	}
	cidr := fmt.Sprintf("%d.%d.%d.%d/%d", a[0], a[1], a[2], a[3], createRandomCIDRMask())
	_, ipn, err := net.ParseCIDR(cidr)
	if err != nil {
		panic("Failed to net.ParseCIDR(): " + err.Error())
	}

	return ipn
}

func createRandomIPAddress(ipn *net.IPNet) *ipnet.IPAddresss {
	mask, _ := ipn.Mask.Size()
	addressNumber := int(math.Exp2(float64(32 - mask)))
	hostNumber := addressNumber - 2
	additionalNumber := 1 + rand.Intn(hostNumber)

	ip := make(net.IP, net.IPv4len)
	copy(ip, ipn.IP)
	for i := 0; i < additionalNumber; i++ {
		ip = incrementAddress(ip)
	}

	return ipnet.NewIPAddress(ip, ipn)
}

func createRandomInterfaceType() *device.InterfaceType {
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(device.InterfaceTypes))
	return device.InterfaceTypes[index]
}

func createRandomHardwareAddress() device.MacAddress {
	m := [6]uint8{}
	for i := 0; i < 5; i++ {
		m[i] = uint8(rand.Intn(256))
	}
	m[5] = uint8(0)

	return m
}

func incrementAddress(ip net.IP) net.IP {
	nextIP := ip
	for i := len(nextIP) - 1; i > -1; i-- {
		nextIP[i]++
		if nextIP[i] != 0 {
			break
		}
	}
	return nextIP
}
