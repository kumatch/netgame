package terminal

import (
	"github.com/chzyer/readline"
	"github.com/kumatch/netgame/device"
)

const (
	stateInitialHostname = "Network"
)

type state struct {
	instance     *readline.Instance
	device       *device.Device
	mode         mode
	hostname     string
	configIFName string
	fin          bool
}

func (s *state) GetPrompt() string {
	return s.mode.prompt(s.hostname)
}

func (s *state) SetGlobalMode() {
	s.mode = &globalMode{}
	s.instance.SetPrompt(s.GetPrompt())

	commands := newGlobalCommands()
	s.instance.Config.AutoComplete = commands.completer
}

func (s *state) isGlobalMode() bool {
	_, ok := s.mode.(*globalMode)
	return ok
}

func (s *state) SetConfigureMode() {
	s.mode = &configureMode{}
	s.instance.SetPrompt(s.GetPrompt())

	commands := newConfigureCommands()
	s.instance.Config.AutoComplete = commands.completer
}

func (s *state) isConfigureMode() bool {
	_, ok := s.mode.(*configureMode)
	return ok
}

func (s *state) SetConfigureInterfaceMode(configIFName string) {
	s.mode = &configureInterfaceMode{}
	s.instance.SetPrompt(s.GetPrompt())

	commands := newConfigureInterfaceCommands()
	s.instance.Config.AutoComplete = commands.completer
	s.configIFName = configIFName
}

func (s *state) isConfigureInterfaceMode() bool {
	_, ok := s.mode.(*configureInterfaceMode)
	return ok
}

func (s *state) SetHostname(hostname string) {
	s.hostname = hostname
	s.instance.SetPrompt(s.GetPrompt())
}

func (s *state) Exit() {
	s.fin = true
}

func (s *state) cleanup() {
	s.device = nil
	s.configIFName = ""
	s.hostname = stateInitialHostname
	s.fin = false

	s.SetGlobalMode()
}

func (s *state) reset(d *device.Device) {
	s.cleanup()
	s.device = d
}

func newState(instance *readline.Instance) *state {
	state := &state{instance: instance}
	state.cleanup()

	return state
}
