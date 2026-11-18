package main

import (
	"log"
	"project/internal/graph"
	"project/internal/gui"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	gr := graph.NewGraph()

	game := gui.NewGame(&gr)

	ebiten.SetWindowSize(gui.ScreenW, gui.ScreenH)
	ebiten.SetWindowTitle("IA04 simulation rond point BF")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
