package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
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

func createCards(lines []string) map[int][]Card {
	cards := make(map[int][]Card, 0)
	for i, line := range lines {
		winNums := make([]int, 0)
		playerNums := make([]int, 0)

		cardString := strings.Split(line, ":")[1]

		winNumString := strings.Split(strings.Trim(strings.Split(cardString, "|")[0], " "), " ")
		playerNumString := strings.Split(strings.Trim(strings.Split(cardString, "|")[1], " "), " ")
		for _, winString := range winNumString {
			if winString == "" {
				continue
			}
			num, _ := strconv.Atoi(winString)
			winNums = append(winNums, num)
		}
		for _, playerString := range playerNumString {
			if playerString == "" {
				continue
			}
			num, _ := strconv.Atoi(playerString)
			playerNums = append(playerNums, num)
		}

		cards[i+1] = []Card{{-1, winNums, playerNums}}
	}
	return cards
}

type Card struct {
	matching   int
	winNums    []int
	playerNums []int
}

func containsNum(target int, arr []int) bool {
	for _, num := range arr {
		if num == target {
			return true
		}
	}

	return false
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	cards := createCards(lines)

	sum := 0
	for _, cardList := range cards {
		winningNums := 0
		for _, winNum := range cardList[0].winNums {
			if containsNum(winNum, cardList[0].playerNums) {
				winningNums++
			}
		}
		cardList[0].matching = winningNums
		if winningNums == 0 {
			continue
		}
		points := int(math.Pow(2, float64(winningNums)-1))
		sum += points
	}

	for cardNum := 1; cardNum < len(cards)+1; cardNum++ {
		cardList := cards[cardNum]
		matching := cardList[0].matching
		for x := 0; x < len(cardList); x++ {
			for i := 0; i < matching; i++ {
				newCardNum := cardNum + i + 1
				cards[newCardNum] = append(cards[newCardNum], cards[newCardNum][0])
			}
		}
	}

	fmt.Println("scratchcard is worth", sum)
	totalScratch := 0
	for _, card := range cards {
		totalScratch += len(card)
	}
	fmt.Println("total scratch is", totalScratch)

}
