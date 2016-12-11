package device

import (
	"strconv"
)

type Device struct {
	interfaces []*NetInterface
	routeTable *routeTable 
}

func (d *Device) GetInterfaceSize() int {
	return len(d.interfaces)
}

func (d *Device) GetInterfaces() []*NetInterface {
	size := len(d.interfaces)
	interfaces := make([]*NetInterface, size, size)

	for index, netIF := range d.interfaces {
		interfaces[index] = netIF
	}
	return interfaces
}

func (d *Device) GetInterface(name string) *NetInterface {
	for _, netIF := range d.interfaces {
		if netIF.name == name {
			return netIF
		}
	}
	return nil
}

func (d *Device) GetInterfaceByIdx(idx int) *NetInterface {
	if idx < 0 || idx >= len(d.interfaces) {
		return nil
	}
	return d.interfaces[idx]
}

func NewDevice(ifType *InterfaceType, macAddress MacAddress, ifNum int) *Device {
	d := &Device{
		interfaces: make([]*NetInterface, ifNum),
		routeTable: &routeTable{},
	}
	for i := 0; i < ifNum; i++ {
		n := i + 1
		mac := macAddress.Increase(n)
		d.interfaces[i] = createNetInterface(ifType, mac, n)
	}

	return d
}

func createNetInterface(ifType *InterfaceType, macAddress MacAddress, num int) *NetInterface {
	name := "eth" + strconv.Itoa(num)
	return NewNetInterface(name, ifType, macAddress)
}
