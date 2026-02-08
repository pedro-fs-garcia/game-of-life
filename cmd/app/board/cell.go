package board

type Cell struct {
	row      int
	col      int
	isAlive  bool
	willLive bool
}

func (c *Cell) Kill() {
	c.willLive = false
}

func (c *Cell) Born() {
	c.willLive = true
}

func (c *Cell) NeighborIndexes(height, width int) [8][2]int {
	indexes := [8][2]int{
		{c.row - 1, c.col - 1}, {c.row - 1, c.col}, {c.row - 1, c.col + 1},
		{c.row, c.col - 1}, {c.row, c.col + 1},
		{c.row + 1, c.col - 1}, {c.row + 1, c.col}, {c.row + 1, c.col + 1},
	}
	for k := range 8 {
		indexes[k][0] = correctIndex(indexes[k][0], height)
		indexes[k][1] = correctIndex(indexes[k][1], width)
	}
	return indexes
}

func (c *Cell) rollRound() {
	c.isAlive = c.willLive
}

func correctIndex(current int, limit int) int {
	if current >= limit {
		current = current % limit
	} else if current < 0 {
		current = limit - 1
	}
	return current
}

func NewCell(row, col int, isAlive bool) *Cell {
	newCell := Cell{row, col, isAlive, isAlive}
	return &newCell
}
