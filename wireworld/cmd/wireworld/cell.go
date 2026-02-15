package main

type CellState uint8

const (
	EMPTY CellState = iota
	WIRE
	HEAD
	TAIL
)

var stateSymbols = map[CellState]rune{
	EMPTY: ' ',
	WIRE:  '·',
	HEAD:  '●',
	TAIL:  '○',
}

type Cell struct {
	coord     Coord
	state     CellState
	nextState CellState
}

func (c *Cell) Symbol() rune {
	return stateSymbols[c.state]
}

func WireNextState() CellState {
	return WIRE
}

func (c *Cell) SetNextState(g *Grid) {
	switch c.state {
	case WIRE:
		n := c.HeadNeighbors(g)
		if n == 1 || n == 2 {
			c.nextState = HEAD
		} else {
			c.nextState = WIRE
		}
	case HEAD:
		c.nextState = TAIL
	case TAIL:
		c.nextState = WIRE
	default:
		c.nextState = EMPTY
	}
}

func (c *Cell) RunClock() {
	c.state = c.nextState
}

func (c *Cell) HeadNeighbors(g *Grid) uint8 {
	headCount := uint8(0)
	n := c.coord.Neighbors(g.rows, g.cols)
	for i := range n {
		if g.cells[n[i].row][n[i].col].state == HEAD {
			headCount++
		}
	}
	return headCount
}

func NewCell(row, col uint8, state CellState) Cell {
	return Cell{Coord{row, col}, state, EMPTY}
}
