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

func addLine(floor map[point]int, start, end point) int {
	newOverlapping := 0
	if start.x == end.x {
		if start.y > end.y {
			tmp := end.y
			end.y = start.y
			start.y = tmp
		}
		for y := start.y; y <= end.y; y+= 1 {
			floor[point{start.x,y}] += 1
			if floor[point{start.x,y}] == 2 {
				newOverlapping += 1
			}
		}
	} else if start.y == end.y {
		if start.x > end.x {
			tmp := end.x
			end.x = start.x
			start.x = tmp
		}
		for x := start.x; x <= end.x; x+= 1 {
			floor[point{x,start.y}] += 1
			if floor[point{x,start.y}] == 2 {
				newOverlapping += 1
			}
		}
	}
	return newOverlapping 
}

func part1() int {
	input := util.LoadInput()
	seaFloor := map[point]int{}
	overlapping := 0
	for _, line := range input {
		fmt.Println(line)
		start, end := parseLine(line)
		overlapping += addLine(seaFloor, start, end)
	}

	return overlapping
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	//fmt.Printf("Part 2: %d\n", part2())
}
