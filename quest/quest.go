package quest

import (
	"context"
	"time"

	"github.com/kumatch/netgame/quest/spec"
	"github.com/kumatch/netgame/terminal"
)

func Run() {
	cxt, cancel := context.WithCancel(context.Background())
	deliver, verifier := spec.NewSpecDeliver()
	term := terminal.NewTerminal()
	open("http://localhost:18080")

	go quest(cxt, term, deliver, verifier)

	term.Run()
	cancel()
}

func quest(parent context.Context, term *terminal.Terminal, deliver *spec.Deliver, verifier *spec.Verifier) {
	var fin bool
	ctx, cancel := context.WithCancel(parent)

	for {
		fixture := newQuestFixture()
		world := newQuestWorld(fixture)
		sheet := createSpecSheet(fixture, world)

		term.D <- world.device.device

		succeed := verify(ctx, world, verifier)

		go deliver.Delivery(sheet)

		func() {
			for {
				select {
				case <-succeed:
					return
				case <-parent.Done():
					cancel()
					fin = true
					return
				}
			}
		}()

		if fin {
			return
		}
	}
}

func verify(parent context.Context, world *questWorld, verifier *spec.Verifier) chan struct{} {
	succeed := make(chan struct{})

	go func() {
		for {
			select {
			case <-verifier.ReceiveRequest():
				result := world.verify()
				verifier.Result(result)
				if result {
					time.Sleep(5 * time.Second) 
					succeed <- struct{}{}
					return
				}
			case <-parent.Done():
				return
			}
		}
	}()

	return succeed
}

func createSpecSheet(fixture *questFixture, world *questWorld) *spec.Sheet {
	var seg1, seg2 *spec.Segment

	{
		network := fixture.srcIPNetwork
		destIP := world.srcCompanion.ipAddress
		netIF := world.device.getSourceNetworkInterface()
		seg1 = spec.NewSegment(network, destIP, netIF.GetName())
	}
	{
		network := fixture.destIPNetwork
		destIP := world.destCompanion.ipAddress
		netIF := world.device.getDestinationNetworkInterface()
		seg2 = spec.NewSegment(network, destIP, netIF.GetName())
	}

	sheet := spec.NewSpecSheet()
	sheet.AddSegument(seg1)
	sheet.AddSegument(seg2)

	return sheet
}
