package main

import (
	"strings"
)

type Strip struct {
	cells  []Cell
	length int
	rule   uint8
}

func (s *Strip) RollRound() {
	for i := range s.cells {
		s.cells[i].RollRound()
	}
}

func (s *Strip) NewGeneration() {
	for i := range s.cells {
		s.cells[i].CalculateNextState(
			&s.cells[(i+s.length-1)%s.length],
			&s.cells[(i+s.length+1)%s.length],
			s.rule)
	}
}

func (s *Strip) String() string {
	var sb strings.Builder
	for i := range s.cells {
		sb.WriteString(" " + s.cells[i].Symbol() + " ")
	}
	return sb.String()
}

func NewStrip(length int, rule uint8) *Strip {
	cells := make([]Cell, length)
	for i := range cells {
		cells[i] = Cell{false, false}
	}
	cells[length/2].currentState = true
	return &Strip{cells, length, rule}
}
