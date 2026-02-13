package main

import (
	"strings"
)

type Coord struct {
	row uint8
	col uint8
}

type Board struct {
	size  uint8 // board accepted size: 0 - 256
	cells [][]Cell
	ant   Ant
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
	ant := Ant{
		Coord{10, 10},
		AntDirection(0),
	}
	for row := range size {
		cells[row] = make([]Cell, size)
		for col := range size {
			coord := Coord{row, col}
			cells[row][col] = Cell{coord, false}
		}
	}
	b := &Board{size, cells, ant}
	return b
}
