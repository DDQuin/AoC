package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
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

func getNetworkAndInstructions(lines []string) (map[string][]string, []int, []string) {
	instructions := make([]int, 0)
	network := make(map[string][]string, 0)
	nodesEndA := make([]string, 0)
	insString := lines[0]
	for _, instruction := range insString {
		if instruction == 'L' {
			instructions = append(instructions, 0)
		} else if instruction == 'R' {
			instructions = append(instructions, 1)
		} else {
			log.Fatalf("Not a valid instruction: %s", string(instruction))
		}
	}

	for i := 2; i < len(lines); i++ {
		curString := strings.Split(lines[i], "=")
		node := strings.Trim(curString[0], " ")
		if node[len(node)-1:] == "A" {
			nodesEndA = append(nodesEndA, node)
		}

		rightSide := strings.Trim(strings.Split(strings.Trim(curString[1], " "), ",")[1], " ")[:3]
		leftSide := strings.Split(strings.Trim(curString[1], " "), ",")[0][1:]

		network[node] = []string{leftSide, rightSide}
	}

	return network, instructions, nodesEndA
}

func getSteps(network map[string][]string, instructions []int, startNode string) int {
	steps := 0
	iCounter := 0
	curNode := startNode
	for !allNodesZ([]string{curNode}) {
		curNode = network[curNode][instructions[iCounter]]
		steps++
		iCounter = (iCounter + 1) % len(instructions)
	}
	return steps
}

func allNodesZ(nodes []string) bool {
	for _, node := range nodes {
		if node[len(node)-1:] != "Z" {
			return false
		}
	}
	return true
}

func getStepsAll(network map[string][]string, instructions []int, startNodes []string) int {
	steps := 0
	iCounter := 0
	curNodes := startNodes
	for !allNodesZ(curNodes) {
		for i := range curNodes {
			curNodes[i] = network[curNodes[i]][instructions[iCounter]]
		}
		steps++
		iCounter = (iCounter + 1) % len(instructions)
	}
	return steps
}

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

// find Least Common Multiple (LCM) via GCD
func LCMD(nums []int) int {
	result := nums[0] * nums[1] / GCD(nums[0], nums[1])

	for i := 2; i < len(nums); i++ {
		result = LCM(result, nums[i])
	}

	return result
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	network, instructions, nodesA := getNetworkAndInstructions(lines)
	allSteps := make([]int, 0)

	for _, node := range nodesA {
		allSteps = append(allSteps, getSteps(network, instructions, node))
	}
	fmt.Println(LCMD(allSteps))

}
