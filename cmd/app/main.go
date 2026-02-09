package main

import (
	"GameOfLife/cmd/app/game"
	"fmt"
	"time"
)

func main() {
	var b *game.Board = game.NewBoard(20, 20)
	fmt.Println(b)
	for {
		b, changes := b.Round()
		fmt.Println(b)
		if b.Stop(changes) {
			break
		}
		time.Sleep(2 * time.Second)
	}
}
