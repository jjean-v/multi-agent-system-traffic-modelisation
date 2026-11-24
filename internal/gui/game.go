package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"

	"project/internal/graph"
)

const (
	ScreenW             = 1280
	ScreenH             = 720
	LabelCreativeOn     = "Creative Mode : ON"
	LabelCreativeOff    = "Creative Mode : OFF"
	LabelCreateNode     = "Node Mode"
	LabelCreateArc      = "Arc Mode"
	Radius              = 20
	SpaceBetweenButtons = 40
	ButtonX             = 10
	ButtonCreativeY     = 10
	ButtonCreateY       = ButtonCreativeY + SpaceBetweenButtons
	ButtonW             = 200
	ButtonH             = 30
	NbSettingsButton    = 2
	SettingsMargins     = 5
	SettingsX           = ButtonX - SettingsMargins
	SettingsY           = ButtonCreativeY - SettingsMargins
	SettingsW           = ButtonW + 2*SettingsMargins
	SettingsH           = NbSettingsButton*ButtonH + 2*SettingsMargins
)

var (
	buttonColor       = color.RGBA{70, 70, 70, 255}    //gris foncé
	buttonColorHover  = color.RGBA{100, 100, 100, 255} //gris clair
	backgroundColor   = color.RGBA{40, 40, 40, 255}    //gris anthracite
	nodeColor         = color.RGBA{70, 130, 220, 255}  //bleu
	nodePreviewColor  = color.RGBA{90, 150, 250, 75}   //bleu clair transparent
	nodeSelectedColor = color.RGBA{120, 170, 255, 150} //bleu clair
	arcColor          = color.RGBA{255, 100, 100, 255} //rouge
	arcPreviewColor   = color.RGBA{200, 130, 130, 255} //rouge pale
	hoveredNodeColor  = color.RGBA{200, 80, 200, 150}  //violet
)

type Game struct {
	graph                 *graph.Graph
	creativeButton        *Button
	createNodeorArcButton *Button
	clearGraphButton      *Button
	creativeMode          bool
	createArc             bool
	arcStart              *graph.Node
	nodeSelected          *graph.Node
}

func NewGame(gr *graph.Graph) *Game {
	g := &Game{graph: gr}
	g.creativeButton = NewButton(ButtonX, ButtonCreativeY, ButtonW, ButtonH, LabelCreativeOff, func() { g.setCreativeMode() }, buttonColor, buttonColorHover)
	g.createNodeorArcButton = NewButton(ButtonX, ButtonCreateY, ButtonW, ButtonH, LabelCreateNode, func() { g.setCreateMode() }, buttonColor, buttonColorHover)
	g.creativeButton = NewButton(ButtonX, ButtonCreativeY, ButtonW, ButtonH, LabelCreativeOff, func() { g.setCreativeMode() }, buttonColor, buttonColorHover)
	g.createNodeorArcButton = NewButton(ButtonX, ButtonCreateY, ButtonW, ButtonH, LabelCreateNode, func() { g.setCreateMode() }, buttonColor, buttonColorHover)
	return g
}

func (g *Game) setCreativeMode() {
	g.nodeSelected = nil
	g.arcStart = nil
	g.creativeMode = !g.creativeMode
	if g.creativeMode {
		g.creativeButton.SetLabel(LabelCreativeOn)
	} else {
		g.creativeButton.SetLabel(LabelCreativeOff)
	}
}

func (g *Game) setCreateMode() {
	g.nodeSelected = nil
	g.arcStart = nil
	g.createArc = !g.createArc
	if g.createArc {
		g.createNodeorArcButton.SetLabel(LabelCreateArc)
	} else {
		g.createNodeorArcButton.SetLabel(LabelCreateNode)
	}
}

func (g *Game) isSettingsHovered() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	fmouseX, fmouseY := float32(mouseX), float32(mouseY)
	settingsHovered := fmouseX >= SettingsX && fmouseX <= SettingsX+SettingsW && fmouseY >= SettingsY && fmouseY <= SettingsY+SettingsH
	return settingsHovered
}

func (g *Game) Update() error {
	overCreativeButton := g.creativeButton.Update()
	g.createNodeorArcButton.Update()
	if !overCreativeButton {
		if !g.isSettingsHovered() {
			g.HandleEditInput()
		}
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(backgroundColor)
	drawArcs(screen, g.graph.GetArcs(), arcColor)
	drawNodes(screen, g.graph.GetNodes(), nodeColor)
	if g.creativeMode {
		if g.createArc {
			mouseX, mouseY := ebiten.CursorPosition()
			node := g.graph.IsThereNode(float32(mouseX), float32(mouseY), 0)
			if node != nil {
				hoveredNode(screen, node.GetPose(), float32(node.GetRadius()))
			}
			if g.arcStart != nil {
				drawArcPreview(screen, g.arcStart)
			}
		} else {
			mouseX, mouseY := ebiten.CursorPosition()
			node := g.graph.IsThereNode(float32(mouseX), float32(mouseY), 0)
			if node != nil {
				hoveredNode(screen, node.GetPose(), float32(node.GetRadius()))
			} else if !g.isSettingsHovered() {
				drawNodePreview(screen)
			}

			if g.nodeSelected != nil {
				drawNodeSelected(screen, g.nodeSelected)
			}
		}
		g.createNodeorArcButton.Draw(screen)
	}
	g.creativeButton.Draw(screen)
}

func (g *Game) Layout(W, H int) (int, int) {
	return ScreenW, ScreenH
}
