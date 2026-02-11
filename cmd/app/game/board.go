package game

import (
	"strings"
)

type Board struct {
	size        int
	cells       [][]Cell
	livingCount int
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
	for i := range b.size {
		for j := range b.size {
			current := &b.cells[i][j]
			current.rollRound()
		}
	}
}

func (b *Board) Round() (*Board, int) {
	changes := 0
	livingCount := 0
	for i := range b.size {
		for j := range b.size {
			current := &b.cells[i][j]
			willLive := current.shouldLive(b)
			if willLive {
				current.makeAlive()
				livingCount++
			} else {
				current.kill()
			}
			if current.isAlive != current.willLive {
				changes++
			}
		}
	}
	b.rollCellRounds()
	b.livingCount = livingCount
	return b, changes
}

func (b *Board) Stop(changes int) bool {
	return b.livingCount == 0 || changes == 0
}

func NewBoard(size int, livingCoords map[Coord]struct{}) *Board {
	cells := make([][]Cell, size)
	for i := range cells {
		cells[i] = make([]Cell, size)
		for j := range cells[i] {
			c := Coord{i, j}
			if c.belongsTo(livingCoords) {
				cells[i][j] = newCell(i, j, true)
				continue
			}
			cells[i][j] = newCell(i, j, false)
		}
	}
	var board *Board = &Board{size, cells, len(livingCoords)}
	return board
}
