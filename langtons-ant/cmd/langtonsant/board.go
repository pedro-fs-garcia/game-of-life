package main

import (
	"strings"
)

type Coord struct {
	row uint8
	col uint8
}

func (c Coord) belongsTo(coordMap map[Coord]struct{}) bool {
	_, ok := coordMap[c]
	return ok
}

type Board struct {
	size  uint8 // board accepted size: 0 - 256
	cells [][]Cell
	ant   Ant
}

func (b *Board) AntPosition() (uint8, uint8) {
	return b.ant.position.row, b.ant.position.col
}

func (b *Board) String() string {
	var sb strings.Builder
	antRow, antCol := b.AntPosition()
	for i := range b.size {
		for j := range b.size {
			var square string
			if i == antRow && j == antCol {
				square = b.cells[i][j].String(true) + b.ant.String()
			} else {
				square = b.cells[i][j].String(false)
			}
			sb.WriteString(square + " ")
		}
		sb.WriteString("\n")
	}
	return sb.String()
}

func (b *Board) RunRound() {
	r, c := b.AntPosition()
	if b.cells[r][c].state {
		b.ant.TurnRight()
	} else {
		b.ant.TurnLeft()
	}
	b.cells[r][c].Flip()
	b.ant.Walk(b.size)
}

func NewBoard(size uint8, whiteCoordinates map[Coord]struct{}) *Board {
	cells := make([][]Cell, size)
	center := size / 2
	ant := Ant{
		Coord{center, center},
		NORTH,
	}
	for row := range size {
		cells[row] = make([]Cell, size)
		for col := range size {
			coord := Coord{row, col}
			if coord.belongsTo(whiteCoordinates) {
				cells[row][col] = Cell{coord, true}
			} else {
				cells[row][col] = Cell{coord, false}
			}
		}
	}
	b := &Board{size, cells, ant}
	return b
}
