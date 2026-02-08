package board

import (
	"strings"
)

type Board struct {
	width, height int
	cells         [][]*Cell
	livingCells   map[*Cell]bool
}

func (b *Board) String() string {
	var sb strings.Builder

	for i := range b.cells {
		for j := range b.cells[i] {
			cell := b.cells[i][j]
			if cell.isAlive {
				sb.WriteString("■ ")
			} else {
				sb.WriteString("□ ")
			}
		}
		sb.WriteByte('\n')
	}
	return sb.String()
}

func Round(b *Board) *Board {
	cells := make([][]*Cell, b.height)
	livingCells := make(map[*Cell]bool)

	for i := range cells {
		cells[i] = make([]*Cell, b.width)
		for j := range cells[i] {
			current := b.cells[i][j]
			isAlive := LivingCell(b, current)
			cells[i][j] = NewCell(i, j, isAlive)
			if isAlive {
				livingCells[cells[i][j]] = true
			}
		}
	}
	board := &Board{b.width, b.height, cells, livingCells}
	return board
}

func NewBoard(width, height int) *Board {
	cells := make([][]*Cell, height)
	livingCells := make(map[*Cell]bool)
	for i := range cells {
		cells[i] = make([]*Cell, width)
		for j := range cells[i] {
			if i == j || i+j == 19 {
				cells[i][j] = NewCell(i, j, true)
				livingCells[cells[i][j]] = true
				continue
			}
			cells[i][j] = NewCell(i, j, false)
		}
	}
	var board *Board = &Board{width, height, cells, livingCells}
	return board
}
