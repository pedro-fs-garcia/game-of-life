package main

type CellState uint8

const (
	OFF CellState = iota
	ON
	DYING
)

var stateSymbolMap = map[CellState]rune{
	ON:    '■',
	OFF:   ' ',
	DYING: '_',
}

type Cell struct {
	state     CellState
	nextState CellState
}

func (c *Cell) Rune() rune {
	return stateSymbolMap[c.state]
}

func (c *Cell) Tick() {
	c.state = c.nextState
}

func (c *Cell) rollNextState() {
	c.nextState = (c.state + 1) % 3
}
