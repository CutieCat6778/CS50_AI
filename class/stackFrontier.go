package class

import "errors"

type NodeClass struct {
	state  [2]int
	parent *NodeClass
	action [2]int
}

type StackFrontierClass struct {
	frontier []NodeClass
}

func (stack StackFrontierClass) Add(node NodeClass) {
	stack.frontier = append(stack.frontier, node)
}

func (stack StackFrontierClass) ContainSelf(state [2]int) []NodeClass {
	var nodes []NodeClass
	for _, n := range stack.frontier {
		if state == n.state {
			nodes = append(nodes, n)
		}
	}
	return nodes
}

func (stack StackFrontierClass) Remove() NodeClass {
	if stack.Empty() {
		panic(errors.New("empty frontier"))
	} else {
		node := stack.frontier[len(stack.frontier)-1]
		stack.frontier = stack.frontier[:len(stack.frontier)-1]
		return node
	}
}

func (stack StackFrontierClass) Empty() bool {
	return len(stack.frontier) == 0
}
