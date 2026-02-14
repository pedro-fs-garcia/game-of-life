package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func GetInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func askforRule(reader *bufio.Reader) (uint8, error) {
	fmt.Println("Insert rule to be used. Values accepted: 1 - 256. Press enter for default value (30)")
	localError := fmt.Errorf("Invalid input. Enter a number between 1 and 256")
	input, err := GetInput(reader)
	if err != nil {
		return 0, localError
	}
	if input == "" {
		return 30, nil
	}
	n, err := strconv.Atoi(input)
	if err != nil {
		return 0, localError
	}
	if n < 1 || n > 256 {
		return 0, localError
	}
	return uint8(n), nil
}

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

func GetRule(reader *bufio.Reader) ([]string, error) {
	n, err := askforRule(reader)
	if err != nil {
		return nil, err
	}
	return NewRule(n), nil
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	rule, err := GetRule(reader)
	if err != nil {
		fmt.Println(err)
		return
	}
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
