package quest

import (
	"github.com/kumatch/netgame/ipnet"
)

type companion struct {
	ipAddress *ipnet.IPAddresss
}

func newCompanion(ipAddress *ipnet.IPAddresss) *companion {
	return &companion{
		ipAddress: ipAddress,
	}
}
