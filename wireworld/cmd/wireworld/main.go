package main

import (
	"fmt"
	"time"
)

func setCircuit() *Grid {
	grid := NewGrid(20)
	grid.InsertDiode(Coord{9, 6}, LEFT_TO_RIGHT)
	return grid
}

func StartCircuit(grid *Grid) {
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
