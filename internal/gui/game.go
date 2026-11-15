package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"project/internal/graph"
)

const (
	ScreenW = 1280
	ScreenH = 720
)

type Game struct {
	nodes []*graph.Node
}

func NewGame(nodes []*graph.Node) *Game {
	return &Game{nodes: nodes}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 40, 40, 255})
	drawNodes(screen, g.nodes)
}

func (g *Game) Layout(W, H int) (int, int) {
	return ScreenW, ScreenH
}
