package main

import "fmt"

type Graph struct {
	graph []*Node
}

type Node struct {
	Value   int
	Edges   []*Node
	Visited bool
}

func NewNode(val int) *Node {
	return &Node{
		Value: val,
		Edges: []*Node{},
	}
}

func (n *Node) AddEdge(node *Node) {
	n.Edges = append(n.Edges, node)
}

func NewGraph() *Graph {
	return &Graph{graph: []*Node{}}
}

func (g *Graph) AddNode(n *Node) {
	g.graph = append(g.graph, n)
}

func (g *Graph) AddEdge(n1, n2 *Node) {
	n1.AddEdge(n2)
	n2.AddEdge(n1)
}

func (g *Graph) DFS(n *Node) {
	n.Visited = true
	fmt.Println(n.Value)

	for _, node := range n.Edges {
		if !node.Visited {
			g.DFS(node)
		}
	}
}

func (g *Graph) BFS(queue []*Node) {
	if len(queue) == 0 {
		return
	}

	curr := queue[0]
	queue = queue[1:]
	fmt.Println(curr.Value)

	for _, node := range curr.Edges {
		if !node.Visited {
			node.Visited = true
			queue = append(queue, node)
		}
	}

	g.BFS(queue)
}

func main() {
	n1 := NewNode(1)
	n2 := NewNode(2)
	n3 := NewNode(3)
	n4 := NewNode(4)

	g := NewGraph()
	g.AddNode(n1)
	g.AddNode(n2)
	g.AddNode(n3)
	g.AddNode(n4)

	g.AddEdge(n1, n2)
	g.AddEdge(n1, n3)
	//g.AddEdge(n2, n3)
	//g.AddEdge(n2, n4)
	//g.AddEdge(n3, n4)

	g.DFS(n1)
}
