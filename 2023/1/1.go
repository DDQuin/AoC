package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func ReadLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func getNumFromString(line string) string {
	digitMaps := map[string]string{"one": "1", "two": "2", "three": "3", "four": "4",
		"five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	for key, value := range digitMaps {
		if strings.Contains(line, key) {
			return value
		}

	}
	return ""

}

func getcalValue(line string) int {
	firstDigit := ""
	secondDigit := ""

	for i := 0; i < len(line); i++ {
		secondIndex := len(line) - i - 1
		if firstDigit == "" {
			if line[i] >= '0' && line[i] <= '9' {
				firstDigit = string(line[i])
			} else {
				firstDigit = getNumFromString(line[0 : i+1])
			}
		}
		if secondDigit == "" {
			if line[secondIndex] >= '0' && line[secondIndex] <= '9' {
				secondDigit = string(line[secondIndex])
			} else {
				secondDigit = getNumFromString(line[secondIndex:])
			}
		}
	}
	calVal, _ := strconv.Atoi(firstDigit + secondDigit)
	return calVal
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	total := 0
	for _, line := range lines {
		num := getcalValue(line)
		total += num
	}

	fmt.Println("Total is", total)

}
