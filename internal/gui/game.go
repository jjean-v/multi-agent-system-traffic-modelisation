package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"project/internal/graph"
)

const (
	ScreenW          = 1280
	ScreenH          = 720
	LabelCreativeOn  = "Creative Mode : ON"
	LabelCreativeOff = "Creative Mode : OFF"
)

var (
	buttonColor      = color.RGBA{70, 70, 70, 255}    //gris foncé
	buttonColorHover = color.RGBA{100, 100, 100, 255} //gris clair
	backgroundColor  = color.RGBA{40, 40, 40, 255}    //gris anthracite
	nodeColor        = color.RGBA{70, 130, 220, 255}  // bleu
	arcColor         = color.RGBA{255, 100, 100, 255} //rouge
)

type Game struct {
	graph          *graph.Graph
	creativeButton *Button
	creativeMode   bool
}

func NewGame(gr *graph.Graph) *Game {
	g := &Game{graph: gr}
	g.creativeButton = NewButton(10, 10, 200, 30, LabelCreativeOff, func() {
		g.setCreativeMode()
	}, buttonColor, buttonColorHover)
	return g
}

func (g *Game) setCreativeMode() {
	g.creativeMode = !g.creativeMode
	if g.creativeMode {
		g.creativeButton.SetLabel(LabelCreativeOn)
	} else {
		g.creativeButton.SetLabel(LabelCreativeOff)
	}
}

func (g *Game) Update() error {
	overButton := g.creativeButton.Update()
	if !overButton {
		g.HandleEditInput()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	drawArcs(screen, g.graph.GetArcs(), arcColor)
	drawNodes(screen, g.graph.GetNodes(), nodeColor)
	g.creativeButton.Draw(screen)
}

func (g *Game) Layout(W, H int) (int, int) {
	return ScreenW, ScreenH
}
