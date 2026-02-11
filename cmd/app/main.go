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

func GetInput(reader *bufio.Reader) string {
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error when reading input: ", err)
	}
	input = strings.TrimSpace(input)
	return input
}

func ParseCoord(s string) ([]game.Coord, error) {
	var coords []game.Coord
	var x, y int
	r := strings.NewReader((s))
	for {
		_, err := fmt.Fscanf(r, " { (%d, %d)", &x, &y)
		if err == nil {
			coords = append(coords, game.Coord{X: x, Y: y})
		} else if err == io.EOF {
			break
		} else {
			return nil, err
		}
		fmt.Fscanf(r, ",")
		fmt.Fscanf(r, "}")
	}
	return coords, nil
}

func getInitialCoords() map[game.Coord]struct{} {
	fmt.Println("Starting the Game Of Life")
	fmt.Println("Choose the initial coordinates. Write them in the format '{(x,y),(x,y),(x,y),...}'")
	fmt.Println("Press Enter to start with a standard board.")
	reader := bufio.NewReader(os.Stdin)
	input := GetInput(reader)
	coords, err := ParseCoord(input)
	if err != nil {
		return game.StandardInitialCoordinates(20)
	}
	var coordMap map[game.Coord]struct{}
	for i := range coords {
		coordMap[coords[i]] = struct{}{}
	}
	return coordMap
}

func main() {
	c := game.StandardInitialCoordinates(20)
	var b *game.Board = game.NewBoard(20, c)
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
