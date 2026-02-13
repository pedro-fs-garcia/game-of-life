package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var parsingError = errors.New("Input string is not valid. Follow the model.")

func getInput(reader *bufio.Reader) (string, error) {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error when reading input.", err)
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func parseCoords(input string) ([]Coord, error) {
	var coords []Coord
	var row, col uint8
	r := strings.NewReader((input))
	_, err := fmt.Fscanf(r, "{")
	if err != nil {
		return nil, parsingError
	}

	for {
		_, err := fmt.Fscanf(r, "(%d,%d)", &row, &col)
		if err == nil {
			coords = append(coords, Coord{row, col})
		} else if err == io.EOF {
			break
		} else {
			return nil, parsingError
		}

		var c rune
		if _, err := fmt.Fscanf(r, "%c", &c); err != nil {
			return nil, err
		}
		if c == '}' {
			break
		}
		if c != ',' {
			return nil, fmt.Errorf("expected ',' or '}', got %q", c)
		}
	}
	return coords, nil
}

func getInitialCoords() (map[Coord]struct{}, error) {
	fmt.Println("Choose the initial coordinates. Write them in the format '{(x,y),(x,y),(x,y),...}'")
	fmt.Println("Press Enter to start with a standard board.")
	reader := bufio.NewReader(os.Stdin)
	input, err := getInput(reader)
	if err != nil {
		return nil, parsingError
	}
	if input == "" {
		return nil, nil
	}
	coords, err := parseCoords(input)
	if err != nil {
		return nil, err
	}
	coordMap := make(map[Coord]struct{})
	for i := range coords {
		coordMap[coords[i]] = struct{}{}
	}
	return coordMap, nil
}

func main() {
	fmt.Println("Starting Langton's Ant")
	coordsMap, err := getInitialCoords()
	if err != nil {
		fmt.Println(err)
		return
	}
	b := NewBoard(20, coordsMap)
	fmt.Println(b)
	for {
		time.Sleep(50 * time.Millisecond)
		b.RunRound()
		fmt.Println(b)
	}
}
