package gui

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"

	"project/internal/graph"
)

func drawNodes(screen *ebiten.Image, nodes []*graph.Node) {
	for _, n := range nodes {
		drawNode(screen, n)
	}
}

func drawNode(screen *ebiten.Image, n *graph.Node) {
	pos := n.GetPose()
	x, y := pos.GetX(), pos.GetY()
	r := float32(n.GetRadius())
	arcs := n.GetArcs()
	for _, a := range arcs {
		drawArc(screen, a, n, a.GetNextNode())
	}
	vector.FillCircle(screen, x, y, r, color.RGBA{70, 130, 220, 255}, true)
	ebitenutil.DebugPrintAt(screen, fmt.Sprint(n.GetId()), int(x)-3, int(y)-8)

}

func drawArc(screen *ebiten.Image, a *graph.Arc, currentNode *graph.Node, nextNode *graph.Node) {
	posStart := currentNode.GetPose()
	xStart, yStart := posStart.GetX(), posStart.GetY()
	posEnd := nextNode.GetPose()
	xEnd, yEnd := posEnd.GetX(), posEnd.GetY()
	vector.StrokeLine(screen, xStart, yStart, xEnd, yEnd, 3, color.RGBA{255, 100, 100, 255}, true)
}
