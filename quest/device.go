package quest

import (
	"math/rand"

	"github.com/kumatch/netgame/device"
)

type questDevice struct {
	device                     *device.Device
	srcNetworkInterfaceNumber  int
	destNetworkInterfaceNumber int
}

func (d *questDevice) getSourceNetworkInterface() *device.NetInterface {
	return d.device.GetInterfaceByIdx(d.srcNetworkInterfaceNumber)
}

func (d *questDevice) getDestinationNetworkInterface() *device.NetInterface {
	return d.device.GetInterfaceByIdx(d.destNetworkInterfaceNumber)
}

func pickupInterfaceNumbers(number int) (if1, if2 int) {
	if number < 2 {
		panic("set 2 or more integer")
	}

	if1 = rand.Intn(number)
	for {
		if2 = rand.Intn(number)
		if if1 != if2 {
			return
		}
	}
}

func newQuestDevice() *questDevice {
	ifType := createRandomInterfaceType()
	macAddress := createRandomHardwareAddress()
	num := createRandomInterfaceNumber()

	d := device.NewDevice(ifType, macAddress, num)

	if1, if2 := pickupInterfaceNumbers(num)
	qd := &questDevice{device: d}
	qd.srcNetworkInterfaceNumber = if1
	qd.destNetworkInterfaceNumber = if2

	return qd
}
