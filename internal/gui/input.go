package gui

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) HandleEditInput() {
	if g.creativeMode != true {
		return
	}
	if g.createArc {
		g.HandleArcInput()
	} else {
		g.HandleNodeInput()
	}

}

func (g *Game) HandleArcInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.arcStart = nil
		return
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(ebiten.MouseButtonLeft)) {
		mouseX, mouseY := ebiten.CursorPosition()
		node := g.graph.IsThereNode(float32(mouseX), float32(mouseY), 0)
		if g.arcStart == nil {
			g.arcStart = node
		} else if node == g.arcStart {
			g.arcStart = nil //deselectionne quand on reclique dessus
		} else {
			g.graph.AddConnection(g.arcStart.GetId(), node.GetId())
			g.arcStart = nil
		}
	}
}

func (g *Game) HandleNodeInput() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		g.nodeSelected = nil
		return
	}
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButton(ebiten.MouseButtonLeft)) {
		mouseX, mouseY := ebiten.CursorPosition()
		node := g.graph.IsThereNode(float32(mouseX), float32(mouseY), 20)
		if node != nil {
			g.nodeSelected = node
		} else {
			g.nodeSelected = nil
			g.graph.CreateNode(mouseX, mouseY, 20)
		}
	} else if inpututil.IsKeyJustPressed(ebiten.KeyDelete) && g.nodeSelected != nil {
		g.graph.DeleteNode(g.nodeSelected.GetId())
		g.nodeSelected = nil
	}
}
