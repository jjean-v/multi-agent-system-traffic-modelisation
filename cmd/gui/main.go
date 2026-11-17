package main

import (
	"log"
	"project/internal/graph"
	"project/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	n1 := graph.NewNode(1, graph.NewVec2f(500, 360), 20)
	n2 := graph.NewNode(2, graph.NewVec2f(780, 360), 20)
	rule1 := graph.NewRule(30, graph.PriorityRoad)
	a1 := graph.NewArc(1, *rule1, &n2)
	n1.AddArc(a1)
	nodes := []*graph.Node{&n1, &n2}

	game := gui.NewGame(nodes)

	ebiten.SetWindowSize(gui.ScreenW, gui.ScreenH)
	ebiten.SetWindowTitle("IA04 simulation rond point BF")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
