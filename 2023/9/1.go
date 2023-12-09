package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func getSequence(line string) []int {
	seq := make([]int, 0)
	seqSplit := strings.Split(line, " ")
	for _, numString := range seqSplit {
		num, _ := strconv.Atoi(numString)
		seq = append(seq, num)
	}
	return seq
}

func isSeqAllZeroes(seq []int) bool {
	for _, num := range seq {
		if num != 0 {
			return false
		}
	}
	return true
}

func createDiffSeq(seq []int) []int {
	diff := make([]int, 0)
	for i := 0; i < len(seq)-1; i++ {
		numOne := seq[i]
		numTwo := seq[i+1]
		diff = append(diff, numTwo-numOne)
	}
	return diff
}

func getAllSequences(seq []int) [][]int {
	allSeq := make([][]int, 0)
	allSeq = append(allSeq, seq)
	curSeq := seq
	for !isSeqAllZeroes(curSeq) {
		curSeq = createDiffSeq(curSeq)
		allSeq = append(allSeq, curSeq)
	}
	return allSeq
}

func getNextValue(allSeq [][]int) int {
	lastIndex := len(allSeq) - 1
	allSeq[lastIndex] = append(allSeq[lastIndex], 0)
	for i := lastIndex - 1; i >= 0; i-- {
		curSeq := allSeq[i]
		prevSeq := allSeq[i+1]
		newVal := prevSeq[len(prevSeq)-1] + curSeq[len(curSeq)-1]
		allSeq[i] = append(allSeq[i], newVal)
	}
	return allSeq[0][len(allSeq[0])-1]
}

func getNextValueBack(allSeq [][]int) int {
	lastIndex := len(allSeq) - 1
	allSeq[lastIndex] = append(allSeq[lastIndex], 0)
	for i := lastIndex - 1; i >= 0; i-- {
		curSeq := allSeq[i]
		prevSeq := allSeq[i+1]
		newVal := curSeq[0] - prevSeq[len(prevSeq)-1]
		allSeq[i] = append(allSeq[i], newVal)
	}
	return allSeq[0][len(allSeq[0])-1]
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	total := 0
	totalBack := 0

	for _, line := range lines {
		seq := getSequence(line)
		allSeq := getAllSequences(seq)
		nextVal := getNextValue(allSeq)
		nextValBack := getNextValueBack(allSeq)
		total += nextVal
		totalBack += nextValBack
	}

	fmt.Println("Total is", total, " Back total is", totalBack)

}
