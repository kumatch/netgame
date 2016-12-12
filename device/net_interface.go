package device

import (
	"bytes"
	"net"
	"strconv"
	"text/template"

	"github.com/kumatch/netgame/ipnet"
)

type ifStatus bool

type MacAddress [6]uint8

func (m MacAddress) String() (mac string) {
	for i, v := range m {
		s := strconv.FormatUint(uint64(v), 16)
		if len(s) > 1 {
			mac += s
		} else {
			mac += "0" + s
		}
		if (i%2) > 0 && i < 5 {
			mac += "."
		}
	}
	return
}

func (m MacAddress) Increase(num int) MacAddress {
	var mac MacAddress
	for i, v := range m {
		mac[i] = v
	}
	mac[5] += uint8(num)

	return mac
}

func (i ifStatus) String() string {
	if i {
		return "up"
	}
	return "down"
}

type InterfaceType struct {
	hwName   string
	bandWith int
	delay    int
}

type NetInterface struct {
	name      string
	status    ifStatus
	ifType    *InterfaceType
	l2Address MacAddress
	l3Address *ipnet.IPAddresss
}

func (netIF *NetInterface) GetName() string {
	return netIF.name
}

func (netIF *NetInterface) Up() {
	netIF.status = true
}

func (netIF *NetInterface) Down() {
	netIF.status = false
}

func (netIF *NetInterface) SetIPAddress(ipAddress *ipnet.IPAddresss) {
	netIF.l3Address = ipAddress
}

func (netIF *NetInterface) SendPacket() net.IP {
	if !netIF.status {
		return nil
	}
	if netIF.l3Address == nil {
		return nil
	}
	if netIF.l3Address.IsBroadcastAddress() {
		return nil
	}
	if netIF.l3Address.IsNetworkAddress() {
		return nil
	}
	return netIF.l3Address.IP
}

func (netIF *NetInterface) ReceivePacket(ip net.IP) bool {
	if !netIF.status {
		return false
	}
	if netIF.l3Address == nil {
		return false
	}
	if ip == nil {
		return false
	}
	if !netIF.l3Address.Network.Contains(ip) {
		return false
	}
	return netIF.l3Address.IP.String() != ip.String()
}

func (i *NetInterface) Show() string {
	text := `{{.name}} is {{.isUp}}
Hardware is {{.ifTypeName}} Port, address is {{.mac}}
Internet address is {{.ip}}
MTU 1500 bytes, BW {{.bw}} Kbit, DLY {{.dly}} usec
Encapsulation ARPA, loopback not set
Keepalive set (10 sec)
Auto-duplex, Auto Speed`

	var ip string
	if i.l3Address != nil {
		ip = i.l3Address.String()
	} else {
		ip = "unknown"
	}

	values := map[string]string{
		"name":       i.name,
		"isUp":       i.status.String(),
		"ifTypeName": i.ifType.hwName,
		"mac":        i.l2Address.String(),
		"ip":         ip,
		"bw":         strconv.Itoa(i.ifType.bandWith),
		"dly":        strconv.Itoa(i.ifType.delay),
	}

	t, err := template.New("show_interfaces").Parse(text)
	if err != nil {
		panic(err)
	}

	buff := bytes.NewBufferString("")
	err = t.Execute(buff, values)
	if err != nil {
		panic(err)
	}

	return buff.String()
}

func NewNetInterface(name string, ifType *InterfaceType, l2Address MacAddress) *NetInterface {
	return &NetInterface{
		name:      name,
		status:    false,
		ifType:    ifType,
		l2Address: l2Address,
	}
}

var InterfaceTypes []*InterfaceType = []*InterfaceType{
	&InterfaceType{
		hwName:   "Ethernet",
		bandWith: 100000,
		delay:    100,
	},
	&InterfaceType{
		hwName:   "Gigabit Ethernet",
		bandWith: 1000000,
		delay:    10,
	},
	&InterfaceType{
		hwName:   "10 Gigabit Ethernet",
		bandWith: 10000000,
		delay:    10,
	},
}
