package main

type Coord struct {
	row, col uint8
}

var dirs = [8][2]int8{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func (c *Coord) Neighbors(rows uint8, cols uint8) []Coord {
	var n []Coord
	for d := range dirs {
		row := int8(c.row) + dirs[d][0]
		col := int8(c.col) + dirs[d][1]
		if row < 0 || row >= int8(rows) || col < 0 || col >= int8(cols) {
			continue
		}
		n = append(n, Coord{uint8(row), uint8(col)})
	}
	return n
}
