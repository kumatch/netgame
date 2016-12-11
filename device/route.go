package device

import (
	"fmt"
	"net"
)

type route interface {
	String() string
}

type connectedRoute struct {
	netIF *NetInterface
}

func newConnectedRoute(netIF *NetInterface) *connectedRoute {
	return &connectedRoute{
		netIF: netIF,
	}
}

func (r *connectedRoute) Enabled() bool {
	return r.netIF.status == true
}

func (r *connectedRoute) SameNetwork(network *net.IPNet) bool {
	return r.netIF.l3Address.Network.String() == network.String()
}

func (r *connectedRoute) ContainIP(ip net.IP) bool {
	return r.netIF.l3Address.Network.Contains(ip)
}

func (r *connectedRoute) String() string {
	network := r.netIF.l3Address.Network
	return fmt.Sprintf("%s is directly connected, %s", network.String(), r.netIF.GetName())
}

type staticRoute struct {
	netIF   *NetInterface
	network *net.IPNet
	nextIP  net.IP
}

func newStaticRoute(netIF *NetInterface, network *net.IPNet, nextIP net.IP) *staticRoute {
	return &staticRoute{
		netIF:   netIF,
		network: network,
		nextIP:  nextIP,
	}
}

func (r *staticRoute) Enabled() bool {
	return r.netIF.status == true
}

func (r *staticRoute) String() string {
	return fmt.Sprintf("%s via %s, %s", r.network.String(), r.nextIP, r.netIF.GetName())
}

func (r *staticRoute) SameNetwork(network *net.IPNet) bool {
	return r.netIF.l3Address.Network.String() == network.String()
}

type routeTable struct {
	connections map[string]*connectedRoute
	statics     map[string]*staticRoute
}

func newRouteTable() *routeTable {
	return &routeTable{
		connections: map[string]*connectedRoute{},
	}
}

func (table *routeTable) addConnectedRoute(r *connectedRoute) {
	name := r.netIF.GetName()
	table.connections[name] = r
}

func (table *routeTable) matchesRoute(network *net.IPNet) route {
	for _, r := range table.connections {
		if r.SameNetwork(network) {
			return r
		}
	}
	return nil
}

func (table *routeTable) show() []string {
	routes := []string{}
	for _, r := range table.connections {
		if r.Enabled() {
			routes = append(routes, r.String())
		}
	}
	return routes
}
