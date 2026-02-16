package gates

import "wireworld/cell"

var batteryPattern = [][]cell.CellState{
	{cell.TAIL, cell.HEAD, cell.WIRE},
	{cell.WIRE, cell.EMPTY, cell.EMPTY},
}

func NewBattery(origin cell.Coord) *Gate {
	return NewGate(origin, LEFT_TO_RIGHT, batteryPattern)
}
