package main

import (
	"fmt"
	"strings"
)

var Cells []Cell = make([]Cell, 30)

func main() {
	for i := range Cells {
		if i%2 == 0 {
			Cells[i] = Cell{true, false}
		} else {
			Cells[i] = Cell{false, false}
		}
	}
	var sb strings.Builder
	for i := range Cells {
		sb.WriteString(" " + Cells[i].String() + " ")
	}
	fmt.Println(sb.String())
}
