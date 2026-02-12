package main

import (
	"strings"
)

type Coord struct {
	x uint8
	y uint8
}

type Ant struct {
	position  Coord
	direction string
}

type Cell struct {
	coord Coord
	state bool
}

type Board struct {
	size  uint8 // board accepted size: 0 - 256
	cells [][]Cell
}

func (b *Board) String() string {
	var sb strings.Builder
	for i := range b.size {
		for j := range b.size {
			s := b.cells[i][j].state
			if s {
				sb.WriteString("■ ")
			} else {
				sb.WriteString("□ ")
			}
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func NewBoard(size uint8) *Board {
	cells := make([][]Cell, size)
	for x := range size {
		cells[x] = make([]Cell, size)
		for y := range size {
			coord := Coord{x, y}
			cells[x][y] = Cell{coord, false}
		}
	}
	b := &Board{size, cells}
	return b
}
