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

func getWaysBeat(time int, distance int) int {
	beat := 0

	for t := 0; t <= time; t++ {
		curDistance := -(t * t) + time*t
		if curDistance > distance {
			beat++
		}
	}

	return beat
}

func getTimesAndDists(lines []string) ([]int, []int) {
	timeString := strings.Split(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ")
	times := make([]int, 0)
	for _, timeS := range timeString {
		if timeS != "" {
			num, err := strconv.Atoi(timeS)
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			times = append(times, num)
		}
	}

	distString := strings.Split(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ")
	distances := make([]int, 0)
	for _, distS := range distString {
		if distS != "" {
			num, err := strconv.Atoi(distS)
			if err != nil {
				log.Fatalf("readLines: %s", err)
			}
			distances = append(distances, num)
		}
	}
	return times, distances
}

func getTimeAndDist(lines []string) (int, int) {
	timeString := strings.ReplaceAll(strings.TrimSpace(strings.Split(lines[0], ":")[1]), " ", "")
	time, _ := strconv.Atoi(timeString)

	distString := strings.ReplaceAll(strings.TrimSpace(strings.Split(lines[1], ":")[1]), " ", "")
	distance, _ := strconv.Atoi(distString)

	return time, distance
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	num := 1
	times, distances := getTimesAndDists(lines)
	time, distance := getTimeAndDist(lines)

	for i := 0; i < len(times); i++ {
		num = num * getWaysBeat(times[i], distances[i])
	}

	fmt.Println("Num is ", num)

	fmt.Println("Big num is ", getWaysBeat(time, distance))

}
