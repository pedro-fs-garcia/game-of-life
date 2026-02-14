package main

import (
	"math"
	"strconv"
	"strings"
)

func toDecimal(bin string) int {
	decimal := 0
	digits := strings.Split(bin, "")
	l := len(bin)
	for i := range bin {
		index := (i + l - 1) % l
		d, _ := strconv.Atoi(digits[index])
		if d == 0 {
			continue
		}
		decimal += int(math.Pow(2, float64(i)))
	}
	return decimal
}

func getDigitFromRule(decimal int, rule []string) int {
	i := len(rule) - 1 - decimal
	digit, _ := strconv.Atoi(rule[i])
	return digit
}

func toBinary(decimal uint8) string {
	var sb strings.Builder
	for decimal > 0 {
		rest := decimal % 2
		sb.WriteString(strconv.Itoa(int(rest)))
		decimal = decimal / 2
	}
	temp := strings.Split(sb.String(), "")
	var bin strings.Builder
	for i := len(sb.String()) - 1; i >= 0; i-- {
		bin.WriteString(temp[i])
	}
	return bin.String()
}
