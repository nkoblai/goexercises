package main

import (
	"fmt"
)

type Graph struct {
	nodes []*node
}

type node struct {
	edges     []edge
	index     int
	container Node
}

type Node struct {
	node  *node
	Value *interface{}
}

type edge struct {
	weight int
	end    *node
}

type Edge struct {
	Weight int
	Start  Node
	End    Node
}

// MakeNode creates a node, adds it to the graph and returns the new node.
func (g *Graph) MakeNode() Node {
	newNode := &node{index: len(g.nodes)}
	newNode.container = Node{node: newNode, Value: new(interface{})}
	g.nodes = append(g.nodes, newNode)
	return newNode.container
}

func main() {
	graph := &Graph{}
	nodes := make(map[rune]Node, 0)
	nodes['a'] = graph.MakeNode()
	nodes['b'] = graph.MakeNode()
	nodes['c'] = graph.MakeNode()
	nodes['a'] = graph.MakeNode()
	fmt.Println(nodes)
}
