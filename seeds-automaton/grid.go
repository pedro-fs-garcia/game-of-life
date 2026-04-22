package main

import "strings"

type CellState uint8

const (
	DEAD CellState = iota
	ALIVE
)

var StateSymbol = map[CellState]string{
	ALIVE: "■ ",
	DEAD:  "□ ",
}

type Grid struct {
	size  int16
	cells [][]CellState
}

func (g *Grid) String() string {
	var sb strings.Builder
	for i := range g.size {
		for j := range g.size {
			sb.WriteString(StateSymbol[g.cells[i][j]])
			sb.WriteRune(' ')
		}
		sb.WriteRune('\n')
	}
	return sb.String()
}

func (g *Grid) ValidateIndex(row int16, col int16) bool {
	if row < 0 || col < 0 {
		return false
	}
	if row > g.size-1 || col > g.size-1 {
		return false
	}
	return true
}

func (g *Grid) Neighborhood(row int16, col int16) [][2]int16 {
	indexes := [][2]int16{
		{row - 1, col - 1}, {row - 1, col}, {row - 1, col + 1},
		{row, col - 1}, {row, col + 1},
		{row + 1, col - 1}, {row + 1, col}, {row + 1, col + 1},
	}

	var n [][2]int16
	for i := range indexes {
		if g.ValidateIndex(indexes[i][0], indexes[i][1]) {
			n = append(n, indexes[i])
		}
	}

	return n
}

func (g *Grid) CellNextState(row int16, col int16) CellState {
	if g.cells[row][col] == ALIVE {
		return DEAD
	}
	neighbors := g.Neighborhood(row, col)
	sum := 0
	for n := range neighbors {
		i := neighbors[n][0]
		j := neighbors[n][1]
		if g.cells[i][j] == ALIVE {
			sum++
		}
	}
	if sum == 2 {
		return ALIVE
	}
	return DEAD
}

func (g *Grid) GridNextState() [][]CellState {
	cells := make([][]CellState, g.size)
	for i := range g.size {
		cells[i] = make([]CellState, g.size)
		for j := range g.size {
			cells[i][j] = g.CellNextState(i, j)
		}
	}
	return cells
}

func (g *Grid) Tick() {
	nextState := g.GridNextState()
	g.cells = nextState
}

func (g *Grid) IsDead() bool {
	for i := range g.size {
		for j := range g.size {
			if g.cells[i][j] == ALIVE {
				return false
			}
		}
	}
	return true
}

func NewGrid(size int16) *Grid {
	cells := make([][]CellState, size)
	for i := range size {
		cells[i] = make([]CellState, size)
		for j := range size {
			if i == j || i+j == size-1 {
				cells[i][j] = ALIVE
			} else {
				cells[i][j] = DEAD
			}
		}
	}
	return &Grid{size, cells}
}
