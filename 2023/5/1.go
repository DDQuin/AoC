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

type MapConvert struct {
	destination int
	source      int
	rangeLen    int
}

func (mapConvert *MapConvert) mapSeed(seed int) int {
	if seed >= mapConvert.source && seed < mapConvert.source+mapConvert.rangeLen {
		return seed - mapConvert.source + mapConvert.destination
	}
	return -1
}

func createSeedRange(start int, rangeNum int) []int {
	seeds := make([]int, 0)
	for i := 0; i < rangeNum; i++ {
		seeds = append(seeds, start+i)
	}
	return seeds
}

func createMapsSeeds(lines []string) (map[string][]MapConvert, []int, []string) {
	seeds := make([]int, 0)
	maps := make(map[string][]MapConvert, 0)
	convertOrder := make([]string, 0)
	seedNums := strings.Split(strings.Trim(strings.Split(lines[0], ":")[1], " "), " ")
	for i := 0; i < len(seedNums); i += 2 {
		fmt.Println(i)
		//fmt.Println(seedNums[i], seedNums[i+1])
		seedNumStart, _ := strconv.Atoi(seedNums[i])
		seedNumRange, _ := strconv.Atoi(seedNums[i+1])
		seeds = append(seeds, createSeedRange(seedNumStart, seedNumRange)...)
	}
	// for _, num := range seedNums {
	// 	seedNum, _ := strconv.Atoi(num)
	// 	seeds = append(seeds, seedNum)
	// }

	curMapName := ""
	curMapConverts := make([]MapConvert, 0)
	for i := 2; i < len(lines); i++ {
		curLine := lines[i]
		if len(curLine) == 0 {
			maps[curMapName] = curMapConverts
			curMapName = ""
			curMapConverts = nil

		} else {
			splitLine := strings.Split(curLine, " ")
			if splitLine[1] == "map:" {
				curMapName = splitLine[0]
				convertOrder = append(convertOrder, splitLine[0])
			} else {
				destNum, err := strconv.Atoi(splitLine[0])
				sourceNum, err := strconv.Atoi(splitLine[1])
				rangeNum, err := strconv.Atoi(splitLine[2])
				if err != nil {
					log.Fatalf("Number converting %s", err)
				}
				curMapConverts = append(curMapConverts, MapConvert{destNum, sourceNum, rangeNum})
			}
		}

	}
	maps[curMapName] = curMapConverts
	curMapName = ""
	curMapConverts = nil
	return maps, seeds, convertOrder
}

func main() {
	lines, err := ReadLines("input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}

	maps, seeds, convertOrder := createMapsSeeds(lines)
	//fmt.Println(seeds)
	//fmt.Println(maps)
	//.Println(convertOrder)

	for _, convertString := range convertOrder {
		//fmt.Println("Stage", convertString)
		for i, seed := range seeds {
			//oldSeed := seed
			for _, convert := range maps[convertString] {
				newPossibleMap := convert.mapSeed(seed)
				if newPossibleMap != -1 {
					seeds[i] = newPossibleMap
					break
				}
			}
			//fmt.Print("Converting ", oldSeed, " to ", seeds[i], " ")
		}
		//fmt.Println()
	}

	lowestSeed := seeds[0]
	for _, seed := range seeds {
		if seed < lowestSeed {
			lowestSeed = seed
		}
		//fmt.Println("Final seed", seed)
	}

	fmt.Println("Lowest seed is", lowestSeed)

}
