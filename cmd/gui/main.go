package main

import (
	"fmt"
	"log"
	"project/internal/graph"
	"project/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {

	// Create Graph for A* algorithm
	gr := graph.NewGraph()
	n1 := gr.CreateNode(100, 200, 20)
	n2 := gr.CreateNode(250, 200, 20)
	n3 := gr.CreateNode(300, 300, 20)
	n4 := gr.CreateNode(400, 400, 20)
	n5 := gr.CreateNode(500, 400, 20)
	n6 := gr.CreateNode(600, 300, 20)
	n7 := gr.CreateNode(650, 200, 20)
	n8 := gr.CreateNode(800, 200, 20)
	n9 := gr.CreateNode(500, 500, 20)
	n10 := gr.CreateNode(450, 350, 20)

	gr.AddConnection(n1.GetId(), n2.GetId())
	gr.AddConnection(n2.GetId(), n3.GetId())
	gr.AddConnection(n3.GetId(), n4.GetId())
	gr.AddConnection(n4.GetId(), n5.GetId())
	gr.AddConnection(n5.GetId(), n6.GetId())
	gr.AddConnection(n6.GetId(), n7.GetId())
	gr.AddConnection(n7.GetId(), n8.GetId())
	gr.AddConnection(n5.GetId(), n9.GetId())
	gr.AddConnection(n3.GetId(), n10.GetId())
	gr.AddConnection(n10.GetId(), n6.GetId())

	// Create two types of rules
	rulesafe := graph.NewRule(30, graph.PriorityRoad, 0)
	rulecautious := graph.NewRule(30, graph.PriorityRoad, 4)

	// Add rules to each Arc
	for _, arc := range gr.GetArcs() {
		if arc.GetNextNode().GetId() == n10.GetId() {
			arc.SetRule(&rulecautious)
		} else {
			arc.SetRule(&rulesafe)
		}
	}

	// Create Profile
	prudent := &graph.Profile{Lambda: 30}
	direct := &graph.Profile{Lambda: 1}

	// Launch A* algorithm
	chemin1, _, _ := graph.FindPath(n1, n7, prudent)
	chemin2, _, _ := graph.FindPath(n1, n7, direct)

	//Print Result
	fmt.Println("Chemin Safe")
	for _, node := range chemin1 {
		fmt.Print("->", node.GetId())
	}
	fmt.Println()

	fmt.Println("Chemnin dangereux")
	for _, node := range chemin2 {
		fmt.Print("->", node.GetId())
	}

	//Gui Interface
	game := gui.NewGame(&gr)

	ebiten.SetWindowSize(gui.ScreenW, gui.ScreenH)
	ebiten.SetWindowTitle("IA04 simulation rond point BF")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
