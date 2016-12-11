package device

import (
	"strconv"

	"github.com/kumatch/netgame/ipnet"
)

type Device struct {
	interfaces []*NetInterface
	routeTable *routeTable
}

func NewDevice(ifType *InterfaceType, macAddress MacAddress, ifNum int) *Device {
	d := &Device{
		interfaces: make([]*NetInterface, ifNum),
		routeTable: newRouteTable(),
	}
	for i := 0; i < ifNum; i++ {
		n := i + 1
		mac := macAddress.Increase(n)
		d.interfaces[i] = createNetInterface(ifType, mac, n)
	}

	return d
}

func (d *Device) getInterface(name string) *NetInterface {
	for _, netIF := range d.interfaces {
		if netIF.name == name {
			return netIF
		}
	}
	return nil
}

func (d *Device) GetInterfaceSize() int {
	return len(d.interfaces)
}

func (d *Device) GetInterfaces() []string {
	size := len(d.interfaces)
	interfaces := make([]string, size, size)
	for index, netIF := range d.interfaces {
		interfaces[index] = netIF.Show()
	}
	return interfaces
}

func (d *Device) GetInterfaceByIdx(idx int) *NetInterface {
	if idx < 0 || idx >= len(d.interfaces) {
		return nil
	}
	return d.interfaces[idx]
}

func (d *Device) SetIPAddress(name string, ipAddress *ipnet.IPAddresss) {
	netIF := d.getInterface(name)
	netIF.SetIPAddress(ipAddress)

	route := newConnectedRoute(netIF)
	d.routeTable.addConnectedRoute(route)
}

func (d *Device) SetInterfaceStatus(name string, up bool) {
	netIF := d.getInterface(name)
	if up {
		netIF.Up()
	} else {
		netIF.Down()
	}
}

func (d *Device) GetRouteTable() []string {
	return d.routeTable.show()
}

func (d *Device) Ping(destIF *NetInterface) bool {
	network := destIF.l3Address.Network
	route := d.routeTable.matchesRoute(network)
	if route == nil {
		return false
	}

	switch r := route.(type) {
	case *connectedRoute:
		senderIP := r.netIF.l3Address.IP
		return destIF.ReceivePacket(senderIP)
	}

	return false
}

func createNetInterface(ifType *InterfaceType, macAddress MacAddress, num int) *NetInterface {
	name := "eth" + strconv.Itoa(num)
	return NewNetInterface(name, ifType, macAddress)
}
