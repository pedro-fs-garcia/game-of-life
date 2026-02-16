package main

type Orientation uint8

const (
	LEFT_TO_RIGHT Orientation = iota
	RIGHT_TO_LEFT
	TOP_TO_BOTTOM
	BOTTOM_TO_TOP
)

func Diode(origin Coord, orient Orientation) [][]Cell {
	pattern := [3][5]CellState{
		{EMPTY, WIRE, WIRE, EMPTY, EMPTY},
		{WIRE, WIRE, EMPTY, WIRE, WIRE},
		{EMPTY, WIRE, WIRE, EMPTY, EMPTY},
	}

	var rows, cols uint8
	switch orient {
	case LEFT_TO_RIGHT, RIGHT_TO_LEFT:
		rows, cols = 3, 5
	case TOP_TO_BOTTOM, BOTTOM_TO_TOP:
		rows, cols = 5, 3
	}

	diode := make([][]Cell, rows)
	for r := uint8(0); r < rows; r++ {
		diode[r] = make([]Cell, cols)
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
			coord := Coord{origin.row + r, origin.col + c}
			diode[r][c] = Cell{coord, pattern[sr][sc], pattern[sr][sc]}
		}
	}
	return diode
}
