package utils

import (
	"errors"
	"project/internal/graph"
)

var ErrOutOfGrid = errors.New("cell is outside the grid")

type Grid struct {
	cells       [][]int
	cellCount   graph.Vec2i
	cellSize    graph.Vec2f
	cellSizeInv graph.Vec2f
}

func NewGrid(size graph.Vec2f, cellCount graph.Vec2i) (*Grid, error) {
	cellCountInv, err1 := cellCount.ToVec2f().Inv()
	if err1 != nil {
		return nil, err1
	}

	cellSize := size.Mult(cellCountInv)
	cellSizeInv, err2 := cellSize.Inv()
	if err2 != nil {
		return nil, err2
	}

	return &Grid{
		make([][]int, cellCount.GetX()*cellCount.GetY()),
		cellCount,
		cellSize,
		cellSizeInv,
	}, nil
}

func (g *Grid) scale(v graph.Vec2f) graph.Vec2f {
	return v.Mult(g.cellSizeInv)
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

func (g *Grid) GetValuesAt(cell graph.Vec2f) ([]int, error) {
	i, ok := g.index(g.scale(cell).FloorToVec2i())
	if !ok {
		return nil, ErrOutOfGrid
	}
	return g.cells[i], nil
}

func (g *Grid) GetCellsAroundMe(cell graph.Vec2f) [][]int {
	tempCell := g.scale(cell).FloorToVec2i()
	cx, cy := tempCell.GetX(), tempCell.GetY()
	around := make([][]int, 0, 8)

	for dx := -1; dx <= 1; dx++ {
		for dy := -1; dy <= 1; dy++ {
			if i, ok := g.indexXY(cx+dx, cy+dy); ok {
				around = append(around, g.cells[i])
			}
		}
	}
	return around
}

func (g *Grid) AddTo(cell graph.Vec2f, id int) error {
	i, ok := g.index(g.scale(cell).FloorToVec2i())
	if !ok {
		return ErrOutOfGrid
	}
	g.cells[i] = append(g.cells[i], id)
	return nil
}
