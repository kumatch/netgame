package terminal

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type action struct {
	commands []string
	config   *actionConfig
}

type actionConfig struct {
	method       actionMethod
	minArgNumebr int
	maxArgNumber int
}

func (c *actionConfig) IsLesserArguments(args []string) bool {
	return len(args) < c.minArgNumebr
}

func (c *actionConfig) IsGreaterArguments(args []string) bool {
	return len(args) > c.maxArgNumber
}

func (a *action) splitAction() (cmd string, childAction *action) {
	if len(a.commands) == 0 {
		panic("Cannot split action")
	}

	cmd = a.commands[0]
	if len(a.commands) > 1 {
		childAction = &action{
			commands: a.commands[1:],
			config:   a.config,
		}
	}
	return
}

type actionMethod func(args []string, s *state, out io.Writer)

type actionNode struct {
	command      string
	children     *actionNodes
	actionConfig *actionConfig
}

func (n *actionNode) IsSameCommand(cmd string) bool {
	return n.command == cmd
}

func (n *actionNode) IsEnd() bool {
	if n.actionConfig == nil {
		return false
	}
	return n.actionConfig.method != nil
}

func (n *actionNode) addChild(child *action) {
	n.children.add(child)
}

func (n *actionNode) setActionConfig(c *actionConfig) {
	n.actionConfig = c
}

func newActionNode(cmd string) *actionNode {
	return &actionNode{
		command:  cmd,
		children: &actionNodes{},
	}
}

type actionNodes []*actionNode

func (nodes *actionNodes) add(a *action) {
	var isExistsNode bool
	cmd, childAction := a.splitAction()

	node := nodes.search(cmd)
	if node == nil {
		node = newActionNode(cmd)
		*nodes = append(*nodes, node)
		sort.Sort(*nodes)
	} else {
		isExistsNode = true
	}

	if childAction != nil {
		if isExistsNode && node.IsEnd() {
			panic(fmt.Sprintf("Fatal: adding action [%s] to end node.", strings.Join(a.commands, " ")))
		}
		node.addChild(childAction)
	} else {
		if isExistsNode && !node.IsEnd() {
			panic(fmt.Sprintf("Fatal: adding action [%s] to not end node.", strings.Join(a.commands, " ")))
		}
		node.setActionConfig(a.config)
	}
}

func (nodes *actionNodes) search(command string) *actionNode {
	for _, n := range *nodes {
		if n.command == command {
			return n
		}
	}
	return nil
}

func (nodes actionNodes) Len() int {
	return len(nodes)
}

func (nodes actionNodes) Less(i, j int) bool {
	return nodes[i].command < nodes[j].command
}

func (nodes actionNodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}

func actionsToActionNodes(actions []*action) *actionNodes {
	nodes := &actionNodes{}
	for _, a := range actions {
		nodes.add(a)
	}
	return nodes
}
