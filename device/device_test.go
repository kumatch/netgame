package device

import (
	"fmt"
	"github.com/kumatch/netgame/ipnet"
	"testing"
)

func TestNewDevice(t *testing.T) {
	var macAddress MacAddress = [6]uint8{10, 20, 30, 40, 50, 60}
	ifType := InterfaceTypes[0]
	size := 5

	d := NewDevice(ifType, macAddress, size)

	// Device interfaces
	if len(d.interfaces) != size {
		t.Errorf("device interface size want %d, got %d", size, len(d.interfaces))
	}
	if d.GetInterfaceSize() != size {
		t.Errorf("device GetInterfaceSize() want %d, got %d", size, d.GetInterfaceSize())
	}
	for i := 0; i < size; i++ {
		name := fmt.Sprintf("eth%d", i+1)
		mac := macAddress.Increase(i + 1)
		if d.interfaces[i].GetName() != name {
			t.Errorf("device interface name want %v, got %v", name, d.interfaces[i].GetName())
		}
		if d.interfaces[i].ifType != ifType {
			t.Errorf("device interface type want %v, got %v", ifType, d.interfaces[i].ifType)
		}
		if d.interfaces[i].l2Address != mac {
			t.Errorf("device interface L2 address want %v, got %v", mac, d.interfaces[i].l2Address)
		}
		if d.interfaces[i].l3Address != nil {
			t.Errorf("device interface L3 address want nil, got %v", d.interfaces[i].l3Address)
		}
		if d.interfaces[i].status != false {
			t.Errorf("device interface status want false, got true")
		}
	}

	// Device route table
	if len(d.routeTable.connections) != 0 {
		t.Errorf("device route table connected routes want empty, got %v", d.routeTable.connections)
	}
	if len(d.routeTable.statics) != 0 {
		t.Errorf("device static table connected routes want empty, got %v", d.routeTable.statics)
	}
}

func TestGetInterface(t *testing.T) {
	d := &Device{
		interfaces: []*NetInterface{
			&NetInterface{name: "foo"},
			&NetInterface{name: "bar"},
			&NetInterface{name: "baz"},
		},
	}

	if d.getInterface("bar").GetName() != "bar" {
		t.Errorf(`device getInterface("bar") want bar NetInterface, got %s`, d.getInterface("bar").GetName())
	}
}

func TestGetInterfaceByIdx(t *testing.T) {
	d := &Device{
		interfaces: []*NetInterface{
			&NetInterface{name: "foo"},
			&NetInterface{name: "bar"},
			&NetInterface{name: "baz"},
		},
	}

	if d.GetInterfaceByIdx(1).GetName() != "bar" {
		t.Errorf(`device getInterfaceByIdx(1) want bar NetInterface, got %s`, d.GetInterfaceByIdx(1).GetName())
	}
}

func TestSetIPAddress(t *testing.T) {
	netIF := &NetInterface{name: "foo"}
	routeTable := newRouteTable()
	d := &Device{
		interfaces: []*NetInterface{netIF},
		routeTable: routeTable,
	}

	ip, _ := ipnet.NewIPAddressByCIDR("192.168.0.1/24")
	d.SetIPAddress("foo", ip)

	if netIF.l3Address != ip {
		t.Errorf("want %s, got %s", ip, netIF.l3Address)
	}
	if routeTable.connections["foo"].netIF != netIF {
		t.Errorf("want %s, got %s", netIF, routeTable.connections["foo"].netIF)
	}
}
