package gates

import "wireworld/cell"

var diodePattern = [][]cell.CellState{
	{cell.EMPTY, cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY},
	{cell.WIRE, cell.WIRE, cell.EMPTY, cell.WIRE, cell.WIRE},
	{cell.EMPTY, cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY},
}

func NewDiode(origin cell.Coord, orient Orientation) *Gate {
	return NewGate(origin, orient, diodePattern)
}
