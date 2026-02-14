package main

import (
	"strconv"
	"strings"
)

func toDecimal(bin string) int {
	decimal := 0
	for i := 0; i < len(bin); i++ {
		decimal = decimal*2 + int(bin[i]-'0')
	}
	return decimal
}

func getDigitFromRule(decimal int, rule []string) int {
	i := len(rule) - 1 - decimal
	digit, _ := strconv.Atoi(rule[i]) // rule is astring made of "0"s and "1"s only
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
