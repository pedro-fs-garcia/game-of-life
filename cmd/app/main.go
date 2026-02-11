package main

import (
	"GameOfLife/cmd/app/game"
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

func getInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error when reading input: ", err)
	}
	input = strings.TrimSpace(input)
	return input
}

func parseCoord(s string) ([]game.Coord, error) {
	var coords []game.Coord
	var x, y int
	r := strings.NewReader((s))
	if _, err := fmt.Fscanf(r, "{"); err != nil {
		return nil, err
	}

	for {
		_, err := fmt.Fscanf(r, "(%d,%d)", &x, &y)
		if err == nil {
			coords = append(coords, game.Coord{X: x, Y: y})
		} else if err == io.EOF {
			break
		} else {
			return nil, err
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

func getInitialCoords() map[game.Coord]struct{} {
	fmt.Println("Starting the Game Of Life")
	fmt.Println("Choose the initial coordinates. Write them in the format '{(x,y),(x,y),(x,y),...}'")
	fmt.Println("Press Enter to start with a standard board.")
	reader := bufio.NewReader(os.Stdin)
	input := getInput(reader)
	coords, err := parseCoord(input)
	if err != nil {
		fmt.Println(err)
		return game.StandardInitialCoordinates(20)
	}
	coordMap := make(map[game.Coord]struct{})
	for i := range coords {
		coordMap[coords[i]] = struct{}{}
	}
	return coordMap
}

func main() {
	c := getInitialCoords()
	var b *game.Board = game.NewBoard(20, c)
	fmt.Println(b)
	time.Sleep(1 * time.Second)
	for {
		b, changes := b.Round()
		fmt.Println(b)
		if b.Stop(changes) {
			break
		}
		time.Sleep(1 * time.Second)
	}
}
