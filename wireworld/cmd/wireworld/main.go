package main

import (
	"fmt"
	"time"
)

func main() {
	grid := NewGrid(20)
	c := Coord{grid.rows - 21, 10}
	fmt.Println(c)
	for t := 0; t < 100; t++ {
		fmt.Printf("%d\n", t)
		fmt.Println(grid)
		time.Sleep(1 * time.Second)
		grid.NextGeneration()
		grid.RunClock()
	}
}
