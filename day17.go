package main

import (
	"fmt"
	"log"
)

func createDirectionList(line string) []Point {
	directionList := make([]Point, 0)
	for _, char := range line {
		if char == '>' {
			directionList = append(directionList, Point{x: 1, y: 0})
		} else if char == '<' {
			directionList = append(directionList, Point{x: -1, y: 0})
		}
	}

	return directionList
}

func createPossibleRocks() [][]Point {
	rocks := make([][]Point, 0)

	lineRock := []Point{{x: 2, y: 0}, {x: 3, y: 0}, {x: 4, y: 0}, {x: 5, y: 0}}
	plusRock := []Point{{x: 3, y: 2}, {x: 2, y: 1}, {x: 3, y: 1}, {x: 4, y: 1}, {x: 3, y: 0}}
	lRock := []Point{{x: 4, y: 2}, {x: 4, y: 1}, {x: 2, y: 0}, {x: 3, y: 0}, {x: 4, y: 0}}
	tallRock := []Point{{x: 2, y: 3}, {x: 2, y: 2}, {x: 2, y: 1}, {x: 2, y: 0}}
	boxRock := []Point{{x: 2, y: 1}, {x: 3, y: 1}, {x: 2, y: 0}, {x: 3, y: 0}}

	rocks = append(rocks, lineRock)
	rocks = append(rocks, plusRock)
	rocks = append(rocks, lRock)
	rocks = append(rocks, tallRock)
	rocks = append(rocks, boxRock)

	return rocks
}

func getUnitsTall(placedRocks [][]Point) int {
	tallestY := 0
	for _, rock := range placedRocks {
		curY := rock[0].y
		if curY > tallestY {
			tallestY = curY
		}
	}
	return tallestY + 1
}

func getNextRock(newRock []Point, placedRocks [][]Point) []Point {
	lastRockY := -1
	if len(placedRocks) != 0 {
		lastRock := placedRocks[len(placedRocks)-1]
		lastRockY = lastRock[0].y
	}
	nextRock := make([]Point, 0)
	// First index of eaech rock is highest tpoint
	for _, point := range newRock {
		nextRock = append(nextRock, Point{point.x, point.y + lastRockY + 4})
	}
	return nextRock
}

func moveRock(curRock []Point, move Point) []Point {
	newRock := make([]Point, 0)
	for _, point := range curRock {
		newRock = append(newRock, Point{x: point.x + move.x, y: point.y + move.y})
	}
	return newRock
}

func isRockTouching(rock []Point, placedRocks [][]Point) bool {
	for _, point := range rock {
		if point.x < 0 || point.x > 6 || point.y < 0 {
			return true
		}

		for _, rockd := range placedRocks {
			for _, pointRock := range rockd {
				if point.x == pointRock.x && point.y == pointRock.y {
					return true
				}

			}
		}
	}
	return false
}

func Day17() {
	lines, err := ReadLines("resources/day17input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	jets := createDirectionList(lines[0])
	jetIndex := 0
	rocks := createPossibleRocks()
	rockIndex := 0
	placedRocks := make([][]Point, 0)

	for len(placedRocks) < 2022 {
		curRock := getNextRock(rocks[rockIndex%len(rocks)], placedRocks)
		rockIndex++
		isRested := false
		for !isRested {
			curJet := jets[jetIndex%len(jets)]
			jetRock := moveRock(curRock, curJet)
			//fmt.Println("try direction", curJet, jetRock)
			if !isRockTouching(jetRock, placedRocks) {
				curRock = jetRock
				//fmt.Println("move direction")
			}
			jetIndex++
			nextRock := moveRock(curRock, Point{x: 0, y: -1})
			if !isRockTouching(nextRock, placedRocks) {
				curRock = nextRock

				//	fmt.Println("rock is falling ", curRock)
			} else {
				isRested = true
			}
		}
		fmt.Println("Placed rrock at ", curRock)
		placedRocks = append(placedRocks, curRock)
	}

	fmt.Println("UNits tall were ", getUnitsTall(placedRocks))

}
