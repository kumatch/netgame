package quest

type questWorld struct {
	device        *questDevice
	srcCompanion  *companion
	destCompanion *companion
}

func (w *questWorld) verify() bool {
	if !w.device.getSourceNetworkInterface().Ping(w.srcCompanion.ipAddress) {
		return false
	}

	if !w.device.getDestinationNetworkInterface().Ping(w.destCompanion.ipAddress) {
		return false
	}

	return true
}

func newQuestWorld(f *questFixture) *questWorld {
	return &questWorld{
		srcCompanion: newCompanion(f.srcIPAddress), 			
		destCompanion: newCompanion(f.destIPAddress),
		device:        newQuestDevice(),
	}
}
