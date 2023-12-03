package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

func printSymbols(symbols [][]byte) {
	for y := 0; y < len(symbols); y++ {
		for x := 0; x < len(symbols[0]); x++ {
			fmt.Print(string(symbols[y][x]))
		}
		fmt.Println()
	}
}

type Point struct {
	x int
	y int
}

type Gear struct {
	point    Point
	partNums []*PartNumber
}

type PartNumber struct {
	number       int
	numberPoints []Point
}

func createPartsSymbols(lines []string) ([][]byte, []PartNumber, *map[Point]Gear) {
	gears := make(map[Point]Gear, 0)
	symbols := make([][]byte, len(lines))
	partNumbers := make([]PartNumber, 0)
	//total := 0
	for y, line := range lines {
		symbols[y] = make([]byte, len(line))
		curDigit := ""
		numberPoints := make([]Point, 0)
		for x, _ := range line {
			symbols[y][x] = line[x]
			if line[x] == '*' {
				gears[Point{x: x, y: y}] = Gear{point: Point{x: x, y: y}, partNums: make([]*PartNumber, 0)}
			}
			if line[x] >= '0' && line[x] <= '9' {
				curDigit += string(line[x])
				numberPoints = append(numberPoints, Point{x: x, y: y})
			} else {
				if curDigit != "" {
					num, _ := strconv.Atoi(curDigit)
					partNumbers = append(partNumbers, PartNumber{number: num, numberPoints: numberPoints})
					curDigit = ""
					numberPoints = nil
				}
			}
		}
		if curDigit != "" {
			num, _ := strconv.Atoi(curDigit)
			partNumbers = append(partNumbers, PartNumber{number: num, numberPoints: numberPoints})
			curDigit = ""
			numberPoints = nil
		}

	}
	return symbols, partNumbers, &gears
}

func isPointAdjToSymbol(point Point, symbols [][]byte) bool {
	directions := []Point{Point{0, 1}, Point{1, 1}, Point{1, 0}, Point{1, -1}, Point{0, -1},
		Point{-1, -1}, Point{-1, 0}, Point{-1, 1}}
	isAdj := false
	for _, direction := range directions {
		newPoint := Point{x: direction.x + point.x, y: direction.y + point.y}
		if newPoint.x < 0 || newPoint.x >= len(symbols[0]) || newPoint.y < 0 || newPoint.y >= len(symbols) {
			continue
		}
		symbol := symbols[newPoint.y][newPoint.x]
		isDigit := symbol >= '0' && symbol <= '9'
		if symbol != '.' && !isDigit {
			isAdj = true
			break
		}
	}
	return isAdj
}

func partEquals(part *PartNumber, part2 *PartNumber) bool {

	//fmt.Println(part, part2)
	if part.number != part2.number {
		return false
	}
	if part.number == part2.number {
		for i := 0; i < len(part.numberPoints); i++ {
			if part.numberPoints[i] != part2.numberPoints[i] {
				return false
			}
		}
	}
	return true
}

func containsPartGear(part *PartNumber, gear Gear) bool {

	for _, partGear := range gear.partNums {
		if partEquals(part, partGear) {
			//fmt.Println(part, partGear)
			return true
		}
	}
	return false
}

func isPointAdjToGear(point Point, curPart PartNumber, symbols [][]byte, gears *map[Point]Gear) {
	directions := []Point{Point{0, 1}, Point{1, 1}, Point{1, 0}, Point{1, -1}, Point{0, -1},
		Point{-1, -1}, Point{-1, 0}, Point{-1, 1}}
	for _, direction := range directions {
		newPoint := Point{x: direction.x + point.x, y: direction.y + point.y}
		if newPoint.x < 0 || newPoint.x >= len(symbols[0]) || newPoint.y < 0 || newPoint.y >= len(symbols) {
			continue
		}
		symbol := symbols[newPoint.y][newPoint.x]
		if symbol == '*' {
			gear := (*gears)[newPoint]
			if !containsPartGear(&curPart, gear) {

				gear.partNums = append(gear.partNums, &curPart)
				(*gears)[newPoint] = gear
			}

		}
	}
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	symbols, partNumbers, gears := createPartsSymbols(lines)
	total := 0

	for _, part := range partNumbers {
		isPartNumber := false
		for _, point := range part.numberPoints {
			if isPointAdjToSymbol(point, symbols) {
				isPartNumber = true
			}

			isPointAdjToGear(point, part, symbols, gears)
		}
		if isPartNumber {
			total += part.number
		} else {
		}

	}

	gearRatioSum := 0
	for _, gear := range *gears {
		curGearRatio := 1
		if len(gear.partNums) == 2 {
			for _, part := range gear.partNums {

				curGearRatio = curGearRatio * part.number
			}
			gearRatioSum += curGearRatio
		}

	}
	fmt.Println("Sum is", total)
	fmt.Println("Gear ratui", gearRatioSum)

}
