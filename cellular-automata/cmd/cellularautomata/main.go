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

func askForRule(reader *bufio.Reader) (uint8, error) {
	fmt.Println("Insert rule to be used. Values accepted: 0 - 255. Press enter for default value (30)")
	localError := fmt.Errorf("Invalid input. Enter a number between 0 and 255")
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
	if n < 0 || n > 255 {
		return 0, localError
	}
	return uint8(n), nil
}

func GetRule(reader *bufio.Reader) (uint8, error) {
	rule, err := askForRule(reader)
	if err != nil {
		return 0, err
	}
	return rule, nil
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
