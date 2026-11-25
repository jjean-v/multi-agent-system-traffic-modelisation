package main

import (
	"fmt"
	"project/internal/agt"
	"project/internal/graph"
)

func main() {
	vec1 := graph.NewVec2f(3.4, 5.8)
	n1 := graph.NewNode(0, vec1, 3)

	vec2 := graph.NewVec2f(-4.6, 9.1)
	n2 := graph.NewNode(1, vec2, 4)

	rule1 := graph.NewRule(30, graph.Yield, 3)
	arc1 := graph.NewArc(100, &rule1, n1, n2)

	pieton1 := agt.NewPedestrian(1, graph.NewVec2f(3, 4), graph.NewVec2f(1, 1), agt.NewFieldOfVision(3, 35), nil)
	car1 := agt.NewCar(2, graph.NewVec2f(6, 8), graph.NewVec2f(1, 1), agt.NewFieldOfVision(3, 35), nil)

	agents := []agt.Agent{pieton1, car1}

	for _, a := range agents {
		a.Percept()
		a.Deliberate() // appelle la bonne version selon le type
		a.Act()
		fmt.Println(a)

	}

	n1.AddArc(arc1)

	//n1.Print()
	//n2.Print()

}
