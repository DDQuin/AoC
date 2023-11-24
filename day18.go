package main

import (
	"fmt"
	"log"
	"math"
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

func floodFillRange(pointsMap map[Point3]Point3, point Point3, directions []Point3, visited map[Point3]Point3, max int, min int) bool {
	_, isVisited := visited[Point3{x: point.x, y: point.y, z: point.z}]
	if isVisited {
		return false
	}
	_, exists := pointsMap[Point3{x: point.x, y: point.y, z: point.z}]
	if exists {
		return false
	}

	if point.x <= min || point.y <= min || point.z <= min {
		return true
	}

	if point.x >= max || point.y >= max || point.z >= max {
		return true
	}
	visited[Point3{x: point.x, y: point.y, z: point.z}] = Point3{x: point.x, y: point.y, z: point.z}

	for _, direction := range directions {
		nextPoint := Point3{x: point.x + direction.x, y: point.y + direction.y, z: point.z + direction.z}
		result := floodFillRange(pointsMap, nextPoint, directions, visited, max, min)
		if result {
			return result
		}
	}
	return false

}

func getMinMaxPoints(pointMap map[Point3]Point3) (int, int) {
	maxX := math.MinInt32
	maxY := math.MinInt32
	maxZ := math.MinInt32
	minX := math.MaxInt32
	minY := math.MaxInt32
	minZ := math.MaxInt32
	for point, _ := range pointMap {
		if point.x < minX {
			minX = point.x
		}
		if point.y < minY {
			minY = point.y
		}
		if point.z < minZ {
			minZ = point.z
		}

		if point.x > maxX {
			maxX = point.x
		}
		if point.y > maxY {
			maxY = point.y
		}
		if point.z > maxZ {
			maxZ = point.z
		}

	}

	min := minX
	if minY < min {
		min = minY
	}
	if minZ < min {
		min = minZ
	}

	max := maxX
	if maxY > max {
		max = maxY
	}
	if maxZ > max {
		max = maxZ
	}
	return max, min
}

func checkSurfaceArea(point Point3, pointsMap map[Point3]Point3, directions []Point3, max int, min int) int {
	surface := 0

	for _, direction := range directions {
		_, exists := pointsMap[Point3{x: point.x + direction.x, y: point.y + direction.y, z: point.z + direction.z}]
		if !exists {
			visited := make(map[Point3]Point3, 0)
			result := floodFillRange(pointsMap, Point3{x: point.x + direction.x, y: point.y + direction.y, z: point.z + direction.z}, directions, visited, max, min)
			if result {
				surface++
			}

		}
	}

	return surface
}

func Day18() {
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
	max, min := getMinMaxPoints(points)

	for k, _ := range points {

		total = total + checkSurfaceArea(k, points, directions, max, min)
	}

	fmt.Println("There are ", total)

}
