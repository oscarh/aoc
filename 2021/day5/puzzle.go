package main

import (
	"fmt"
	"strings"
	"strconv"

	"github.com/oscarh/aoc/util"
)

type point struct {
	x int
	y int
}

func parsePoint(startend string) point {
	xy := strings.Split(startend, ",")
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		panic("Invalid input")
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		panic("Invalid input")
	}

	return point{x, y}
}

func parseLine(line string) (point, point) {
	points := strings.Split(line, " -> ")
	start := parsePoint(points[0])
	end := parsePoint(points[1])
	return start, end
}

func delta(s, e int) int {
	if s < e {
		return 1
	} else if s > e {
		return -1
	} else {
		return 0
	}
}

func deltaPoint(start, end point) (int, int) {
	return delta(start.x, end.x), delta(start.y, end.y)
}

func addLine(floor map[point]int, start, end point) int {
	newOverlapping := 0
	dx, dy := deltaPoint(start, end)

	x := start.x
	y := start.y
	for {
		floor[point{x,y}] += 1
		if floor[point{x,y}] == 2 {
			newOverlapping += 1
		}

		if x == end.x && y == end.y {
			break
		}

		x += dx
		y += dy
	}

	return newOverlapping 
}

func part1() int {
	input := util.LoadInput()
	seaFloor := map[point]int{}
	overlapping := 0
	for _, line := range input {
		start, end := parseLine(line)
		if start.x != end.x &&  start.y != end.y {
			// Don't consider these in part 1
			continue
		}
		overlapping += addLine(seaFloor, start, end)
	}

	return overlapping
}

func part2() int {
	input := util.LoadInput()
	seaFloor := map[point]int{}
	overlapping := 0
	for _, line := range input {
		start, end := parseLine(line)
		overlapping += addLine(seaFloor, start, end)
	}

	return overlapping
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
