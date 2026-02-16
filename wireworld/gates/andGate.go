package gates

import "wireworld/cell"

var andPattern = [][]cell.CellState{}

func NewAndGate(origin cell.Coord, orient Orientation) *Gate {
	return NewGate(origin, orient, andPattern)
}
