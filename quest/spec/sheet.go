package spec

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"unicode"

	"github.com/kumatch/netgame/ipnet"
)

type Segment struct {
	//IPNetwork       string `json:"ip_network"`
	DestinationHost string `json:"destination_host"`
	InterfaceName   string `json:"device_interface"`
}

func NewSegment(ipn *net.IPNet, dest *ipnet.IPAddresss, name string) *Segment {
	return &Segment{
		//IPNetwork:       ipn.String(),
		DestinationHost: destinationHost(dest),
		InterfaceName:   name,
	}
}

type Sheet struct {
	SpecNo   string     `json:"spec_no"`
	Segments []*Segment `json:"segments"`
}

func NewSpecSheet() *Sheet {
	alpha := string(65 + rand.Intn(26))
	num := rand.Intn(9999)
	return &Sheet{
		SpecNo: fmt.Sprintf("%s-%04d", alpha, num),
	}
}

func (s *Sheet) AddSegument(seg *Segment) {
	s.Segments = append(s.Segments, seg)
}

func destinationHost(ipAddress *ipnet.IPAddresss) string {
	var host string
	n := rand.Intn(30)
	if n > 1 {
		host = ipAddress.String()
	} else if n == 1 {
		host = up(ipAddress.String())
	} else {
		host = up1(ipAddress.String())
	}
	return host
}

func up(str string) string {
	c := unicode.SpecialCase{
		unicode.CaseRange{
			Lo: 0x0030, // 0
			Hi: 0x0039, // 9
			Delta: [unicode.MaxCase]rune{
				0xff10 - 0x0030, // UpperCase
				0,               // LowerCase
				0,               // TitleCase
			},
		},
	}
	return strings.ToUpperSpecial(c, str)
}

func up1(str string) string {
	index := rand.Intn(len(str))
	res := ""
	for i := 0; i < index; i++ {
		res += string(str[i])
	}
	res += up(string(str[index]))
	for j := index + 1; j < len(str); j++ {
		res += string(str[j])
	}
	return res
}
