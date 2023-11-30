package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"
)

type MonkeyNode struct {
	name     string
	left     string
	right    string
	operator string
	value    int
}

func createMonkeyMap(lines []string) map[string]MonkeyNode {
	nodesMap := make(map[string]MonkeyNode, 0)
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		name := tokens[0][0 : len(tokens[0])-1]
		var left string = ""
		var right string = ""
		var operator string = ""
		value := -1
		if len(tokens) == 2 { // Value nodde
			num, _ := strconv.Atoi(tokens[1])
			value = num
		} else { // extra node
			left = tokens[1]
			operator = tokens[2]
			right = tokens[3]
		}
		node := MonkeyNode{name: name, left: left, right: right, operator: operator, value: value}
		nodesMap[name] = node
	}
	return nodesMap
}

func evalMonkey(monkeyMap map[string]MonkeyNode, monkey string) int {
	node := monkeyMap[monkey]
	if node.operator == "" {
		return node.value
	}
	if node.operator == "+" {
		return evalMonkey(monkeyMap, node.left) + evalMonkey(monkeyMap, node.right)
	} else if node.operator == "-" {
		return evalMonkey(monkeyMap, node.left) - evalMonkey(monkeyMap, node.right)
	} else if node.operator == "*" {
		return evalMonkey(monkeyMap, node.left) * evalMonkey(monkeyMap, node.right)
	} else if node.operator == "/" {
		return evalMonkey(monkeyMap, node.left) / evalMonkey(monkeyMap, node.right)
	}

	log.Fatal("Something wrong happened, operator", node.operator, "isnt working")
	return -1
}

func Day21() {
	lines, err := ReadLines("resources/day21input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	monkeyMap := createMonkeyMap(lines)
	endVal := evalMonkey(monkeyMap, "root")

	leftNode := monkeyMap[monkeyMap["root"].left]
	rightNode := monkeyMap[monkeyMap["root"].right]

	humanVal := monkeyMap["humn"]
	monkeyMap["humn"] = MonkeyNode{name: humanVal.name, left: humanVal.left, right: humanVal.right, value: 0, operator: humanVal.operator}

	leftVal := evalMonkey(monkeyMap, leftNode.name)
	rightVal := evalMonkey(monkeyMap, rightNode.name)

	nodeToChange := rightNode
	nodeCorrect := leftNode
	if leftVal < rightVal {
		nodeToChange = leftNode
		nodeCorrect = rightNode
	}

	//Binary search insipiration taken from https://github.com/mnml/aoc/blob/main/2022/21/1.go
	part2, _ := sort.Find(1e16, func(v int) int {
		monkeyMap["humn"] = MonkeyNode{name: humanVal.name, left: humanVal.left, right: humanVal.right, value: v, operator: humanVal.operator}
		return evalMonkey(monkeyMap, nodeCorrect.name) - evalMonkey(monkeyMap, nodeToChange.name)
	})

	fmt.Println("Root yells", endVal)

	fmt.Println("Human val is", part2)
}
