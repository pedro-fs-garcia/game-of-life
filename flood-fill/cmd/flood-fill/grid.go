package main

import (
	"slices"
	"strings"
)

type Color uint8

const (
	RED Color = iota
	GREEN
	BLUE
	WALL
)

var colorSymbol = map[Color]string{
	RED:   "\033[38;2;255;0;0m●\033[0m",
	GREEN: "\033[38;2;0;255;0m●\033[0m",
	BLUE:  "\033[38;2;0;0;255m●\033[0m",
	WALL:  "\033[38;2;0;0;0m●\033[0m",
}

type Grid struct {
	size  uint8
	cells [][]Color
}

func (g *Grid) String() string {
	var sb strings.Builder
	for i := range g.size {
		for j := range g.size {
			sb.WriteString(colorSymbol[g.cells[i][j]])
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (g *Grid) Neighbors(row, col int, change Color) [][2]uint8 {
	ni := [][]int{
		{row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1},
		{row, col - 1}, {row, col + 1},
		{row + 1, col - 1}, {row + 1, col}, {row + 1, col + 1},
	}
	n := [][2]uint8{}
	for i := range ni {
		if ni[i][0] < 0 || ni[i][0] >= int(g.size) || ni[i][1] < 0 || ni[i][1] >= int(g.size) {
			continue
		}
		t := [2]uint8{uint8(ni[i][0]), uint8(ni[i][1])}
		if g.cells[t[0]][t[1]] == change {
			n = append(n, t)
		}
	}
	return n
}

func (g *Grid) RecursiveFlood(change, target Color, origin [2]uint8) {
	if g.cells[origin[0]][origin[1]] == change {
		g.cells[origin[0]][origin[1]] = target
	} else {
		return
	}
	for _, n := range g.Neighbors(int(origin[0]), int(origin[1]), change) {
		g.RecursiveFlood(change, target, n)
	}
}

func (g *Grid) IterativeFlood(change, target Color, origin [2]uint8) {
	q := [][2]uint8{origin}

	for len(q) > 0 {
		coord := q[0]
		if g.cells[coord[0]][coord[1]] != change {
			q = q[1:]
			continue
		}
		g.cells[coord[0]][coord[1]] = target

		for _, n := range g.Neighbors(int(coord[0]), int(coord[1]), change) {
			if !slices.Contains(q, n) {
				q = append(q, n)
			}
		}

		q = q[1:]
	}

}

func NewGrid(size uint8, initialCells map[Color][][2]uint8) *Grid {
	cells := make([][]Color, size)
	for i := range size {
		cells[i] = make([]Color, size)
		for j := range size {
			cells[i][j] = WALL
		}
	}
	for color, coords := range initialCells {
		for _, c := range coords {
			cells[c[0]][c[1]] = color
		}
	}
	return &Grid{size, cells}
}
