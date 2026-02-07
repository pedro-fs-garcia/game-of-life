package main

import (
	"GameOfLife/cmd/app/board"
	"fmt"
)

func main() {
	var board *board.Board = board.NewBoard(20, 20)
	fmt.Println(board)
}
