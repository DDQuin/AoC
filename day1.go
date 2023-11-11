package main

import (
	"fmt"
	"log"
	"strconv"
)

func day1() {
	lines, err := readLines("resources/day1input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var total int = 0
	var max int = 0
	// print file contents
	for _, line := range lines {
		if line == "" {
			if total > max {
				max = total
			}
			total = 0
		} else {
			calorie, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			total += calorie
		}
	}
	fmt.Println("Max cal is ", max)

}
