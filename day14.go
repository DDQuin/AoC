package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func getPaths(lines []string) ([][]Point, int, int, int) {
	paths := make([][]Point, 0)
	maxY := 0
	maxX := 0
	minX := 10000
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		path := make([]Point, 0)
		for _, token := range tokens {
			if token != "->" {
				pointSplit := strings.Split(token, ",")
				x, err := strconv.Atoi(pointSplit[0])
				y, err := strconv.Atoi(pointSplit[1])
				if y > maxY {

					maxY = y
				}
				if x > maxX {
					maxX = x
				}
				if x < minX {
					minX = x
				}
				if err != nil {
					log.Fatalf("error converting string to num: %s", err)

				}
				path = append(path, Point{x: x, y: y})
			}
		}
		paths = append(paths, path)
	}
	return paths, maxX, maxY, minX
}

func printCave(cave [][]byte, minX int) {
	for y := 0; y < len(cave); y++ {
		for x := minX - len(cave); x < len(cave[0]); x++ {
			fmt.Print(string(cave[y][x]))
		}

		fmt.Println()
	}
}

func getMagnitude(point Point) int {
	length := point.x + point.y
	if length < 0 {
		length = length * -1
	}
	return length
}

func normalisePoint(point Point, magnitude int) Point {
	return Point{x: point.x / magnitude, y: point.y / magnitude}
}

func createLine(cave [][]byte, point1 Point, point2 Point) {
	pointDiff := Point{x: point2.x - point1.x, y: point2.y - point1.y}
	mag := getMagnitude(pointDiff)

	normal := normalisePoint(pointDiff, mag)
	for i := 0; i < mag+1; i++ {
		newPoint := Point{x: point1.x + (normal.x * i), y: point1.y + (normal.y * i)}
		cave[newPoint.y][newPoint.x] = '#'
	}
}

func createCave(paths [][]Point, width int, height int) [][]byte {
	cave := make([][]byte, height)
	for y := 0; y < height; y++ {
		cave[y] = make([]byte, width+height)
		for x := 0; x < width+height; x++ {
			cave[y][x] = '.'
		}
	}

	for _, path := range paths {
		pointStart := path[0]
		for i := 1; i < len(path); i++ {
			currentPoint := path[i]
			createLine(cave, pointStart, currentPoint)
			pointStart = currentPoint
		}

	}
	//createLine(cave, Point{x: 0, y: height - 1}, Point{x: width - 1, y: height - 1})
	return cave
}

func simulateCaveSand(cave [][]byte) int {
	sand := Point{500, 0}
	sandUnits := 0
	for sand.y != len(cave)-1 {
		if cave[sand.y+1][sand.x] == '.' {
			sand = Point{x: sand.x, y: sand.y + 1}
		} else if cave[sand.y+1][sand.x-1] == '.' {
			sand = Point{x: sand.x - 1, y: sand.y + 1}
		} else if cave[sand.y+1][sand.x+1] == '.' {
			sand = Point{x: sand.x + 1, y: sand.y + 1}
		} else {
			sandUnits++
			cave[sand.y][sand.x] = 'o'
			sand = Point{500, 0}
		}
	}
	return sandUnits
}

func Day14() {
	lines, err := ReadLines("resources/day14test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	paths, maxX, maxY, minX := getPaths(lines)

	cave := createCave(paths, maxX+1, maxY+1)
	printCave(cave, minX)

	//sand := simulateCaveSand(cave)

	//fmt.Println("There are", sand, "units of sand before it flows forever")

}
