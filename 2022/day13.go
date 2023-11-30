package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type LinePair struct {
	line1 string
	line2 string
}

type ListIntegerNode struct {
	parent *ListIntegerNode
	value int
	list []*ListIntegerNode

}

func (node ListIntegerNode) printNode(isEnd bool) {
	if node.value != -1 {
		if isEnd {
			fmt.Print(node.value)
		} else {
			fmt.Print(node.value, " ")
		}
		return
	} else {
		fmt.Print("[")
		for i, nodeChild := range node.list {
			if i == len(node.list) - 1{
				nodeChild.printNode(true)
			} else {
				nodeChild.printNode(false)
			}
		}
		fmt.Print("]")
	}

}

func createLinePairs(lines []string) []LinePair {
	linePairs := make([]LinePair, 0)
	linePair := LinePair{line1: "", line2: ""}
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		if linePair.line1 == "" {
			linePair.line1 = line
		} else {
			linePair.line2 = line
			linePairs = append(linePairs, linePair)
			linePair.line1 = ""
			linePair.line2 = ""
		}
		
	}
	return linePairs
}

func createPackets(lines []string) []string {
	packets := make([]string, 0)
	packets = append(packets, "[[2]]")
	packets = append(packets, "[[6]]")
	for _, line := range lines {
		if len(line) == 0 {
			continue
		}
		packets = append(packets, line)
		
	}	
	return packets
}

func convertLineToNode(line string) *ListIntegerNode {
	newLine := line
	var appended int = 0
	for i, char := range line {
		if char == '[' {
			newLine = newLine[:i + 1 + appended] + "," + newLine[i + 1 + appended:]
			appended++;
		}
		if char == ']' {
			newLine = newLine[:i  + appended] + "," + newLine[i  + appended:]
			appended++;
		}
	}


	var curNode *ListIntegerNode

	tokens := strings.Split(newLine, ",")

	for _, token := range tokens {
		if len(token) == 0 {

		} else if token[0:1] != "[" && token[0:1] != "]"  {
			value, err := strconv.Atoi(token)
			if err != nil {
				log.Fatalf("converting string to num: %s", err)
			}
			newValNode := ListIntegerNode{value: value, parent: curNode }
			curNode.list = append(curNode.list, &newValNode)
		} else if token[0:1] == "[" {
			newListNode := ListIntegerNode{value: -1, parent: curNode, list: make([]*ListIntegerNode, 0)}
			if curNode != nil {
				curNode.list = append(curNode.list, &newListNode)
			}
			curNode = &newListNode

		} else if token[0:1] == "]" {
			if curNode.parent != nil {
				curNode = curNode.parent
			}
		}

	}

	return curNode
	
}

func (linePair LinePair) isLinePairRightOrder() bool {
	firstNode := convertLineToNode(linePair.line1)
	secondNode := convertLineToNode(linePair.line2)
	result := compareLineNode(*firstNode, *secondNode)
	if result > 0 {
		return true
	}
	if result < 0 {
		return false
	}

	log.Fatal("Sometghin went wrong, samme list")
	return false
}

func compareLineNode(node1 ListIntegerNode, node2 ListIntegerNode) int {

	if node1.value == -1 && node2.value == -1 {
		
		for i := 0; i < len(node1.list) && i < len(node2.list); i++ {
			result := compareLineNode(*node1.list[i], *node2.list[i])
			if result == 1 {
				return 1
			} else if result == -1 {
				return -1
			} else {

			}
		}
		if len(node1.list) < len(node2.list) {
			return 1
		} else if len(node2.list) < len(node1.list) {
		
			return -1
		}

	}
	if node1.value != -1 && node2.value != -1 {
		if node1.value < node2.value {
			return 1
		} else if node1.value > node2.value {
			return -1
		}
		return 0
	}

	if node1.value == -1 && node2.value != -1 {
	
		node2.list = make([]*ListIntegerNode, 0)
		newValNode := ListIntegerNode{value: node2.value, parent: &node2 }
		node2.list = append(node2.list, &newValNode)
		node2.value = -1
		return compareLineNode(node1, node2)
	}
	if node1.value != -1 && node2.value == -1 {
	
		node1.list = make([]*ListIntegerNode, 0)
		newValNode := ListIntegerNode{value: node1.value, parent: &node1 }
		node1.list = append(node1.list, &newValNode)
		node1.value = -1
		return compareLineNode(node1, node2)
	}
	return 0
	
}

func bubbleSort(packetsToSort []string) {
	length := len(packetsToSort)
	for i := 0; i < length; i++ {
		for j := 0; j < length - 1 - i; j++ {
			listPair := LinePair{line1: packetsToSort[j], line2: packetsToSort[j + 1]}
			if !listPair.isLinePairRightOrder() {
				temp := packetsToSort[j]
				packetsToSort[j] = packetsToSort[j + 1]
				packetsToSort[j + 1] = temp
			}
			
		}
	}

}


func Day13() {

	
	lines, err := ReadLines("resources/day13input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	linePairs := createLinePairs(lines)
	allPackets := createPackets(lines)
	
	bubbleSort(allPackets)
	
	decoder := 1
	for i, packet := range allPackets {
		if packet == "[[2]]" {
			decoder = decoder * (i + 1)
		}
		if packet == "[[6]]" {
			decoder = decoder * (i + 1)
		}
	}
	fmt.Println("decoder is ", decoder)

	sum := 0
	for i, linePair := range linePairs {
		if linePair.isLinePairRightOrder() {
			//fmt.Println("Is right ", i + 1)
			sum = sum + (i + 1)
		}
	}
	
	fmt.Println("Sum of pairs in right order is ", sum)
}
