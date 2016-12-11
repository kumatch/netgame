package spec

import (
    "fmt"
    "math/rand"
	"net"    

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
		DestinationHost: dest.String(),
		InterfaceName:   name,
	}
}

type Sheet struct {
    SpecNo string `json:"spec_no"`
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
