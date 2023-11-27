package main

import (
	"fmt"
	"log"
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
	fmt.Println("Root yells", endVal)
}
