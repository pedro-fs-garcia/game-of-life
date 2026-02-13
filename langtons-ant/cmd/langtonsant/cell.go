package main

type Cell struct {
	coord Coord
	state bool
}

func (c *Cell) Flip() {
	c.state = !c.state
}
