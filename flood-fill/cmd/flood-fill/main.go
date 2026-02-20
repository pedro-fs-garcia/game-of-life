package main

import (
	"fmt"
)

func main() {
	grid := NewGrid(20)
	fmt.Println(grid)
	origin := [2]uint8{10, 10}
	grid.Flood(RED, GREEN, origin)
	fmt.Println(grid)

}
