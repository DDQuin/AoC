package main

import (
	"fmt"
	"log"
	"strconv"
)

func printBoard(board [][]byte, player Point, direction int) {
	for y := 0; y < len(board); y++ {
		for x := 0; x < len(board[0]); x++ {
			if player.x == x && player.y == y {
				if direction == 0 {
					fmt.Print(">")
				} else if direction == 1 {
					fmt.Print("v")
				} else if direction == 2 {
					fmt.Print("<")
				} else if direction == 3 {
					fmt.Print("^")
				}
			} else {
				fmt.Print(string(board[y][x]))
			}
		}
		fmt.Println()
	}
}

func createBoard(lines []string) ([][]byte, string, Point) {
	board := make([][]byte, 0)

	start := Point{x: 0, y: 0}
	biggestLen := 0
	pathLine := ""
	for i, line := range lines {
		length := len(line)
		if length > biggestLen {
			biggestLen = length
		}
		if length != 0 { // Is board
			board = append(board, []byte(line))
		} else {
			pathLine = lines[i+1]
			break
		}
	}

	for y := 0; y < len(board); y++ { // Make sure all lines are same length as biggest
		for len(board[y]) < biggestLen {
			board[y] = append(board[y], ' ')
		}
	}

	for x := 0; x < biggestLen; x++ {
		if board[0][x] == '.' {
			start.x = x
			start.y = 0
			break
		}
	}
	return board, pathLine, start
}

func getInstructionList(path string) []string {
	instructionList := make([]string, 0)
	numString := ""
	for i := 0; i < len(path); i++ {
		curChar := path[i]
		if curChar != 'R' && curChar != 'L' {
			numString += string(curChar)
		} else {
			instructionList = append(instructionList, numString)
			numString = ""
			instructionList = append(instructionList, string(curChar))
		}
	}
	if numString != "" {
		instructionList = append(instructionList, numString)
	}
	return instructionList
}

func getWrapTile(board [][]byte, player Point, direction int) Point {
	oppositeDirection := (direction + 2) % 4
	curPoint := player
	if oppositeDirection == 0 {
		nextPoint := Point{x: player.x + 1, y: player.y}
		for nextPoint.x < len(board[0]) && board[nextPoint.y][nextPoint.x] != ' ' {
			curPoint = Point{x: nextPoint.x, y: nextPoint.y}
			nextPoint.x++
		}

	} else if oppositeDirection == 1 {
		nextPoint := Point{x: player.x, y: player.y + 1}
		for nextPoint.y < len(board) && board[nextPoint.y][nextPoint.x] != ' ' {
			curPoint = Point{x: nextPoint.x, y: nextPoint.y}
			nextPoint.y++
		}
	} else if oppositeDirection == 2 {

		nextPoint := Point{x: player.x - 1, y: player.y}
		for nextPoint.x >= 0 && board[nextPoint.y][nextPoint.x] != ' ' {
			curPoint = Point{x: nextPoint.x, y: nextPoint.y}
			nextPoint.x--
		}

	} else if oppositeDirection == 3 {
		nextPoint := Point{x: player.x, y: player.y - 1}
		for nextPoint.y >= 0 && board[nextPoint.y][nextPoint.x] != ' ' {
			curPoint = Point{x: nextPoint.x, y: nextPoint.y}
			nextPoint.y--
		}
	}

	return curPoint
}

func getNextStepBoard(board [][]byte, player Point, direction int) Point {
	next := player

	if direction == 0 {
		next = Point{x: player.x + 1, y: player.y}
	} else if direction == 1 {
		next = Point{x: player.x, y: player.y + 1}
	} else if direction == 2 {
		next = Point{x: player.x - 1, y: player.y}
	} else if direction == 3 {
		next = Point{x: player.x, y: player.y - 1}
	}

	//Wrap arround
	if next.x >= len(board[0]) || next.x < 0 || next.y >= len(board) || next.y < 0 || board[next.y][next.x] == ' ' {
		next = getWrapTile(board, player, direction)
	}
	if board[next.y][next.x] == '#' {
		next = player
	}
	return next
}

func Day22() {
	lines, err := ReadLines("resources/day22input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	direction := 0
	board, path, player := createBoard(lines)
	instructions := getInstructionList(path)
	printBoard(board, player, direction)
	//fmt.Println(instructions)

	for _, instruction := range instructions {
		if instruction == "L" {
			direction = direction - 1
			if direction == -1 {
				direction = 3
			}
		} else if instruction == "R" {
			direction = (direction + 1) % 4

		} else {
			steps, _ := strconv.Atoi(instruction)
			for i := 0; i < steps; i++ {
				nextStep := getNextStepBoard(board, player, direction)
				//fmt.Println("Next step is", nextStep)
				player = nextStep
			}
		}
	}
	printBoard(board, player, direction)
	fmt.Println("Row", player.y+1, "Column", player.x+1, "Direction", direction)
	fmt.Println("Password is", (1000*(player.y+1))+(4*(player.x+1))+direction)

}
