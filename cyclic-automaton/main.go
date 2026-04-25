package main

import (
	"fmt"
	"time"
)

func main() {
	g := NewGrid(20, 8, 1, 2)
	fmt.Println(g.String())
	for {
		time.Sleep(time.Second)
		g.Tick()
		fmt.Println("\033[H")
		fmt.Println(g.String())
	}
}
