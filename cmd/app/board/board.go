package board

import (
	"strings"
)

type Cell bool

type Board struct {
	width, height int
	cells         [][]Cell
}

func (b *Board) String() string {
	var sb strings.Builder

	for i := range b.cells {
		for j := range b.cells[i] {
			if b.cells[i][j] {
				sb.WriteString(" ■ ")
			} else {
				sb.WriteString(" □ ")
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func NewBoard(width, height int) *Board {
	cells := make([][]Cell, height)
	for i := range cells {
		cells[i] = make([]Cell, width)
	}
	var board *Board = &Board{width, height, cells}
	return board
}
