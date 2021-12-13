package main

import (
	"fmt"
	"sort"

	"github.com/oscarh/aoc/util"
)

func median(nums []int) int {
	tmp := nums
	sort.Ints(tmp)
	middle := len(tmp) / 2
	if  len(tmp) % 2 == 0 {
		return tmp[middle]
	} else {
		return (tmp[middle - 1] + tmp[middle]) / 2
	}
}

func abs(v int) int {
	if v < 0 {
		return -1 * v
	} else {
		return v
	}
}

func move(positions []int, target int) int {
	cost := 0
	for _, pos := range positions {
		cost += abs(target - pos)
	}
	return cost
}

func part1() int {
	positions := util.LoadCommaSeparatedInts()
	target := median(positions)
	return move(positions, target)
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

