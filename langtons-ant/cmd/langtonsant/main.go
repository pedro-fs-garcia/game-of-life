package main

import (
	"fmt"
	"time"
)

func main() {
	b := NewBoard(20)
	fmt.Println(b)
	for {
		time.Sleep(50 * time.Millisecond)
		b.RunRound()
		fmt.Println(b)
	}
}
