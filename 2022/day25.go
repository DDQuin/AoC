package main

import (
	"fmt"
	"log"
	"math"
)

func Day25() {
	lines, err := ReadLines("resources/day25input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	digitMap := map[byte]int{'2': 2, '1': 1, '0': 0, '-': -1, '=': -2}
	total := 0
	for _, line := range lines {
		total += convertSNAFUToDecimal(line, 5, digitMap)

	}

	fmt.Println("Sum of SNAFU as decimal numbers is", total)

	fmt.Println("Snafu is", convertDecimalToSnafu(total))

}

func convertDecimalToSnafu(num int) string {
	// This function code taken from //from https://github.com/mnml/aoc/blob/main/2022/25/1.go
	snafu := ""
	for num > 0 {
		snafu = string("=-012"[(num+2)%5]) + snafu
		num = (num + 2) / 5
	}
	return snafu
}

func convertSNAFUToDecimal(snafu string, base int, digitMap map[byte]int) int {
	decimal := 0
	for i := 0; i < len(snafu); i++ {
		curDigit := snafu[i]
		position := len(snafu) - 1 - i
		start := int(math.Pow(float64(base), float64(position)))
		curDec := start * digitMap[curDigit]
		decimal += curDec
	}
	return decimal
}
