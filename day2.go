package main

import (
	"fmt"
	"log"
)

var rpsMap map[string]int = make(map[string]int)
var rpsWin map[string]string = make(map[string]string)

var rpsWinOpp map[string]string = make(map[string]string)
var rpsLossOpp map[string]string = make(map[string]string)
var rpsDrawOpp map[string]string = make(map[string]string)

func RpsScore(opp string, you string) int {
	var score int = 0
	score += rpsMap[you]
	var won bool = rpsWin[you] == opp
	if won {
		score += 6
	} else if (you == "X" && opp == "A") || (you == "Y" && opp == "B") || (you == "Z" && opp == "C") {
		score += 3
	} else {
		// lost
	}
	return score
}

func GetChoice(opp string, result string) string {
	if result == "Y" {
		return rpsDrawOpp[opp]
	}
	if result == "X" {
		return rpsWinOpp[opp]
	}
	return rpsLossOpp[opp]
}

func Day2() {
	rpsMap["X"] = 1
	rpsMap["Y"] = 2
	rpsMap["Z"] = 3

	rpsWin["X"] = "C"
	rpsWin["Y"] = "A"
	rpsWin["Z"] = "B"

	rpsWinOpp["A"] = "Z"
	rpsWinOpp["B"] = "X"
	rpsWinOpp["C"] = "Y"

	rpsLossOpp["A"] = "Y"
	rpsLossOpp["B"] = "Z"
	rpsLossOpp["C"] = "X"

	rpsDrawOpp["A"] = "X"
	rpsDrawOpp["B"] = "Y"
	rpsDrawOpp["C"] = "Z"

	lines, err := ReadLines("resources/day2input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var total int = 0

	for _, line := range lines {
		var opp string = line[0:1]
		var you string = line[2:3]
		var youSecond string = GetChoice(opp, you)
		total += RpsScore(opp, youSecond)
	}
	fmt.Println("total score is ", total)

}
