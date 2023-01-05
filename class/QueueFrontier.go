package class

import "errors"

type QueueFrontierClass struct {
	StackFrontierClass
}

func (stack QueueFrontierClass) Remove() NodeClass {
	if stack.Empty() {
		panic(errors.New("empty frontier"))
	} else {
		node := stack.frontier[len(stack.frontier)-1]
		stack.frontier = stack.frontier[:len(stack.frontier)-1]
		return node
	}
}
