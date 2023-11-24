package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Point3 struct {
	x int
	y int
	z int
}

func getPointsMap(lines []string) (map[Point3]Point3, Point3) {
	points := make(map[Point3]Point3, 0)
	var startingPoint Point3
	for _, line := range lines {
		tokens := strings.Split(line, ",")
		x, err := strconv.Atoi(tokens[0])
		y, err := strconv.Atoi(tokens[1])
		z, err := strconv.Atoi(tokens[2])
		if err != nil {
			log.Fatalf("converting string to num: %s", err)
		}
		startingPoint = Point3{x: x, y: y}
		points[Point3{x: x, y: y, z: z}] = Point3{x: x, y: y, z: z}
	}
	return points, startingPoint
}

func floodFill(pointsMap map[Point3]Point3, point Point3, directions []Point3, visited map[Point3]Point3, distance int) {
	if len(visited) >= distance {
		return
	}
	_, isVisited := visited[Point3{x: point.x, y: point.y, z: point.z}]
	if isVisited {
		return
	}
	_, exists := pointsMap[Point3{x: point.x, y: point.y, z: point.z}]
	if exists {
		return
	}
	visited[Point3{x: point.x, y: point.y, z: point.z}] = Point3{x: point.x, y: point.y, z: point.z}

	for _, direction := range directions {
		nextPoint := Point3{x: point.x + direction.x, y: point.y + direction.y, z: point.z + direction.z}
		floodFill(pointsMap, nextPoint, directions, visited, distance)
	}

}

func checkSurfaceArea(point Point3, pointsMap map[Point3]Point3, directions []Point3, distance int) int {
	surface := 0

	for _, direction := range directions {
		_, exists := pointsMap[Point3{x: point.x + direction.x, y: point.y + direction.y, z: point.z + direction.z}]
		if !exists {
			visited := make(map[Point3]Point3, 0)
			floodFill(pointsMap, Point3{x: point.x + direction.x, y: point.y + direction.y, z: point.z + direction.z}, directions, visited, distance)
			if len(visited) >= distance {
				surface++
			}

		}
	}

	return surface
}

func Day18() {
	//Could improbe by creating max andd min points and then when flood fill check if point is
	//ooutside max and min cube
	lines, err := ReadLines("resources/day18input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	points, _ := getPointsMap(lines)

	directions := make([]Point3, 0)
	directions = append(directions, Point3{x: 1, y: 0, z: 0})
	directions = append(directions, Point3{x: -1, y: 0, z: 0})
	directions = append(directions, Point3{x: 0, y: 1, z: 0})
	directions = append(directions, Point3{x: 0, y: -1, z: 0})
	directions = append(directions, Point3{x: 0, y: 0, z: 1})
	directions = append(directions, Point3{x: 0, y: 0, z: -1})
	total := 0
	distance := 1003
	for k, _ := range points {

		total = total + checkSurfaceArea(k, points, directions, distance)
	}

	fmt.Println("There are ", total)

}
