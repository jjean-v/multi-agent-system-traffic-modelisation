package utils

import (
	"errors"
	"project/internal/graph"
)

var ErrOutOfGrid = errors.New("cell is outside the grid")

type Grid struct {
	cells     [][]int
	cellCount graph.Vec2i
	cellSize  graph.Vec2i
}

func NewGrid(cellCount, cellSize graph.Vec2i) *Grid {
	return &Grid{
		cells:     make([][]int, cellCount.GetX()*cellCount.GetY()),
		cellCount: cellCount,
		cellSize:  cellSize,
	}
}

func (g *Grid) indexXY(x, y int) (int, bool) {
	if x < 0 || x >= g.cellCount.GetX() || y < 0 || y >= g.cellCount.GetY() {
		return 0, false
	}
	return x*g.cellCount.GetY() + y, true
}

func (g *Grid) index(cell graph.Vec2i) (int, bool) {
	return g.indexXY(cell.GetX(), cell.GetY())
}

func (g *Grid) Clear() {
	for i := range g.cells {
		g.cells[i] = g.cells[i][:0]
	}
}

func (g *Grid) GetValuesAt(cell graph.Vec2i) ([]int, error) {
	i, ok := g.index(cell)
	if !ok {
		return nil, ErrOutOfGrid
	}
	return g.cells[i], nil
}

func (g *Grid) GetCellsAroundMe(cell graph.Vec2i) [][]int {
	cx, cy := cell.GetX(), cell.GetY()
	around := make([][]int, 0, 8)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if dx == 0 && dy == 0 {
				continue // la case elle-même
			}
			if i, ok := g.indexXY(cx+dx, cy+dy); ok {
				around = append(around, g.cells[i])
			}
		}
	}
	return around
}

func (g *Grid) AddTo(cell graph.Vec2i, id int) error {
	i, ok := g.index(cell)
	if !ok {
		return ErrOutOfGrid
	}
	g.cells[i] = append(g.cells[i], id)
	return nil
}
