package terminal

import (
	"github.com/chzyer/readline"
)

type terminalCommands struct {
	actions   []*action
	nodes     *actionNodes
	completer *readline.PrefixCompleter
}

func newGlobalCommands() *terminalCommands {
	acts := []*action{
		setConfigureModeAction(), // configure terminal
		showInterfacesAction(),   // show interfaces
		showIPRouteAction(),      // show ip route
		showClockAction(),        // show clock
		exitAction(),             // exit
	}

	commands := &terminalCommands{}
	commands.actions = acts
	commands.nodes = actionsToActionNodes(acts)
	commands.completer = createPrefixCompleter(commands.nodes)
	return commands
}

func newConfigureCommands() *terminalCommands {
	acts := []*action{
		setConfigureTerminalModeAction(), // interface eth1
		setHostnameAction(),              // hostname example
		endAction(),                      // end
	}

	commands := &terminalCommands{}
	commands.actions = acts
	commands.nodes = actionsToActionNodes(acts)
	commands.completer = createPrefixCompleter(commands.nodes)
	return commands
}

func newConfigureInterfaceCommands() *terminalCommands {
	acts := []*action{
		setConfigureTerminalModeAction(), // interface eth1
		interfaceShutdownAction(),        // shutdown
		interfaceNoShutdownAction(),      // no shutdown
		interfaceSetAddressAction(),      // ip address ...
		endAction(),                      // end
	}

	commands := &terminalCommands{}
	commands.actions = acts
	commands.nodes = actionsToActionNodes(acts)
	commands.completer = createPrefixCompleter(commands.nodes)
	return commands
}
