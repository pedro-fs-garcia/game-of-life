package main

import (
	"fmt"
	"math/rand/v2"
	"strings"
)

type Cell struct {
	state     int
	nextState int
}

func hsvToRGB(h, s, v float64) (uint8, uint8, uint8) {
	i := int(h * 6)
	f := h*6 - float64(i)
	p := v * (1 - s)
	q := v * (1 - f*s)
	t := v * (1 - (1-f)*s)

	var r, g, b float64

	switch i % 6 {
	case 0:
		r, g, b = v, t, p
	case 1:
		r, g, b = q, v, p
	case 2:
		r, g, b = p, v, t
	case 3:
		r, g, b = p, q, v
	case 4:
		r, g, b = t, p, v
	case 5:
		r, g, b = v, p, q
	}

	return uint8(r * 255), uint8(g * 255), uint8(b * 255)
}

func (c *Cell) String(nStates int) string {
	h := float64(c.state) / float64(nStates)
	r, g, b := hsvToRGB(h, 1.0, 1.0)
	return fmt.Sprintf("\033[38;2;%d;%d;%dm●\033[0m", r, g, b)
}

type Grid struct {
	size      int
	threshold int
	states    int
	radius    int
	cells     [][]Cell
}

func (g *Grid) NeighborsSum(row int, col int) int {
	size := int(g.size)
	radius := int(g.radius)
	targetState := (g.cells[row][col].state + 1) % g.states
	sum := 0
	for r := -radius; r <= radius; r++ {
		neighborRow := row + r
		if neighborRow < 0 || neighborRow >= size {
			continue
		}

		for c := -radius; c <= radius; c++ {
			if r == 0 && c == 0 {
				continue
			}

			neighborCol := col + c
			if neighborCol < 0 || neighborCol >= size {
				continue
			}
			neighbor := g.cells[neighborRow][neighborCol]
			if neighbor.state == targetState {
				sum++
			}
		}
	}
	return sum
}

func (g *Grid) SetNextState() {
	for i := range g.size {
		for j := range g.size {
			cell := g.cells[i][j]
			target := (cell.state + 1) % g.states
			s := g.NeighborsSum(int(i), int(j))

			if s >= int(g.threshold) {
				g.cells[i][j].nextState = target
			} else {
				g.cells[i][j].nextState = cell.state
			}
		}
	}
}

func (g *Grid) Tick() {
	g.SetNextState()
	for i := range g.size {
		for j := range g.size {
			g.cells[i][j].state = g.cells[i][j].nextState
		}
	}
}

func (g *Grid) String() string {
	var sb strings.Builder
	for i := range g.size {
		for j := range g.size {
			sb.WriteString(g.cells[i][j].String(g.states) + " ")
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func NewGrid(size int, nStates, threshold, radius int) *Grid {
	cells := make([][]Cell, size)
	for i := range size {
		cells[i] = make([]Cell, size)
		for j := range size {
			s := rand.IntN(int(nStates + 1))
			cells[i][j] = Cell{
				state:     s,
				nextState: s,
			}
		}
	}
	return &Grid{
		size:      size,
		cells:     cells,
		states:    nStates,
		threshold: threshold,
		radius:    radius,
	}
}
