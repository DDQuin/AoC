package main

import (
	"errors"
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x int
	y int
}

func isAdjacent(head *Point, tail *Point) bool {

	var xDiff int = head.x - tail.x
	var yDiff int = head.y - tail.y
	if math.Abs(float64(xDiff)) >= 2 || math.Abs(float64(yDiff)) >= 2 {
		return false
	}

	return true
}

func getNextStep(head *Point, tail *Point) (Point, error) {
	var xDiff int = head.x - tail.x
	var yDiff int = head.y - tail.y
	//All cases when ddierrctly up/down/left/right
	if xDiff == 0 {
		if yDiff > 0 {
			return Point{x: tail.x, y: tail.y + yDiff - 1}, nil
		} else {
			return Point{x: tail.x, y: tail.y + yDiff + 1}, nil
		}
	}
	if yDiff == 0 {
		if xDiff > 0 {
			return Point{x: tail.x + xDiff - 1, y: tail.y}, nil
		} else {
			return Point{x: tail.x + xDiff + 1, y: tail.y}, nil
		}
	}
	// If new is on right, so goo diagonally up or down to right
	if xDiff > 0 {
		if yDiff > 0 {
			return Point{x: tail.x + 1, y: tail.y + 1}, nil
		} else {
			return Point{x: tail.x + 1, y: tail.y - 1}, nil
		}
	}
	// If new step is left so go diagonally up or down to left
	if xDiff < 0 {
		if yDiff > 0 {
			return Point{x: tail.x - 1, y: tail.y + 1}, nil
		} else {
			return Point{x: tail.x - 1, y: tail.y - 1}, nil
		}
	}

	return Point{x: 99, y: 99}, errors.New("Something wrrong")
}

func hasVisitedPoint(tail *Point, visited *[]Point) bool {
	for _, point := range *visited {
		if point.x == tail.x && point.y == tail.y {
			return true
		}
	}
	return false
}

func Day9() {
	var numKnots int = 10
	lines, err := ReadLines("resources/day9input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	var visitedPositions []Point
	directions := make(map[string]Point, 4)
	directions["U"] = Point{x: 0, y: 1}
	directions["R"] = Point{x: 1, y: 0}
	directions["D"] = Point{x: 0, y: -1}
	directions["L"] = Point{x: -1, y: 0}
	knots := make([]Point, numKnots)
	for i := range knots {
		knots[i] = Point{x: 0, y: 0}
	}
	visitedPositions = append(visitedPositions, Point{knots[numKnots-1].x, knots[numKnots-1].y})

	for _, line := range lines {
		tokens := strings.Split(line, " ")
		var directionString string = tokens[0]
		var nextStep Point = directions[directionString]

		steps, err := strconv.Atoi(tokens[1])
		if err != nil {
			log.Fatalf("converting string to num : %s", err)
		}

		fmt.Println("Direction", directionString, "Steps:", steps)

		for step := 0; step < steps; step++ {
			knots[0].x += nextStep.x
			knots[0].y += nextStep.y
			for i := 0; i < numKnots-1; i++ {
				if !isAdjacent(&knots[i], &knots[i+1]) {
					newPoint, err := getNextStep(&knots[i], &knots[i+1])
					if err != nil {
						log.Fatalf("wrong next step: %s", err)
					}
					knots[i+1].x = newPoint.x
					knots[i+1].y = newPoint.y

				}
			}
			if !hasVisitedPoint(&knots[numKnots-1], &visitedPositions) {
				visitedPositions = append(visitedPositions, Point{knots[numKnots-1].x, knots[numKnots-1].y})

			}
			fmt.Println("Head ", knots[0])
			fmt.Println("Tail ", knots[1])
			fmt.Println("----")
		}

	}

	fmt.Println("Tail Visited", len(visitedPositions), "unique spaces")

}
