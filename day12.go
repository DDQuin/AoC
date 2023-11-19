package main

import (
	"fmt"
	"log"
	"sort"
)

type GridPoint struct {
	x       int
	y       int
	signal  string
	isEnd   bool
	visited bool
	parent  *GridPoint
}

func isOutOfBoundsGrid(x int, y int, yLength int, xLength int) bool {
	if y < 0 || x < 0 || y >= yLength || x >= xLength {
		return true
	}
	return false
}

func getNextPossibleSteps(point *GridPoint, heightMap [][]GridPoint) []*GridPoint {

	//fmt.Println(heightMap)
	nextSteps := make([]*GridPoint, 0)
	directions := make(map[string]Point, 4)
	directions["U"] = Point{x: 0, y: 1}
	directions["R"] = Point{x: 1, y: 0}
	directions["D"] = Point{x: 0, y: -1}
	directions["L"] = Point{x: -1, y: 0}
	for _, direction := range directions {
		var newPoint Point = Point{x: point.x + direction.x, y: point.y + direction.y}
		if !isOutOfBoundsGrid(newPoint.x, newPoint.y, len(heightMap), len(heightMap[0])) {

			var nextPossiblePoint *GridPoint = &heightMap[newPoint.y][newPoint.x]
			//	fmt.Println(nextPossiblePoint)
			if nextPossiblePoint.signal[0] > point.signal[0] {
				var signalDiff int = int(nextPossiblePoint.signal[0] - point.signal[0])
				if signalDiff > 1 {
					continue
				}
			}
			nextSteps = append(nextSteps, nextPossiblePoint)

		}
	}
	return nextSteps

}

func createHeightMap(lines []string) ([][]GridPoint, []*GridPoint) {
	aRoots := make([]*GridPoint, 0)
	heightMap := make([][]GridPoint, len(lines))

	for y, line := range lines {
		heightMap[y] = make([]GridPoint, len(line))
		for x := range line {
			point := GridPoint{
				x:       x,
				y:       y,
				visited: false,
				signal:  line[x : x+1],
				isEnd:   false,
				parent:  nil,
			}
			if line[x:x+1] == "S" {
				point.signal = "a"
				//aRoots = append(aRoots, &point)
			} else if line[x:x+1] == "E" {
				point.signal = "z"
				point.isEnd = true
			}
			if point.signal == "a" {
				aRoots = append(aRoots, &point)
			}
			heightMap[y][x] = point
		}
	}
	return heightMap, aRoots
}

func bfsGrid(heightMap [][]GridPoint, root *GridPoint) int {
	var steps int = -1
	queue := make([]GridPoint, 0)
	queue = append(queue, *root)
	var end GridPoint
	for len(queue) != 0 {
		vPoint := queue[0]
		queue = queue[1:]
		if vPoint.isEnd {
			steps = 0
			end = vPoint
			break
		}
		for _, point := range getNextPossibleSteps(&vPoint, heightMap) {
			if !point.visited {
				point.visited = true
				point.parent = &vPoint
				queue = append(queue, *point)
			}
		}

	}

	for end.parent != nil {
		//fmt.Println("Step ", end)
		end = *end.parent
		steps++
	}
	//fmt.Println("Final Step ", end, " steps", steps)
	return steps
}

func Day12() {

	lines, err := ReadLines("resources/day12input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	heightMap, aRoots := createHeightMap(lines)
	stepsList := make([]int, 0)
	for i := 0; i < len(aRoots); i++ {
		root := aRoots[i]
		heightMap, _ = createHeightMap(lines)
		step := bfsGrid(heightMap, root)
		if step != -1 {
			stepsList = append(stepsList, step)
		}
	}
	sort.Ints(stepsList)
	fmt.Println(stepsList)

}
