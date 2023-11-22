package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
)

func getSensorBeaconMap(lines []string) (map[Point]Point, int, int, int, int) {
	sensorMap := make(map[Point]Point)
	minX, minY := 100000000, 100000000
	maxX, maxY := -1000, -1000
	for _, line := range lines {
		tokens := strings.Split(line, " ")
		sensorXToken := tokens[2]
		sensorXValue := strings.Split(sensorXToken, "=")[1]
		sensorXValue = sensorXValue[0 : len(sensorXValue)-1]

		sensorYToken := tokens[3]
		sensorYValue := strings.Split(sensorYToken, "=")[1]
		sensorYValue = sensorYValue[0 : len(sensorYValue)-1]

		sX, err := strconv.Atoi(sensorXValue)
		sY, err := strconv.Atoi(sensorYValue)

		beaconXToken := tokens[8]
		beaconXValue := strings.Split(beaconXToken, "=")[1]
		beaconXValue = beaconXValue[0 : len(beaconXValue)-1]

		beaconYToken := tokens[9]
		beaconYValue := strings.Split(beaconYToken, "=")[1]

		bX, err := strconv.Atoi(beaconXValue)
		bY, err := strconv.Atoi(beaconYValue)

		if bX < minX {
			minX = bX
		}
		if sX < minX {
			minX = sX
		}
		if bY < minY {
			minY = bY
		}
		if sY < minY {
			minY = sY
		}

		if bX > maxX {
			maxX = bX
		}
		if sX > maxX {
			maxX = sX
		}
		if bY > maxY {
			maxY = bY
		}
		if sY > maxY {
			maxY = sY
		}

		if err != nil {
			log.Fatalf("error converting: %s", err)
		}

		sensor := Point{x: sX, y: sY}
		beacon := Point{x: bX, y: bY}
		sensorMap[sensor] = beacon
	}

	return sensorMap, minX, minY, maxX, maxY

}

type Zone struct {
	zone [][]byte
	minX int
	minY int
	maxX int
	maxY int
}

func (z *Zone) print() {
	for y := 0; y < len(z.zone); y++ {
		for x := 0; x < len(z.zone[0]); x++ {
			fmt.Print(string(z.zone[y][x]))
		}
		fmt.Println()
	}
}

func (z *Zone) getPoint(point Point) Point {
	return Point{x: point.x - z.minX, y: point.y - z.minY}
}

func (z *Zone) isOOB(point Point) bool {
	if point.x < 0 || point.x >= len(z.zone[0]) || point.y < 0 || point.y >= len(z.zone) {
		return true
	}
	return false
}

func createZone(sensorMap map[Point]Point, minX int, minY int, maxX int, maxY int) Zone {
	width := maxX - minX + 1
	height := maxY - minY + 1

	zone := make([][]byte, height)
	for y := 0; y < height; y++ {
		zone[y] = make([]byte, width)
		for x := 0; x < width; x++ {
			zone[y][x] = '.'
		}
	}
	zoneMain := Zone{zone: zone, minX: minX, minY: minY, maxX: maxX, maxY: maxY}

	for sensor, beacon := range sensorMap {
		sensorPoint := zoneMain.getPoint(sensor)
		beaconPoint := zoneMain.getPoint(beacon)

		zoneMain.zone[sensorPoint.y][sensorPoint.x] = 'S'
		zoneMain.zone[beaconPoint.y][beaconPoint.x] = 'B'
	}
	return zoneMain
}

func getManDist(point1 Point, point2 Point) int {
	return int(math.Abs(float64(point1.x-point2.x)) + math.Abs(float64(point1.y-point2.y)))
}

func markSensorRange(zone *Zone, sensor Point, manhattan int) {
	realPoint := zone.getPoint(sensor)

	for x := -manhattan; x <= manhattan; x++ {
		for y := -manhattan; y <= manhattan; y++ {
			curPoint := Point{x: realPoint.x + x, y: realPoint.y + y}
			if !zone.isOOB(curPoint) && getManDist(realPoint, curPoint) <= manhattan {
				if zone.zone[curPoint.y][curPoint.x] != 'S' && zone.zone[curPoint.y][curPoint.x] != 'B' {
					zone.zone[curPoint.y][curPoint.x] = '#'
				}
			}
		}
	}
}

func countNoBeacon(y int, zone *Zone) int {
	realPoint := zone.getPoint(Point{x: 0, y: y})
	count := 0
	for _, signal := range zone.zone[realPoint.y] {
		if signal == '#' {
			count++
		}
	}
	return count
}

