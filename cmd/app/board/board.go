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

func (b *Board) rollCellRounds() {
	for i := range b.height {
		for j := range b.width {
			current := b.cells[i][j]
			current.rollRound()
		}
	}
}

func Round(b *Board) (*Board, int) {
	livingCells := make(map[*Cell]bool)
	changes := 0
	for i := range b.height {
		for j := range b.width {
			current := b.cells[i][j]
			willLive := LivingCell(b, current)
			if willLive {
				current.Born()
				livingCells[current] = willLive
			} else {
				current.Kill()
			}
			if current.isAlive != current.willLive {
				changes++
			}
		}
	}
	b.livingCells = livingCells
	b.rollCellRounds()
	return b, changes
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
