package main

type AntDirection uint8

const (
	NORTH AntDirection = iota
	EAST
	SOUTH
	WEST
)

type Ant struct {
	position  Coord
	direction AntDirection
}

func (a *Ant) String() string {
	directionSymbol := []string{"↑", "→", "↓", "←"}
	return directionSymbol[a.direction]
}

func (a *Ant) TurnRight() {
	a.direction = (a.direction + 1) % 4
}

func (a *Ant) TurnLeft() {
	a.direction = (a.direction + 3) % 4
}

func (a *Ant) Walk(size uint8) {
	switch a.direction {
	case 0:
		a.position.row = (a.position.row + size - 1) % size
	case 1:
		a.position.col = (a.position.col + 1) % size
	case 2:
		a.position.row = (a.position.row + 1) % size
	case 3:
		a.position.col = (a.position.col + size - 1) % size
	}
}
