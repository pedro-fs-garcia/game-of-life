package main

import (
	"GameOfLife/cmd/app/board"
	"fmt"
	"time"
)

func main() {
	var b *board.Board = board.NewBoard(20, 20)
	fmt.Println(b)
	for {
		b, changes := board.Round(b)
		fmt.Println(b)
		if board.Stop(b, changes) {
			break
		}
		time.Sleep(2 * time.Second)
	}
}
