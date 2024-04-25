package main

import "fmt"

type Node struct {
	Row int
	Col int
}

func (node Node) String() string {
	return fmt.Sprintf("{%d %d}", node.Row, node.Col)
}
