package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
)

func Day1() {
	lines, err := ReadLines("resources/day1input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var totalCalElf []int
	var total int = 0

	for _, line := range lines {
		if line == "" {
			totalCalElf = append(totalCalElf, total)
			total = 0
		} else {
			calorie, err := strconv.Atoi(line)
			if err != nil {
				log.Fatalf("converting to %s int: %s", line, err)
			}
			total += calorie
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(totalCalElf)))
	fmt.Println("top 3 sum is ", totalCalElf[0]+totalCalElf[1]+totalCalElf[2])

}
