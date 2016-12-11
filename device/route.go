package device

import (
	"fmt"
	"net"
)

type route interface {
	Enabled() bool
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

type routeTable struct {
	routes []route
}

func (table *routeTable) addRoute(r route) {
	table.routes = append(table.routes, r)
}

func (table *routeTable) show() []string {
	routes := []string{}
	for _, r := range table.routes {
		if r.Enabled() {
			routes = append(routes, r.String())
		}
	}
	return routes
}
