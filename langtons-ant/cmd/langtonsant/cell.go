package main

import "strings"

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
	var sb strings.Builder
	if c.state {
		sb.WriteString(whiteSqr)
	} else {
		sb.WriteString(blackSqr)
	}
	return sb.String()
}
