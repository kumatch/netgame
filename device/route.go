package device

import (
	"net"
)

const (
    routeCodeConnected = iota
    routeCodeStatic
)

type route struct {
    code int
    address net.IPNet
    next net.IP
}

type routeTable struct {
    routes []*route
}
