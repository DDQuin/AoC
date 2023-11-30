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
	var markLength int = 14

	for j := 0; j < len(signalLine)-markLength-1; j++ {

		for i := range signalLine[j : j+markLength] {

			var currentChar string = signalLine[j+i : j+i+1]
			if contains(signal, currentChar) {
				signal = nil
				break
			}
			signal = append(signal, currentChar)

			if len(signal) == markLength {
				lastMarker = j + i + 1
				break
			}
		}
		if lastMarker != -1 {
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
