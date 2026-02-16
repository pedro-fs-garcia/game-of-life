package main

import (
	"fmt"
	"time"
	"wireworld/board"
	"wireworld/cell"
	"wireworld/gates"
)

func setCircuit() *board.Grid {
	grid := board.NewGrid(20)
	grid.InsertDiode(cell.Coord{Row: 9, Col: 6}, gates.LEFT_TO_RIGHT)
	return grid
}

func StartCircuit(grid *board.Grid) {
	for t := 0; t < 100; t++ {
		fmt.Printf("t = %d\n", t)
		fmt.Println(grid)
		time.Sleep(1 * time.Second)
		grid.NextGeneration()
		grid.RunClock()
	}
}

func main() {
	grid := setCircuit()
	StartCircuit(grid)
}
