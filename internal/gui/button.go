package gui

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Button struct {
	x, y, w, h float32
	label      string
	onClick    func()
	hovered    bool
	color      color.Color
	colorHover color.Color
}

func NewButton(x, y, w, h float32, label string, onClick func(), color, colorHover color.Color) *Button {
	return &Button{x: x, y: y, w: w, h: h, label: label, onClick: onClick, color: color, colorHover: colorHover}
}

func (b *Button) SetLabel(label string) {
	b.label = label
}

func (b *Button) Update() bool {
	mouseX, mouseY := ebiten.CursorPosition()
	fmouseX, fmouseY := float32(mouseX), float32(mouseY)
	b.hovered = fmouseX >= b.x && fmouseX <= b.x+b.w && fmouseY >= b.y && fmouseY <= b.y+b.h
	if b.hovered && inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		b.onClick()
	}
	return b.hovered
}

func (b *Button) Draw(screen *ebiten.Image) {
	c := b.color
	if b.hovered {
		c = b.colorHover
	}
	vector.FillRect(screen, b.x, b.y, b.w, b.h, c, false)

	tx := int(b.x+b.w/2) - len(b.label)*3
	ty := int(b.y+b.h/2) - 8
	ebitenutil.DebugPrintAt(screen, b.label, tx, ty)
}
