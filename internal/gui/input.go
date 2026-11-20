package gui

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) HandleEditInput() {
	if g.creativeMode != true {
		return
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {

		mouseX, mouseY := ebiten.CursorPosition()
		node := g.graph.IsThereNode(float32(mouseX), float32(mouseY), 20)
		if node != nil {
			fmt.Println("A node is already present")
		} else {
			g.graph.CreateNode(mouseX, mouseY, 20)

		}
	}
}
