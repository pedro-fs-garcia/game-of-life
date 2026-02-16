package gates

import "wireworld/cell"

var xorPattern = [][]cell.CellState{
	{cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY, cell.EMPTY, cell.EMPTY, cell.EMPTY},
	{cell.EMPTY, cell.EMPTY, cell.WIRE, cell.EMPTY, cell.EMPTY, cell.EMPTY, cell.EMPTY},
	{cell.EMPTY, cell.WIRE, cell.WIRE, cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY},
	{cell.EMPTY, cell.WIRE, cell.EMPTY, cell.EMPTY, cell.WIRE, cell.WIRE, cell.WIRE},
	{cell.EMPTY, cell.WIRE, cell.WIRE, cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY},
	{cell.EMPTY, cell.EMPTY, cell.WIRE, cell.EMPTY, cell.EMPTY, cell.EMPTY, cell.EMPTY},
	{cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY, cell.EMPTY, cell.EMPTY, cell.EMPTY},
}

func NewXorGate(origin cell.Coord, orient Orientation) *Gate {
	return NewGate(origin, orient, xorPattern)
}
