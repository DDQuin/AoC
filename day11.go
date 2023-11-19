package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Monkey struct {
	items          []int64
	operation      func(old int64) int64
	divisible      int64
	monkeyTrue     int
	monkeyFalse    int
	itemsInspected int
}

func createMonkeys(lines []string) []Monkey {
	monkeys := make([]Monkey, 0)

	startingItems := make([]int64, 0)
	var operation func(old int64) int64 = nil
	var divisible int64
	monkeyTrue := 2
	monkeyFalse := 3
	var err error
	for _, line := range lines {
		tokens := strings.Split(strings.TrimSpace(line), " ")
		if len(tokens) <= 1 {
			continue
		}

		if tokens[0] == "Starting" {
			for _, token := range tokens[2:] {
				var itemString string = token
				if itemString[len(itemString)-1:] == "," {
					itemString = itemString[0 : len(itemString)-1]
				}
				itemNum, err := strconv.ParseInt(itemString, 10, 64)
				if err != nil {
					log.Fatalf("convertting string: %s", err)
				}
				startingItems = append(startingItems, itemNum)
			}
		} else if tokens[0] == "Operation:" {
			operator := tokens[4]
			secondArg := tokens[5]
			if operator == "*" {
				if secondArg == "old" {
					operation = func(old int64) int64 {
						return old * old
					}
				} else {
					value, err := strconv.ParseInt(secondArg, 10, 64)
					if err != nil {
						log.Fatalf("convertting string: %s", err)
					}
					operation = func(old int64) int64 {
						return old * value
					}
				}
			} else if operator == "+" {
				if secondArg == "old" {
					operation = func(old int64) int64 {
						return old + old
					}
				} else {
					value, err := strconv.ParseInt(secondArg, 10, 64)
					if err != nil {
						log.Fatalf("convertting string: %s", err)
					}
					operation = func(old int64) int64 {
						return old + value
					}
				}
			}
		} else if tokens[0] == "Test:" {
			divisible, err = strconv.ParseInt(tokens[3], 10, 64)
			if err != nil {
				log.Fatalf("convertting string: %s", err)
			}
		} else if tokens[0] == "If" && tokens[1] == "true:" {
			monkeyTrue, err = strconv.Atoi(tokens[5])
			if err != nil {
				log.Fatalf("convertting string: %s", err)
			}

		} else if tokens[0] == "If" && tokens[1] == "false:" {
			monkeyFalse, err = strconv.Atoi(tokens[5])
			if err != nil {
				log.Fatalf("convertting string: %s", err)
			}
			monkeys = append(monkeys, Monkey{
				items:          startingItems,
				itemsInspected: 0,
				monkeyTrue:     monkeyTrue,
				monkeyFalse:    monkeyFalse,
				divisible:      divisible,
				operation:      operation,
			})
			startingItems = nil

		}

	}
	return monkeys
}

func Day11() {

	lines, err := ReadLines("resources/day11input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	monkeys := createMonkeys(lines)
	var supermod int64 = 1
	for i := range monkeys {
		var monkey *Monkey = &monkeys[i]
		supermod = supermod * monkey.divisible
	}
	endRound := 10000
	for round := 1; round <= endRound; round++ {
		for i := range monkeys {
			var monkey *Monkey = &monkeys[i]
			for len(monkey.items) != 0 {
				monkey.itemsInspected = monkey.itemsInspected + 1
				item := monkey.items[0]
				item = item % supermod // PARRT 2
				newLevel := monkey.operation(item)
				monkey.items = monkey.items[1:]
				//newLevel = newLevel / 3 PART 2 no dividing
				var monkeyToThrow *Monkey
				if newLevel%monkey.divisible == 0 {
					monkeyToThrow = &monkeys[monkey.monkeyTrue]
				} else {
					monkeyToThrow = &monkeys[monkey.monkeyFalse]
				}
				monkeyToThrow.items = append(monkeyToThrow.items, newLevel)
			}
		}
		fmt.Println("Round", round)
		for i, monkey := range monkeys {
			fmt.Println("Monkey ", i, "Items: ", monkey.items)
		}
	}

	for i, monkey := range monkeys {
		fmt.Println("Monkey ", i, "Inspected ", monkey.itemsInspected)
	}

}
