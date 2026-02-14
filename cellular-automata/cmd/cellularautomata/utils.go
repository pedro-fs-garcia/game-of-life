package main

func getDigitFromRule(index uint8, rule uint8) uint8 {
	rule = rule >> index
	return rule & 1
}
