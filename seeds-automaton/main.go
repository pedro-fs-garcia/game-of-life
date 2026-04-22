package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("==== Seed Automaton ====")
	g := NewGrid(20)
	fmt.Println(g)
	for {
		g.Tick()
		time.Sleep(2 * time.Second)
		fmt.Println(g)
		if g.IsDead() {
			break
		}
	}
}
