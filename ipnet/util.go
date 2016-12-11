package ipnet

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

// IPv4MaskFromString ex. mask = "255.255.0.0"
func IPv4MaskFromString(mask string) (net.IPMask, error) {
	p, err := parseipv4parts(mask)
	if err != nil {
		return nil, err
	}

	return net.IPv4Mask(p[0], p[1], p[2], p[3]), nil
}

func parseipv4parts(address string) (ipv4parts [4]byte, err error) {
	parts := strings.Split(address, ".")
	if len(parts) != 4 {
		err = fmt.Errorf("invalid format: %s", address)
		return
	}

	for i := 0; i < 4; i++ {
		v, e := strconv.ParseUint(parts[i], 10, 8)
		if e != nil {
			err = fmt.Errorf("invalid format: %s", address)
			return
		}
		ipv4parts[i] = byte(v)
	}
	return
}
