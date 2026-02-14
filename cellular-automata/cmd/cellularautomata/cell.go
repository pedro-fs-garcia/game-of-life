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

func (c *Cell) Symbol() string {
	if c.currentState {
		return "1"
	}
	return " "
}

func (c *Cell) SetNextState(nextState bool) {
	c.nextState = nextState
}

func (c *Cell) RollRound() {
	c.currentState = c.nextState
}

func (c *Cell) GetPattern(left, right *Cell) string {
	pattern := left.String() + c.String() + right.String()
	return pattern
}

func (c *Cell) GetPatternIndex(left, right *Cell) uint8 {
	idx := uint8(0)
	if left.currentState {
		idx += 4
	}
	if c.currentState {
		idx += 2
	}
	if right.currentState {
		idx += 1
	}
	return idx
}

func (c *Cell) CalculateNextState(left, right *Cell, rule []string) {
	decimal := c.GetPatternIndex(left, right)
	d := getDigitFromRule(decimal, rule)
	if d == 1 {
		c.nextState = true
	} else {
		c.nextState = false
	}
}
