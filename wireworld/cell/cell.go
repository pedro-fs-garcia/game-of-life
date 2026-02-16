package cell

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
	Coord     Coord
	State     CellState
	NextState CellState
}

func (c *Cell) Symbol() rune {
	return stateSymbols[c.State]
}

func WireNextState() CellState {
	return WIRE
}

func (c *Cell) RunClock() {
	c.State = c.NextState
}

func NewCell(row, col uint8, state CellState) Cell {
	return Cell{Coord{Row: row, Col: col}, state, EMPTY}
}
