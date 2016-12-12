package terminal

import (
	"io"
	"strings"
)

type handler struct {
	nodes *actionNodes
}

func (h *handler) invoke(commands []string, s *state, out io.Writer) {
	var nodes *actionNodes
	var args []string
	var method actionMethod
	var matchingPrefix bool

	nodes = h.nodes
	for index, command := range commands {
		node := nodes.search(command)

		if node == nil {
			// undefined command
			if matchingPrefix {
				outInvalidCommand(out, commands)
			}
			return
		}

		matchingPrefix = true

		if len(*node.children) > 0 {
			nodes = node.children
			continue
		}

		/////////
		// finalize action method
		/////////
		if node.actionConfig == nil {
			panic("Fatal: action config is not defined.")
		}
		args = commands[(index + 1):]
		if node.actionConfig.IsLesserArguments(args) || node.actionConfig.IsGreaterArguments(args) {
			outInvalidArgs(out, args)
			return
		}

		method = node.actionConfig.method
		break
	}

	if method == nil {
		outInvalidCommand(out, commands)
		return
	}

	method(args, s, out)
}

func outInvalidCommand(out io.Writer, commands []string) {
	_, err := out.Write([]byte("invalid command: " + strings.Join(commands, " ")))
	if err != nil {
		panic(err)
	}
}

func outInvalidArgs(out io.Writer, args []string) {
	_, err := out.Write([]byte("invalid argments: " + strings.Join(args, " ")))
	if err != nil {
		panic(err)
	}
}
