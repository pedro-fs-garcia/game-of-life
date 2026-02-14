package main

type Cell struct {
	currentState bool
	nextState    bool
}

func (c *Cell) String() string {
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

func (c *Cell) GetPattern(left, right Cell) string {
	pattern := left.String() + c.String() + right.String()
	return pattern
}

func (c *Cell) CalculateNextState(pattern string, rule []string) {
	decimal := toDecimal(pattern)
	d := getDigitFromRule(decimal, rule)
	if d == 1 {
		c.nextState = true
	} else {
		c.nextState = false
	}
}
