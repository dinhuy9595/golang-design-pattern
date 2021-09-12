package main

type ValNode struct {
	val int
}

func (n *ValNode) Interpret() int {
	return n.val
}
