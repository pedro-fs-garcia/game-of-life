package main

import (
	"strings"
)

type Strip struct {
	cells  []Cell
	length int
	rule   []string
}

func (s *Strip) RollRound() {
	for i := range s.cells {
		s.cells[i].RollRound()
	}
}

func (s *Strip) NewGeneration() {
	for i := range s.cells {
		pattern := s.cells[i].GetPattern(
			s.cells[(i+s.length-1)%s.length],
			s.cells[(i+s.length+1)%s.length])
		s.cells[i].CalculateNextState(pattern, s.rule)
	}
}

func (s *Strip) String() string {
	var sb strings.Builder
	for i := range s.cells {
		sb.WriteString(" " + s.cells[i].String() + " ")
	}
	return sb.String()
}

func NewStrip(length int, rule []string) *Strip {
	cells := make([]Cell, length)
	for i := range cells {
		cells[i] = Cell{false, true}
	}
	cells[length/2].currentState = true
	return &Strip{cells, length, rule}
}
