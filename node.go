package main

import "fmt"

type Node struct {
	Row int
	Col int
}

func (node Node) String() string {
	return fmt.Sprintf("\\Node{%d %d}", node.Row, node.Col)
}

func (node Node) Move(dir Dir) Node {
	switch dir {
	case DIR_RIGHT:
		return Node{node.Row, node.Col + 1}
	case DIR_LEFT:
		return Node{node.Row, node.Col - 1}
	case DIR_UP:
		return Node{node.Row - 1, node.Col}
	case DIR_DOWN:
		return Node{node.Row + 1, node.Row}
	default:
		panic(dir) // unreachable
	}
}
