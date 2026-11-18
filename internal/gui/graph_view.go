package gui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"project/internal/graph"
)

func drawNodes(screen *ebiten.Image, nodes map[uint]*graph.Node, colorNode color.Color) {
	for _, n := range nodes {
		drawNode(screen, n, colorNode)
	}
}

func drawArcs(screen *ebiten.Image, arcs map[uint]*graph.Arc, colorArc color.Color) {
	for _, a := range arcs {
		drawArc(screen, a, colorArc)
	}
}

func drawNode(screen *ebiten.Image, n *graph.Node, colorNode color.Color) {
	pos := n.GetPose()
	x, y := pos.GetX(), pos.GetY()
	r := float32(n.GetRadius())
	vector.FillCircle(screen, x, y, r, colorNode, true)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint(n.GetId()), int(x)-3, int(y)-8)
}

func drawArc(screen *ebiten.Image, a *graph.Arc, colorArc color.Color) {
	currentNode := a.GetFrom()
	nextNode := a.GetNextNode()
	posStart := currentNode.GetPose()
	xStart, yStart := posStart.GetX(), posStart.GetY()
	posEnd := nextNode.GetPose()
	xEnd, yEnd := posEnd.GetX(), posEnd.GetY()
	vector.StrokeLine(screen, xStart, yStart, xEnd, yEnd, 3, colorArc, true)
}

func drawNodePreview(screen *ebiten.Image) {
	mouseX, mouseY := ebiten.CursorPosition()
	fmouseX, fmouseY := float32(mouseX), float32(mouseY)
	settingsHovered := fmouseX >= SettingsX && fmouseX <= SettingsX+SettingsW && fmouseY >= SettingsY && fmouseY <= SettingsY+SettingsH
	if !settingsHovered {
		vector.FillCircle(screen, fmouseX, fmouseY, Radius, nodePreviewColor, true)
	}
}
