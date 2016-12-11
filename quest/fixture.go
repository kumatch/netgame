package quest

import (
	"net"

	"github.com/kumatch/netgame/ipnet"
)

type questFixture struct {
	srcIPNetwork  *net.IPNet
	srcIPAddress  *ipnet.IPAddresss
	destIPNetwork *net.IPNet
	destIPAddress *ipnet.IPAddresss
}

func newQuestFixture() *questFixture {
	f := &questFixture{}
	f.srcIPNetwork = createRandomIPNetwork()
	f.srcIPAddress = createRandomIPAddress(f.srcIPNetwork)
	f.destIPNetwork = createRandomIPNetwork()
	f.destIPAddress = createRandomIPAddress(f.destIPNetwork)

	return f
}