func isNoBeacon(point Point, sensorMap map[Point]Point) bool {
	for sensor, beacon := range sensorMap {
		if beacon.x == point.x && beacon.y == point.y {
			return true
		}
		if sensor.x == point.x && sensor.y == point.y {
			return true
		}
	}

	for sensor, beacon := range sensorMap {
		dist := getManDist(sensor, beacon)
		curDist := getManDist(sensor, point)
		if curDist <= dist {
			return true
		}
	}
	return false
}

func sensorsClose(point Point, sensorMap map[Point]Point, minX int, maxX int, minY int, maxY int) int {
	if point.x < minX || point.x > maxX || point.y < minY || point.y > maxY {
		return 50
	}
	count := 0
	for sensor, beacon := range sensorMap {
		dist := getManDist(sensor, beacon)
		curDist := getManDist(sensor, point)
		if curDist <= dist {
			count++
		}
	}
	return count
}

func countBeacon(y int, minX int, maxX int, sensorMap map[Point]Point) int {
	count := 0
	for i := minX * 2; i <= maxX*2; i++ {
		curPoint := Point{x: i, y: y}
		if isNoBeacon(curPoint, sensorMap) {
			count++
		}
	}
	return count
}

func findBeaconPos(minX int, maxX int, minY int, maxY int, sensorMap map[Point]Point) Point {

	for y := minY; y < maxY; y++ {
		for x := minX; x < maxX; x++ {
			curPoint := Point{x: 0, y: y}
			if !isNoBeacon(curPoint, sensorMap) {
				return curPoint
			}
		}

	}

	return Point{x: -1, y: -1}
}

func mazeFind(curPoint Point, minX int, maxX int, minY int, maxY int, sensorMap map[Point]Point) Point {
	if !isNoBeacon(curPoint, sensorMap) {
		return curPoint
	}

	points := make([]Point, 0)
	upPoint := Point{x: curPoint.x, y: curPoint.y - 1}
	downPoint := Point{x: curPoint.x, y: curPoint.y + 1}
	rightPoint := Point{x: curPoint.x + 1, y: curPoint.y}
	leftPoint := Point{x: curPoint.x - 1, y: curPoint.y}

	points = append(points, upPoint)
	points = append(points, downPoint)
	points = append(points, rightPoint)
	points = append(points, leftPoint)

	bestPoint := points[0]
	bestScore := sensorsClose(bestPoint, sensorMap, minX, maxX, minY, maxY)
	for i := 1; i < len(points); i++ {
		score := sensorsClose(points[i], sensorMap, minX, maxX, minY, maxY)
		if score < bestScore {
			bestScore = score
			bestPoint = points[i]
		}
	}
	return mazeFind(bestPoint, minX, maxX, minY, maxY, sensorMap)

}

func sensorContainsNon(sensor Point, manhattan int, sensorMap map[Point]Point, min int, max int) Point {
	point := Point{x: -1, y: -1}
	fmt.Println(manhattan)
	for x := -manhattan; x <= manhattan; x++ {
		for y := -manhattan; y <= manhattan; y++ {
			curPoint := Point{x: sensor.x + x, y: sensor.y + y}
			isBounds := curPoint.x >= min && curPoint.x <= max && curPoint.y >= min && curPoint.y <= max
			if isBounds && !isNoBeacon(curPoint, sensorMap) {
				return curPoint
			}
		}
	}
	return point
}

func getPrem(sensor Point, man int) []Point {
	points := make([]Point, 0)
	start := Point{x: sensor.x, y: sensor.y - man}
	points = append(points, start)
	for i := 0; i < man; i++ {
		start = Point{x: start.x + 1, y: start.y + 1}
		points = append(points, start)
	}
	for i := 0; i < man; i++ {
		start = Point{x: start.x - 1, y: start.y + 1}
		points = append(points, start)
	}

	for i := 0; i < man; i++ {
		start = Point{x: start.x - 1, y: start.y - 1}
		points = append(points, start)
	}

	for i := 0; i < man; i++ {
		start = Point{x: start.x + 1, y: start.y - 1}
		points = append(points, start)
	}
	return points
}

func Day15() {
	lines, err := ReadLines("resources/day15input.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	sensorMap, _, _, _, _ := getSensorBeaconMap(lines)

	min := 0
	max := 4000000

	for sensor, beacon := range sensorMap {
		dist := getManDist(sensor, beacon)
		points := getPrem(sensor, dist+1)

		for _, point := range points {
			isBounds := point.x >= min && point.x <= max && point.y >= min && point.y <= max
			if isBounds && !isNoBeacon(point, sensorMap) {
				fmt.Println("Pog ", point)
				fmt.Println("Pog ", point.x*4_000_000+point.y)
				return
			}
		}
	}

}
