package main

import (
	"fmt"
	"time"
)

func livingIndexes() []int {
	idxs := make([]int, 40)
	for i := range 20 {
		for j := range 20 {
			if i == j || i+j == 19 {
				idxs = append(idxs, 20*i+j)
			}
		}
	}
	return idxs
}

func main() {
	lc := livingIndexes()
	grid := NewGrid(20, lc)
	for {
		fmt.Println(grid)
		time.Sleep(1 * time.Second)
		grid.setNextState()
		grid.Tick()
	}

}
