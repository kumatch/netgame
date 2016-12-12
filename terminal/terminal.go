package terminal

import (
	"io"
	"strings"

	"os"

	"github.com/chzyer/readline"
	"github.com/kumatch/netgame/device"
)

type Terminal struct {
	D        chan *device.Device
	instance *readline.Instance
	state    *state
	out      io.Writer
}

func (t *Terminal) setup() {
	instance, err := readline.NewEx(&readline.Config{
		InterruptPrompt:   "^C",
		HistorySearchFold: true,
	})
	if err != nil {
		panic(err)
	}

	state := newState(instance)
	t.instance = instance
	t.state = state
}

func (t *Terminal) close() {
	t.instance.Clean()
}

func (t *Terminal) ClearScreen() {
	os.Stdout.Write([]byte("\033[H\033[2J"))
	t.instance.Refresh()
}

func (t *Terminal) Run() {
	t.setup()
	defer t.close()

	globalCommands := newGlobalCommands()
	configureCommands := newConfigureCommands()
	configureInterfaceCommands := newConfigureInterfaceCommands()

	globalHandler := &handler{nodes: globalCommands.nodes}
	configureHandler := &handler{nodes: configureCommands.nodes}
	configureInterfaceHandler := &handler{nodes: configureInterfaceCommands.nodes}

	go func() {
		for {
			select {
			case d := <-t.D:
				t.state.reset(d)
				t.ClearScreen()
			}
		}
	}()

	for {
		line, err := t.instance.Readline()
		if err == readline.ErrInterrupt {
			if len(line) == 0 {
				break
			} else {
				continue
			}
		} else if err == io.EOF {
			break
		}

		commands := strings.Fields(strings.TrimSpace(line))
		if len(commands) == 0 {
			continue
		}

		switch {
		case t.state.isGlobalMode():
			globalHandler.invoke(commands, t.state, t.out)
			if t.state.fin {
				return
			}
		case t.state.isConfigureMode():
			configureHandler.invoke(commands, t.state, t.out)
		case t.state.isConfigureInterfaceMode():
			configureInterfaceHandler.invoke(commands, t.state, t.out)
		}
	}
}

func NewTerminal() *Terminal {
	return &Terminal{
		D:   make(chan *device.Device),
		out: newOutout(),
	}
}
