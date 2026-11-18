package main

import (
	"project/internal/graph"
)

func main() {
	vec1 := graph.NewVec2f(3.4, 5.8)
	n1 := graph.NewNode(0, vec1, 3)

	vec2 := graph.NewVec2f(-4.6, 9.1)
	n2 := graph.NewNode(1, vec2, 4)

	rule1 := graph.NewRule(30, graph.Yield)
	arc1 := graph.NewArc(100, &rule1, n1, n2)

	n1.AddArc(arc1)

	n1.Print()
	n2.Print()

}
