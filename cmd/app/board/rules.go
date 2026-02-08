package board

func LivingCell(b *Board, c *Cell) bool {
	indexes := c.NeighborIndexes(b.height, b.width)

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

func Stop(b *Board) bool {
	return len(b.livingCells) == 0
}
