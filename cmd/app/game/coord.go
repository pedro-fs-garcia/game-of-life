package game

type Coord struct {
	X int
	Y int
}

func StandardInitialCoordinates(size int) map[Coord]struct{} {
	alive := make(map[Coord]struct{})
	for i := range size {
		alive[Coord{X: i, Y: i}] = struct{}{}
		alive[Coord{X: i, Y: size - 1 - i}] = struct{}{}
	}
	return alive
}

func (c Coord) belongsTo(coord map[Coord]struct{}) bool {
	_, ok := coord[c]
	return ok
}
