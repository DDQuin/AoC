package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day5() {
	lines, err := ReadLines("resources/day5input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var numStacks int = (len(lines[0]) + 1) / 4
	stacks := make([][]string, numStacks)
	for _, line := range lines {

		if line != "" && line[1:2] != "1" && line[0:1] != "m" {
			for i := 0; i < numStacks; i++ {
				var crateIndex int = i*4 + 1
				var possibleCrate string = line[crateIndex : crateIndex+1]
				if possibleCrate != " " {
					stacks[i] = append(stacks[i], possibleCrate)
				}
			}
		}

		if line != "" && line[0:1] == "m" {

			var tokens []string = strings.Split(line, " ")

			amount, err := strconv.Atoi(tokens[1])
			stackFrom, err := strconv.Atoi(tokens[3])
			stackTo, err := strconv.Atoi(tokens[5])

			if err != nil {
				log.Fatalf("cannot convert string to num: %s amount is %s stackfro mis %s and stackTo is %s", err, line[5:6], line[12:13], line[17:18])
			}

			var stackToChange []string = stacks[stackFrom-1][0:amount]
			var end int = len(stacks[stackFrom-1])
			stacks[stackFrom-1] = stacks[stackFrom-1][amount:end]
			var newStack []string
			// for i := len(stackToChange) - 1; i >= 0; i-- {
			// 	newStack = append(newStack, stackToChange[i])
			// } use this instead of below function for part1

			for _, crate := range stackToChange {
				newStack = append(newStack, crate)
			}

			for _, crate := range stacks[stackTo-1] {
				newStack = append(newStack, crate)
			}
			stacks[stackTo-1] = newStack
		}

	}
	fmt.Println(stacks)

}
