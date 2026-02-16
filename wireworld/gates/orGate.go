package gates

import "wireworld/cell"

var orPattern = [][]cell.CellState{
	{cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY, cell.EMPTY},
	{cell.EMPTY, cell.EMPTY, cell.WIRE, cell.EMPTY, cell.EMPTY},
	{cell.EMPTY, cell.WIRE, cell.WIRE, cell.WIRE, cell.WIRE},
	{cell.EMPTY, cell.EMPTY, cell.WIRE, cell.EMPTY, cell.EMPTY},
	{cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY, cell.EMPTY},
}

func NewOrGate(origin cell.Coord, orient Orientation) *Gate {
	return NewGate(origin, orient, orPattern)
}
