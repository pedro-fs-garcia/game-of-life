package main

import (
	"slices"
	"strings"
)

type Grid struct {
	size  int
	cells []Cell
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
	for i := range g.cells {
		g.cells[i].Tick()
	}
}

func (g *Grid) CountLivingNeighbors(idx int) int {
	l := len(g.cells)
	s := g.size
	ni := []int{
		(idx - s - 1) % l, (idx - s) % l, (idx + 1) % l,
		(idx - 1) % l, (idx + 1) % l,
		(idx + s - 1) % l, (idx + s) % l, (idx + 1) % l,
	}
	living := 0
	for i := range ni {
		if g.cells[ni[i]].state == ON {
			living++
		}
	}
	return living
}

func (g *Grid) setNextState() {
	for i := range g.cells {
		switch g.cells[i].state {
		case ON, OFF:
			g.cells[i].rollNextState()
		default:
			if g.CountLivingNeighbors(i) == 2 {
				g.cells[i].rollNextState()
			}
		}
	}
}

func NewGrid(size int, livingIndexes []int) *Grid {
	cells := make([]Cell, size*size)
	for i := range len(cells) {
		if slices.Contains(livingIndexes, i) {
			cells[i] = Cell{ON, OFF}
			continue
		}
		cells[i] = Cell{OFF, OFF}
	}
	return &Grid{size, cells}
}
