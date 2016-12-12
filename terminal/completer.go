package terminal

import (
	"github.com/chzyer/readline"
)

func createPrefixCompleter(nodes *actionNodes) *readline.PrefixCompleter {
	completers := []readline.PrefixCompleterInterface{}
	for _, n := range *nodes {
		completers = append(completers, actionNodeToPrefixCompleter(n))
	}

	return readline.NewPrefixCompleter(completers...)
}

func actionNodeToPrefixCompleter(n *actionNode) *readline.PrefixCompleter {
	if len(*n.children) == 0 {
		return readline.PcItem(n.command)
	}

	children := make([]readline.PrefixCompleterInterface, len(*n.children))
	for i, childNode := range *n.children {
		children[i] = actionNodeToPrefixCompleter(childNode)
	}
	return readline.PcItem(n.command, children...)
}
