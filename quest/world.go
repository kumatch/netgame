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
	if !deviceSourceNetworkInterface.ReceivePacket(srcCompanionInterface.SendPacket()) {
		return false
	}

	// device -> destination companion
	if !destCompanionInterface.ReceivePacket(deviceDestinationNetworkInterface.SendPacket()) {
		return false
	}

	// destination companion -> device
	if !deviceDestinationNetworkInterface.ReceivePacket(destCompanionInterface.SendPacket()) {
		return false
	}

	// device -> source companion
	if !srcCompanionInterface.ReceivePacket(deviceSourceNetworkInterface.SendPacket()) {
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
