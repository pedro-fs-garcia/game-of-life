package main

import "strings"

type Grid struct {
	rows, cols uint8
	cells      [][]Cell
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

func (g *Grid) NextGeneration() {
	for i := range g.rows {
		for j := range g.cols {
			g.cells[i][j].SetNextState(g)
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

func NewCircuit() []Cell {
	c := make([]Cell, 20)
	for i := range 20 {
		switch i {
		case 1:
			c[i] = NewCell(10, uint8(i), TAIL)
		case 2:
			c[i] = NewCell(10, uint8(i), HEAD)
		default:
			c[i] = NewCell(10, uint8(i), WIRE)
		}
	}
	return c
}

func NewGrid(size uint8) *Grid {
	cells := make([][]Cell, size)
	for i := range size {
		if i == 10 {
			cells[i] = NewCircuit()
			continue
		}
		row := make([]Cell, size)
		for j := range size {
			row[j] = NewCell(uint8(i), uint8(j), EMPTY)
		}
		cells[i] = row
	}
	return &Grid{size, size, cells}
}

func NewEmptyGrid(size uint8) *Grid {
	cells := make([][]Cell, size)
	for i := range size {
		row := make([]Cell, size)
		for j := range size {
			row[j] = NewCell(uint8(i), uint8(j), EMPTY)
		}
		cells[i] = row
	}
	return &Grid{size, size, cells}
}

func (g *Grid) InsertDiode(origin Coord, orient Orientation) {
	diode := Diode(origin, orient)
	for i := range len(diode) {
		for j := range len(diode[i]) {
			c := diode[i][j]
			g.cells[c.coord.row][c.coord.col] = diode[i][j]
		}
	}
}
