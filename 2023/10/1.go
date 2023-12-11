package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Point struct {
	x int
	y int
}

var north Point = Point{0, -1}
var east Point = Point{1, 0}
var south Point = Point{0, 1}
var west Point = Point{-1, 0}

var pipeMap map[byte][]Point = map[byte][]Point{
	'|': {north, south},
	'-': {west, east},
	'L': {north, east},
	'J': {north, west},
	'7': {south, west},
	'F': {south, east},
	'S': {east, south},
}

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

func createPipes(lines []string) ([][]byte, Point) {
	pipes := make([][]byte, 0)
	start := Point{0, 0}
	for y, line := range lines {
		newLine := make([]byte, 0)
		for x := range line {
			newLine = append(newLine, line[x])
			if line[x] == 'S' {
				start.x = x
				start.y = y
			}
		}
		pipes = append(pipes, newLine)
	}

	return pipes, start
}

func printPipes(pipes [][]byte, highlight []Point, outside []Point, inside []Point) {
	for y := range pipes {
		for x := range pipes[y] {
			found := false
			for _, point := range highlight {
				if x == point.x && y == point.y {
					fmt.Print("P")
					found = true
				}
			}
			if !found {
				for _, point := range outside {
					if x == point.x && y == point.y {
						fmt.Print("O")
						found = true
					}
				}
			}

			if !found {
				for _, point := range inside {
					if x == point.x && y == point.y {
						fmt.Print("I")
						found = true
					}
				}
			}
			if !found {
				fmt.Print(string(pipes[y][x]))
			}
		}
		fmt.Println()
	}
}

func getNextSteps(point Point, pipe byte) []Point {
	steps := make([]Point, 0)
	for _, step := range pipeMap[pipe] {
		steps = append(steps, Point{point.x + step.x, point.y + step.y})
	}
	return steps
}

//[Language: C++]

// Part 2 solution is area - circumference. But...
// First I had a large brainfart moment by adding +1 to the area for each "left" and "top"
// border element before I realized this is equal to the solution for Part 1 +1
func getFurthestPoint(pipes [][]byte, start Point) (int, []Point) {
	visited := map[Point]int{start: 0}
	connected := []Point{start}
	max := 0
	queue := []Point{start}
	for len(queue) != 0 {
		curPoint := queue[0]
		queue = queue[1:]
		for _, point := range getNextSteps(curPoint, pipes[curPoint.y][curPoint.x]) {
			_, contains := visited[point]
			if !contains {
				stepsTaken := visited[curPoint] + 1
				if stepsTaken > max {
					max = stepsTaken
				}
				visited[point] = stepsTaken
				connected = append(connected, point)
				queue = append(queue, point)
			}
		}
	}
	return max, connected
}

func getOutsideAndInside(pipes [][]byte, path []Point) ([]Point, []Point) {
	outside := make([]Point, 0)
	insde := make([]Point, 0)

	return outside, insde
}

func main() {
	lines, err := ReadLines("test.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	pipes, start := createPipes(lines)
	max, path := getFurthestPoint(pipes, start)
	outside, inside := getOutsideAndInside(pipes, path)
	fmt.Println(max)
	printPipes(pipes, path, outside, inside)

}
