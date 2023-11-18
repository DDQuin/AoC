package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func drawScreen(screen [][]bool) {
	for y := 0; y < len(screen); y++ {
		for x := 0; x < 40; x++ {
			if screen[y][x] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func convertCycleToPixel(cycle int) (x int, y int) {
	var xPixel int = (cycle - 1) % 40
	var yPixel int = (cycle - 1) / 40
	return xPixel, yPixel
}

func Day10() {

	var cycle int = 1
	var X int = 1
	var sum int = 0
	var isAdding bool = false
	var instructionPointer int = 0
	screen := make([][]bool, 6)
	for y := range screen {
		screen[y] = make([]bool, 40)
	}
	lines, err := ReadLines("resources/day10input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	for instructionPointer < len(lines) {

		pixelX, pixelY := convertCycleToPixel(cycle)
		if pixelX >= X-1 && pixelX <= X+1 {
			screen[pixelY][pixelX] = true
		} else {
			screen[pixelY][pixelX] = false
		}
		if cycle == 20 || cycle == 60 || cycle == 100 || cycle == 140 || cycle == 180 || cycle == 220 {
			sum += cycle * X
		}
		var instructionLine string = lines[instructionPointer]
		tokens := strings.Split(instructionLine, " ")
		var instruction string = tokens[0]
		//fmt.Println("Instriction", instruction, "Cycle", cycle, "IP", instructionPointer, "X", X)
		if instruction == "noop" {
			instructionPointer++
		} else if instruction == "addx" {
			if isAdding {
				value, err := strconv.Atoi(tokens[1])
				if err != nil {
					log.Fatalf("error converting string to num %s", err)
				}
				X += value
				//fmt.Println("ADDED Instriction", instruction, "Cycle", cycle, "IP", instructionPointer, "X", X)
				instructionPointer++
				isAdding = false
			} else {
				//fmt.Println("starting to add Instriction", instruction, "Cycle", cycle, "IP", instructionPointer, "X", X)
				isAdding = true
			}
		} else {
			log.Fatalf("unsupported operation! %s", instruction)
		}

		cycle++

	}

	drawScreen(screen)

	fmt.Println("FINAL SUM IS", sum)

}
