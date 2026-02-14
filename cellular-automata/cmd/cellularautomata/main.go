package main

import (
	"fmt"
	"strings"
	"time"
)

func NewRule(n uint8) []string {
	bin := toBinary(n)
	var rb strings.Builder
	padding := 8 - len(bin)
	for range padding {
		rb.WriteString("0")
	}
	rb.WriteString(bin)
	return strings.Split(rb.String(), "")
}

func main() {
	rule := NewRule(30)
	fmt.Println(rule)
	strip := NewStrip(30, rule)
	fmt.Println(strip)

	for {
		time.Sleep(1 * time.Second)
		strip.NewGeneration()
		strip.RollRound()
		fmt.Println(strip)
	}
}
