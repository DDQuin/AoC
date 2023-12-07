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

func getCardStrength(card byte, isJoker bool) int {
	if card >= '2' && card <= '9' {
		return int(card - '2')
	}
	if card == 'T' {
		return 10
	}
	if card == 'J' {
		if isJoker {
			return -1
		} else {
			return 11
		}
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

func getHandType(hand Hand, isJoker bool) int {
	countToLabel := make(map[int][]byte, 0)
	labelCount := make(map[byte]int, 0)
	for i := range hand.cards {
		labelCount[hand.cards[i]]++
	}

	for label, val := range labelCount {
		countToLabel[val] = append(countToLabel[val], label)
	}

	if isJoker {
		if labelCount['J'] == 4 {
			return FIVEKIND
		} else if labelCount['J'] == 3 {
			if len(countToLabel[2]) != 0 {
				return FIVEKIND
			} else {
				return FOURKIND
			}
		} else if labelCount['J'] == 2 {
			if len(countToLabel[3]) != 0 {
				return FIVEKIND
			} else if len(countToLabel[2]) == 2 {
				return FOURKIND
			} else {
				return THREEKIND
			}
		} else if labelCount['J'] == 1 { //J A B  B D
			if len(countToLabel[4]) != 0 {
				return FIVEKIND
			} else if len(countToLabel[3]) != 0 {
				return FOURKIND
			} else if len(countToLabel[2]) == 2 {
				return FULLHOUSE
			} else if len(countToLabel[2]) != 0 {
				return THREEKIND
			} else if len(countToLabel[1]) != 0 {
				return ONEPAIR
			}
		}
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

func sortHand(hands []Hand, isJoker bool) []Hand {
	sort.Slice(hands, func(i, j int) bool {
		if hands[j].handType == hands[i].handType {
			for idx := range hands[j].cards {
				curCardJ := getCardStrength(hands[j].cards[idx], isJoker)
				curCardI := getCardStrength(hands[i].cards[idx], isJoker)

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
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	hands := createHands(lines)
	handsJoker := createHands(lines)
	for i := range hands {
		handsJoker[i].handType = getHandType(handsJoker[i], true)
		hands[i].handType = getHandType(hands[i], false)
	}
	sortedHands := sortHand(hands, false)
	sortedHandsJoker := sortHand(handsJoker, true)
	winnings := 0
	winningsJoker := 0
	for i, hand := range sortedHands {
		winningsJoker += (sortedHandsJoker[i].bid * (i + 1))
		winnings += (hand.bid * (i + 1))
	}

	fmt.Println(winnings)

	fmt.Println(winningsJoker)

}
