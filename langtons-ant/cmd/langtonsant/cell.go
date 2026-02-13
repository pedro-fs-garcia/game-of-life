package main

const (
	whiteSqr = "■"
	blackSqr = " "
)

type Cell struct {
	coord Coord
	state bool
}

func (c *Cell) Flip() {
	c.state = !c.state
}

func (c *Cell) String(hasAnt bool) string {
	if hasAnt {
		return ""
	}
	if c.state {
		return whiteSqr
	} else {
		return blackSqr
	}
}
