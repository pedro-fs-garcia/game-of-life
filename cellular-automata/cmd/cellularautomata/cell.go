package main

type Cell struct {
	currentState bool
	nextState    bool
}

func (c *Cell) String() string {
	if c.currentState {
		return "1"
	}
	return "0"
}

func (c *Cell) SetNextState(nextState bool) {
	c.nextState = nextState
}

func (c *Cell) RollRound() {
	c.currentState = c.nextState
}
