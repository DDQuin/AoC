package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
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

type Game struct {
	id   int
	sets *[]Set
}

type Set struct {
	red   int
	green int
	blue  int
}

func createGames(lines []string) []Game {
	games := make([]Game, 0)

	for _, line := range lines {
		idSplit := strings.Split(strings.Split(line, "Game")[1], ":")
		id, _ := strconv.Atoi(strings.Trim(idSplit[0], " "))
		sets := make([]Set, 0)
		game := Game{id: id, sets: &sets}
		setsList := strings.Split(idSplit[1], ";")
		for _, setString := range setsList {
			set := Set{red: 0, green: 0, blue: 0}
			cubes := strings.Split(setString, ",")
			for _, cube := range cubes {
				cubeSplit := strings.Split(cube, " ")
				color := cubeSplit[2]
				num, err := strconv.Atoi(cubeSplit[1])
				if err != nil {
					log.Fatalf("readLines: %s", err)
				}
				if color == "red" {
					set.red = num
				}
				if color == "blue" {
					set.blue = num
				}
				if color == "green" {
					set.green = num
				}

			}
			sets = append(sets, set)
		}
		games = append(games, game)

	}
	return games
}

func isSetPossible(redMax int, greenMax int, blueMax int, sets []Set) bool {
	for _, set := range sets {
		if set.red > redMax || set.green > greenMax || set.blue > blueMax {
			return false
		}
	}
	return true
}

func getMinSet(sets []Set) Set {
	firstGameSet := sets[0]
	minSet := Set{red: firstGameSet.red, green: firstGameSet.green, blue: firstGameSet.blue}
	for _, set := range sets {
		if set.red > minSet.red {
			minSet.red = set.red
		}
		if set.green > minSet.green {
			minSet.green = set.green
		}
		if set.blue > minSet.blue {
			minSet.blue = set.blue
		}

	}
	return minSet
}

func getIDSumAndPower(redMax int, greenMax int, blueMax int, games []Game) (int, int) {
	total := 0
	powerTotal := 0
	minSets := make([]Set, 0)
	for _, game := range games {
		if isSetPossible(redMax, greenMax, blueMax, *game.sets) {
			total += game.id
		}
		minSets = append(minSets, getMinSet(*game.sets))
	}
	for _, minSet := range minSets {
		powerTotal += minSet.blue * minSet.green * minSet.red
	}
	return total, powerTotal
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	games := createGames(lines)
	total, powerTotal := getIDSumAndPower(12, 13, 14, games)

	fmt.Println("ID sum is", total)
	fmt.Println("Power total  is", powerTotal)
}
