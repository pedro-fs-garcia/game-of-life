package gates

import "wireworld/cell"

type Orientation uint8

const (
	LEFT_TO_RIGHT Orientation = iota
	RIGHT_TO_LEFT
	TOP_TO_BOTTOM
	BOTTOM_TO_TOP
)

type Gate struct {
	Origin      cell.Coord
	Orientation Orientation
	Cells       [][]cell.Cell
}

func NewGate(origin cell.Coord, orient Orientation, pattern [][]cell.CellState) *Gate {
	var rows, cols = uint8(len(pattern)), uint8(len(pattern[0]))
	if orient == TOP_TO_BOTTOM || orient == BOTTOM_TO_TOP {
		rows, cols = cols, rows
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
	return &Gate{origin, orient, cells}
}
