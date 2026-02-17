package main

import (
	"slices"
	"strings"
)

type Grid struct {
	size        int
	cells       []Cell
	livingCells int
}

func (g *Grid) Index(row, col int) int {
	return g.size*row + col
}

func (g *Grid) String() string {
	var sb strings.Builder
	for i := range g.size {
		for j := range g.size {
			idx := g.Index(i, j)
			sb.WriteRune(g.cells[idx].Rune())
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (g *Grid) Tick() {
	l := 0
	for i := range g.cells {
		g.cells[i].Tick()
		if g.cells[i].state == ON {
			l++
		}
	}
	g.livingCells = l
}

func (g *Grid) CountLivingNeighbors(idx int) int {
	size := g.size

	row := idx / size
	col := idx % size

	living := 0

	for dr := -1; dr <= 1; dr++ {
		for dc := -1; dc <= 1; dc++ {

			if dr == 0 && dc == 0 {
				continue
			}

			r := (row + dr + size) % size
			c := (col + dc + size) % size

			nIdx := r*size + c

			if g.cells[nIdx].state == ON {
				living++
			}
		}
	}

	return living
}

func (g *Grid) setNextState() {
	for i := range g.cells {
		if g.cells[i].state == OFF && g.CountLivingNeighbors(i) != 2 {
			continue
		}
		g.cells[i].rollNextState()
	}
}

func NewGrid(size int, livingIndexes []int) *Grid {
	cells := make([]Cell, size*size)
	living := 0
	for i := range len(cells) {
		if slices.Contains(livingIndexes, i) {
			cells[i] = Cell{ON, OFF}
			living++
			continue
		}
		cells[i] = Cell{OFF, OFF}
	}
	return &Grid{size, cells, living}
}
