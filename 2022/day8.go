package main

import (
	"fmt"
	"log"
	"strconv"
)

type Direction struct {
	x int
	y int
}

func isOutOfBounds(x int, y int, length int) bool {
	if y < 0 || x < 0 || y >= length || x >= length {
		return true
	}
	return false
}

func Day8() {
	lines, err := ReadLines("resources/day8input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	var treeLength = len(lines[0])

	treeGrid := make([][]int, treeLength)
	for y, line := range lines {
		for x := range line {
			treeValue, err := strconv.Atoi(line[x : x+1])
			if err != nil {
				log.Fatalf("Converting string to number: %s", err)
			}
			treeGrid[y] = append(treeGrid[y], treeValue)
		}
	}

	var visible int = 0
	directions := make(map[string]Direction, 4)
	directions["UP"] = Direction{x: 0, y: -1}
	directions["RIGHT"] = Direction{x: 1, y: 0}
	directions["DOWN"] = Direction{x: 0, y: 1}
	directions["LEFT"] = Direction{x: -1, y: 0}

	var maxScore int = 0

	for y := 0; y < len(treeGrid); y++ {
		for x := 0; x < len(treeGrid); x++ {
			var tree int = treeGrid[y][x]
			// For each tree check all 4 directions and see if visble
			// Visible means tree in that direction is shorter than current tree
			// If not visible move to next direction
			// If is visbile keep going in that direction to make sure tree is shorter than current

			var directionScore map[string]int = make(map[string]int, 4)
			directionScore["UP"] = 0
			directionScore["RIGHT"] = 0
			directionScore["DOWN"] = 0
			directionScore["LEFT"] = 0
			for compass, direction := range directions {

				var isVisble bool = true
				var adjacenY = y
				var adjacenX = x
				var adjacentTree int = treeGrid[adjacenY][adjacenX]
				for !isOutOfBounds(adjacenX+direction.x, adjacenY+direction.y, treeLength) {
					directionScore[compass]++
					adjacenX += direction.x
					adjacenY += direction.y
					adjacentTree = treeGrid[adjacenY][adjacenX]
					if adjacentTree >= tree {
						isVisble = false
						break
					}
				}
				// 486540 1684
				if isVisble {
					visible++
					//break Un comment to get visibile trees working
				}

			}
			fmt.Println(directionScore, " tree ", tree)
			var curScore int = 1
			for _, treeVal := range directionScore {
				curScore *= treeVal
			}
			if curScore > maxScore {
				maxScore = curScore
			}
		}
	}

	fmt.Println("Visible trees: ", visible)
	fmt.Println("Max scoer is ", maxScore)

}
