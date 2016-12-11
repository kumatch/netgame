package terminal

import (
	"fmt"
	"io"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/kumatch/netgame/ipnet"
)

func exitAction() *action {
	return &action{
		commands: []string{"exit"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				s.Exit()
			},
		},
	}
}

func endAction() *action {
	return &action{
		commands: []string{"end"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				s.SetGlobalMode()
			},
		},
	}
}

func setConfigureModeAction() *action {
	return &action{
		commands: []string{"configure", "terminal"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				s.SetConfigureMode()
				out.Write([]byte("Enter configuration command mode."))
			},
		},
	}
}

func setConfigureTerminalModeAction() *action {
	return &action{
		commands: []string{"interface"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				// static name.
				ifName := "eth"
				target := args[0]

				if !strings.HasPrefix(target, ifName) {
					outInvalidArgs(out, args)
					return
				}
				num, err := strconv.Atoi(strings.TrimLeft(target, ifName))
				if err != nil || num < 1 || num > s.device.GetInterfaceSize() {
					outInvalidArgs(out, args)
					return
				}
				s.SetConfigureInterfaceMode(target)
			},
			minArgNumebr: 1,
			maxArgNumber: 1,
		},
	}
}

func showInterfacesAction() *action {
	return &action{
		commands: []string{"show", "interfaces"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				for _, showedInterface := range s.device.GetInterfaces() {
					out.Write([]byte(showedInterface))
					out.Write([]byte(""))
				}
			},
			//maxArgNumber: 1
		},
	}
}

func showIPRouteAction() *action {
	return &action{
		commands: []string{"show", "ip", "route"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				for _, route := range s.device.GetRouteTable() {
					out.Write([]byte(route))
				}
			},
		},
	}
}

func showClockAction() *action {
	return &action{
		commands: []string{"show", "clock"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				t := time.Now()
				out.Write([]byte(t.Format("15:04:05.000 MST Mon Jan 2 2006")))
			},
		},
	}
}

func setHostnameAction() *action {
	return &action{
		commands: []string{"hostname"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				hostname := args[0]
				match, _ := regexp.MatchString(`^[a-zA-Z0-9\-]+$`, hostname)
				if !match {
					outInvalidArgs(out, args)
					return
				}
				s.SetHostname(hostname)
			},
			minArgNumebr: 1,
			maxArgNumber: 1,
		},
	}
}

// interfaces

func interfaceShutdownAction() *action {
	return &action{
		commands: []string{"shutdown"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				s.device.SetInterfaceStatus(s.configIFName, false)
			},
		},
	}
}

func interfaceNoShutdownAction() *action {
	return &action{
		commands: []string{"no", "shutdown"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				s.device.SetInterfaceStatus(s.configIFName, true)
			},
		},
	}
}

func interfaceSetAddressAction() *action {
	return &action{
		commands: []string{"ip", "address"},
		config: &actionConfig{
			method: func(args []string, s *state, out io.Writer) {
				ipv4String := args[0]
				maskString := args[1]

				mask, err := ipnet.IPv4MaskFromString(maskString)
				if err != nil {
					fmt.Println(err)
					outInvalidArgs(out, args)
				}
				cidr, _ := mask.Size()

				ipAddress, err := ipnet.NewIPAddressByCIDR(fmt.Sprintf("%s/%d", ipv4String, cidr))
				if err != nil {
					fmt.Println(err)
					outInvalidArgs(out, args)
				}

				s.device.SetIPAddress(s.configIFName, ipAddress)
			},
			minArgNumebr: 2,
			maxArgNumber: 2,
		},
	}
}
