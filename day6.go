package main

import (
	"fmt"
	"log"
)

func Day6() {
	lines, err := ReadLines("resources/day6input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var signal []string
	var lastMarker int = -1
	var signalLine string = lines[0]

	for j := range signalLine {
		var currentChar string = signalLine[j : j+1]
		if contains(signal, currentChar) {
			signal = nil
		}
		signal = append(signal, currentChar)
		if len(signal) == 4 {
			lastMarker = j + 1
			break
		}
	}

	fmt.Println("Last marker is ", lastMarker)
}

func contains(arr []string, needle string) bool {
	for _, str := range arr {
		if str == needle {
			return true
		}
	}
	return false
}
