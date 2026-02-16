package main

import (
	"fmt"
	"time"
	"wireworld/board"
	"wireworld/cell"
	"wireworld/gates"
)

func setCircuit() *board.Grid {
	grid := board.NewGrid(40)
	diode := gates.NewDiode(cell.Coord{Row: 0, Col: 3}, gates.LEFT_TO_RIGHT)
	orGate := gates.NewOrGate(cell.Coord{Row: 1, Col: 8}, gates.LEFT_TO_RIGHT)
	xorGate := gates.NewXorGate(cell.Coord{Row: 3, Col: 13}, gates.LEFT_TO_RIGHT)
	grid.InsertBattery(cell.Coord{Row: 1, Col: 0})
	grid.InsertGate(diode)
	grid.InsertGate(orGate)
	grid.InsertGate(xorGate)
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
