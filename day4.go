package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func Day4() {
	lines, err := ReadLines("resources/day4input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var total int = 0
	var totalOverlaps int = 0
	for _, line := range lines {
		var pairs []string = strings.Split(line, ",")
		var pairOne []string = strings.Split(pairs[0], "-")
		var pairTwo []string = strings.Split(pairs[1], "-")

		pairOneA, err := strconv.Atoi(pairOne[0])
		pairOneB, err := strconv.Atoi(pairOne[1])

		pairTwoA, err := strconv.Atoi(pairTwo[0])
		pairTwoB, err := strconv.Atoi(pairTwo[1])

		if err != nil {
			log.Fatalf("cannot convert string to num: %s", err)
		}

		if pairOneA <= pairTwoA && pairOneB >= pairTwoB || pairTwoA <= pairOneA && pairTwoB >= pairOneB {
			total++
		}

		if !(((pairOneA < pairTwoA) && (pairOneB < pairTwoA)) || ((pairOneA > pairTwoB) && (pairOneB > pairTwoB))) {
			totalOverlaps++
		}

	}
	fmt.Println("Number of assignment pairs fully containing the other is", total)
	fmt.Println("Number of assignment pairs overlapping is", totalOverlaps)

}
