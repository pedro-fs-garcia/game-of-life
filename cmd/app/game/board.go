package game

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

func (b *Board) Round() (*Board, int) {
	livingCells := make(map[*Cell]bool)
	changes := 0
	for i := range b.height {
		for j := range b.width {
			current := b.cells[i][j]
			willLive := current.shouldLive(b)
			if willLive {
				current.makeAlive()
				livingCells[current] = willLive
			} else {
				current.kill()
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

func (b *Board) Stop(changes int) bool {
	return len(b.livingCells) == 0 || changes == 0
}

func NewBoard(width, height int) *Board {
	cells := make([][]*Cell, height)
	livingCells := make(map[*Cell]bool)
	for i := range cells {
		cells[i] = make([]*Cell, width)
		for j := range cells[i] {
			if i == j || i+j == 19 {
				cells[i][j] = newCell(i, j, true)
				livingCells[cells[i][j]] = true
				continue
			}
			cells[i][j] = newCell(i, j, false)
		}
	}
	var board *Board = &Board{width, height, cells, livingCells}
	return board
}
