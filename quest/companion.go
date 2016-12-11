package quest

import (
	"github.com/kumatch/netgame/device"
	"github.com/kumatch/netgame/ipnet"
)

type companion struct {
	ipAddress *ipnet.IPAddresss
	NetIF     *device.NetInterface
}

func newCompanion(ipAddress *ipnet.IPAddresss) *companion {
	netIF := device.NewNetInterface("eth1", createRandomInterfaceType(), createRandomHardwareAddress())
	netIF.Up()
	netIF.SetIPAddress(ipAddress)

	return &companion{
		ipAddress: ipAddress,
		NetIF:     netIF,
	}
}
