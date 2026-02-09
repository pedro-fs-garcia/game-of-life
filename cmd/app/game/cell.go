package game

type Cell struct {
	row      int
	col      int
	isAlive  bool
	willLive bool
}

func (c *Cell) kill() {
	c.willLive = false
}

func (c *Cell) makeAlive() {
	c.willLive = true
}

func (c *Cell) neighborIndexes(height, width int) [8][2]int {
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

func (c *Cell) shouldLive(b *Board) bool {
	indexes := c.neighborIndexes(b.height, b.width)

	neighbors := 0
	for k := range 8 {
		x := indexes[k][0]
		y := indexes[k][1]
		currentCell := b.cells[x][y]
		if _, ok := b.livingCells[currentCell]; ok {
			neighbors++
		}
	}
	if c.isAlive && (neighbors == 2 || neighbors == 3) {
		return true
	} else if !c.isAlive && neighbors == 3 {
		return true
	} else {
		return false
	}
}

func correctIndex(current int, limit int) int {
	if current >= limit {
		current = current % limit
	} else if current < 0 {
		current = limit - 1
	}
	return current
}

func newCell(row, col int, isAlive bool) *Cell {
	newCell := Cell{row, col, isAlive, isAlive}
	return &newCell
}
