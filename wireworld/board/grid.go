package board

import (
	"strings"
	"wireworld/gates"
	"wireworld/cell"
)

type Grid struct {
	rows, cols uint8
	cells      [][]cell.Cell
}

func (g *Grid) String() string {
	var sb strings.Builder
	for i := range g.rows {
		for j := range g.cols {
			sb.WriteRune(g.cells[i][j].Symbol())
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (g *Grid) HeadNeighbors(c *cell.Cell) uint8 {
	headCount := uint8(0)
	n := c.Coord.Neighbors(g.rows, g.cols)
	for i := range n {
		if g.cells[n[i].Row][n[i].Col].State == cell.HEAD {
			headCount++
		}
	}
	return headCount
}

func (g *Grid) SetNextState(c *cell.Cell) {
	switch c.State {
	case cell.WIRE:
		n := g.HeadNeighbors(c)
		if n == 1 || n == 2 {
			c.NextState = cell.HEAD
		} else {
			c.NextState = cell.WIRE
		}
	case cell.HEAD:
		c.NextState = cell.TAIL
	case cell.TAIL:
		c.NextState = cell.WIRE
	default:
		c.NextState = cell.EMPTY
	}
}

func (g *Grid) NextGeneration() {
	for i := range g.rows {
		for j := range g.cols {
			g.SetNextState(&g.cells[i][j])
		}
	}
}

func (g *Grid) RunClock() {
	for i := range g.rows {
		for j := range g.cols {
			g.cells[i][j].RunClock()
		}
	}
}

func NewCircuit() []cell.Cell {
	c := make([]cell.Cell, 20)
	for i := range 20 {
		switch i {
		case 1:
			c[i] = cell.NewCell(10, uint8(i), cell.TAIL)
		case 2:
			c[i] = cell.NewCell(10, uint8(i), cell.HEAD)
		default:
			c[i] = cell.NewCell(10, uint8(i), cell.WIRE)
		}
	}
	return c
}

func NewGrid(size uint8) *Grid {
	cells := make([][]cell.Cell, size)
	for i := range size {
		if i == 10 {
			cells[i] = NewCircuit()
			continue
		}
		row := make([]cell.Cell, size)
		for j := range size {
			row[j] = cell.NewCell(uint8(i), uint8(j), cell.EMPTY)
		}
		cells[i] = row
	}
	return &Grid{size, size, cells}
}

func NewEmptyGrid(size uint8) *Grid {
	cells := make([][]cell.Cell, size)
	for i := range size {
		row := make([]cell.Cell, size)
		for j := range size {
			row[j] = cell.NewCell(uint8(i), uint8(j), cell.EMPTY)
		}
		cells[i] = row
	}
	return &Grid{size, size, cells}
}

func (g *Grid) InsertDiode(origin cell.Coord, orient gates.Orientation) {
	diode := gates.NewDiode(origin, orient)
	for i := range len(diode.Cells) {
		for j := range len(diode.Cells[i]) {
			c := diode.Cells[i][j]
			g.cells[c.Coord.Row][c.Coord.Col] = diode.Cells[i][j]
		}
	}
}
