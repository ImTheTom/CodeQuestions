package main

import (
	"fmt"
	"time"
)

type Node struct {
	Value string
}

type Graph struct {
	Nodes []*Node
	Edges map[Node][]*Node
}

func (g *Graph) Construct() {
	if g.Edges == nil {
		g.Edges = make(map[Node][]*Node)
	}
}

func (g *Graph) AddNode(n *Node) {
	g.Nodes = append(g.Nodes, n)
}

func (g *Graph) AddEdge(node1, node2 *Node) {
	g.Edges[*node1] = append(g.Edges[*node1], node2)
	g.Edges[*node2] = append(g.Edges[*node2], node1)
}

func (n *Node) String() string {
	return fmt.Sprintf("%v", n.Value)
}

func (g *Graph) String() {
	s := ""
	for i := 0; i < len(g.Nodes); i++ {
		s += g.Nodes[i].String() + " -> "
		near := g.Edges[*g.Nodes[i]]
		for j := 0; j < len(near); j++ {
			s += near[j].String() + " "
		}
		s += "\n"
	}
	fmt.Println(s)
}

func main() {
	start := time.Now()

	fmt.Println("Running main")

	g := &Graph{}
	g.Construct()

	nA := Node{"A"}
	nB := Node{"B"}
	nC := Node{"C"}
	nD := Node{"D"}
	nE := Node{"E"}
	nF := Node{"F"}
	g.AddNode(&nA)
	g.AddNode(&nB)
	g.AddNode(&nC)
	g.AddNode(&nD)
	g.AddNode(&nE)
	g.AddNode(&nF)

	g.AddEdge(&nA, &nB)
	g.AddEdge(&nA, &nC)
	g.AddEdge(&nB, &nE)
	g.AddEdge(&nC, &nE)
	g.AddEdge(&nE, &nF)
	g.AddEdge(&nD, &nA)

	g.String()

	elapsed := time.Since(start)
	fmt.Printf("Total execution time is %s\n", elapsed)
}
