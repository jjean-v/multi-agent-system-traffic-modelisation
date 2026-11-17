package main

import (
	"log"
	"project/internal/graph"
	"project/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	gr := graph.NewGraph()
	n1 := gr.CreateNode(500, 360, 20)
	n2 := gr.CreateNode(780, 360, 20)
	gr.AddConnection(n1.GetId(), n2.GetId())

	game := gui.NewGame(&gr)

	ebiten.SetWindowSize(gui.ScreenW, gui.ScreenH)
	ebiten.SetWindowTitle("IA04 simulation rond point BF")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
