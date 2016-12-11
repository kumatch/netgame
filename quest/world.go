package quest

type questWorld struct {
	device        *questDevice
	srcCompanion  *companion
	destCompanion *companion
}

func (w *questWorld) verify() bool {
	srcCompanionInterface := w.srcCompanion.NetIF
	destCompanionInterface := w.destCompanion.NetIF
	deviceSourceNetworkInterface := w.device.getSourceNetworkInterface()
	deviceDestinationNetworkInterface := w.device.getDestinationNetworkInterface()

	// source companion -> device
	if !srcCompanionInterface.SendPacket(deviceSourceNetworkInterface) {
		return false
	}

	// device -> destination companion
	if !deviceDestinationNetworkInterface.SendPacket(destCompanionInterface) {
		return false
	}

	// destination companion -> device
	if !destCompanionInterface.SendPacket(deviceDestinationNetworkInterface) {
		return false
	}

	// device -> source companion
	if !deviceSourceNetworkInterface.SendPacket(srcCompanionInterface) {
		return false
	}

	return true
}

func newQuestWorld(f *questFixture) *questWorld {
	return &questWorld{
		srcCompanion:  newCompanion(f.srcIPAddress),
		destCompanion: newCompanion(f.destIPAddress),
		device:        newQuestDevice(),
	}
}
