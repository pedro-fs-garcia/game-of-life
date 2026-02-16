package gates

import "wireworld/cell"

type Orientation uint8

const (
	LEFT_TO_RIGHT Orientation = iota
	RIGHT_TO_LEFT
	TOP_TO_BOTTOM
	BOTTOM_TO_TOP
)

type Diode struct {
	Origin      cell.Coord
	Orientation Orientation
	Cells       [][]cell.Cell
}

func NewDiode(origin cell.Coord, orient Orientation) *Diode {
	pattern := [3][5]cell.CellState{
		{cell.EMPTY, cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY},
		{cell.WIRE, cell.WIRE, cell.EMPTY, cell.WIRE, cell.WIRE},
		{cell.EMPTY, cell.WIRE, cell.WIRE, cell.EMPTY, cell.EMPTY},
	}

	var rows, cols uint8
	switch orient {
	case LEFT_TO_RIGHT, RIGHT_TO_LEFT:
		rows, cols = 3, 5
	case TOP_TO_BOTTOM, BOTTOM_TO_TOP:
		rows, cols = 5, 3
	}

	cells := make([][]cell.Cell, rows)
	for r := uint8(0); r < rows; r++ {
		cells[r] = make([]cell.Cell, cols)
		for c := uint8(0); c < cols; c++ {
			var sr, sc uint8
			switch orient {
			case LEFT_TO_RIGHT:
				sr, sc = r, c
			case RIGHT_TO_LEFT:
				sr, sc = r, 4-c
			case TOP_TO_BOTTOM:
				sr, sc = c, r
			case BOTTOM_TO_TOP:
				sr, sc = 2-c, r
			}
			coord := cell.Coord{Row: origin.Row + r, Col: origin.Col + c}
			cells[r][c] = cell.Cell{Coord: coord, State: pattern[sr][sc], NextState: pattern[sr][sc]}
		}
	}
	return &Diode{origin, orient, cells}
}
