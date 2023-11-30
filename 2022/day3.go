package main

import (
	"fmt"
	"log"
)

func GetSameItem(rucksack string) string {
	var comp1 string = rucksack[0 : len(rucksack)/2]
	var comp2 string = rucksack[len(rucksack)/2:]
	//fmt.Println("comp1 ", comp1, "comp2", comp2)
	for i, item1 := range comp1 {
		for _, item2 := range comp2 {
			if item1 == item2 {
				//fmt.Println("Item 1 ", string(item1))
				return rucksack[i : i+1]
			}
		}
	}
	return ""
}

func GetBadge(rucksack1 string, rucksack2 string, rucksack3 string) string {
	for i, item1 := range rucksack1 {
		for _, item2 := range rucksack2 {
			for _, item3 := range rucksack3 {
				if item1 == item2 && item2 == item3 {
					return rucksack1[i : i+1]
				}
			}
		}
	}
	return ""
}

func GetItemPrio(item string) int {
	var char int = int(item[0])
	if char >= 'A' && char <= 'Z' {
		return char - 'A' + 27
	} else {
		return char - 'a' + 1
	}
}

func Day3() {
	lines, err := ReadLines("resources/day3input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var totalPrio int = 0
	var totalPrioElf int = 0
	for _, rucksack := range lines {
		var sameItem string = GetSameItem(rucksack)
		totalPrio += GetItemPrio(sameItem)
	}

	for i := 0; i < len(lines)/3; i++ {
		var rucksack1 string = lines[i*3]
		var rucksack2 string = lines[i*3+1]
		var rucksack3 string = lines[i*3+2]
		var badge string = GetBadge(rucksack1, rucksack2, rucksack3)
		totalPrioElf += GetItemPrio(badge)
	}

	fmt.Println("Sum of priorities is", totalPrio)
	fmt.Println("Sum of priorities 3 elf is", totalPrioElf)

}
