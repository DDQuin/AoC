package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
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

const (
	_ int = iota
	HIGHCARD
	ONEPAIR
	TWOPAIR
	THREEKIND
	FULLHOUSE
	FOURKIND
	FIVEKIND
)

func getCardStrength(card byte) int {
	if card >= '2' && card <= '9' {
		return int(card - '2')
	}
	if card == 'T' {
		return 10
	}
	if card == 'J' {
		return 11
	}
	if card == 'Q' {
		return 12
	}
	if card == 'K' {
		return 13
	}
	if card == 'A' {
		return 14
	}
	return -1
}

type Hand struct {
	cards    string
	bid      int
	handType int
}

func createHands(lines []string) []Hand {
	hands := make([]Hand, 0)
	for _, line := range lines {
		cardString := strings.Split(line, " ")[0]
		bidString := strings.Split(line, " ")[1]

		bid, _ := strconv.Atoi(bidString)
		hands = append(hands, Hand{cards: cardString, bid: bid, handType: -1})
	}
	return hands
}

func getHandType(hand Hand) int {
	countToLabel := make(map[int][]byte, 0)
	labelCount := make(map[byte]int, 0)
	for i := range hand.cards {
		labelCount[hand.cards[i]]++
	}

	for label, val := range labelCount {
		countToLabel[val] = append(countToLabel[val], label)
	}

	if len(countToLabel[5]) != 0 {
		return FIVEKIND
	} else if len(countToLabel[4]) != 0 {
		return FOURKIND
	} else if len(countToLabel[3]) != 0 {
		if len(countToLabel[2]) != 0 {
			return FULLHOUSE
		} else {
			return THREEKIND
		}
	} else if len(countToLabel[2]) != 0 {
		if len(countToLabel[2]) == 2 {
			return TWOPAIR
		} else {
			return ONEPAIR
		}
	}
	return HIGHCARD
}

func sortHand(hands []Hand) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[j].handType == hands[i].handType {
			for idx := range hands[j].cards {
				curCardJ := getCardStrength(hands[j].cards[idx])
				curCardI := getCardStrength(hands[i].cards[idx])

				if curCardJ > curCardI {
					return true
				} else if curCardJ < curCardI {
					return false
				}
			}
		}
		return hands[j].handType > hands[i].handType
	})
	return hands
}

func main() {
	lines, err := ReadLines("test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	hands := createHands(lines)
	for i := range hands {
		hands[i].handType = getHandType(hands[i])
	}
	sortedHands := sortHand(hands)
	winnings := 0
	for i, hand := range sortedHands {
		winnings += (hand.bid * (i + 1))
	}

	fmt.Println(winnings)

}
