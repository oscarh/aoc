package main

import (
	"fmt"
	"sort"

	"github.com/oscarh/aoc/util"
)

func mean(nums []int) int {
	sum := 0
	for _, n := range nums {
		sum +=n
	}
	return sum / len(nums)
}

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

func gauss(n int) int {
	return n * (n + 1) / 2
}

func move(positions []int, target int, costFunc func(int) int) int {
	cost := 0
	for _, pos := range positions {
		cost += costFunc(target - pos)
	}
	return cost
}

func part1() int {
	positions := util.LoadCommaSeparatedInts()
	target := median(positions)
	return move(positions, target, abs)
}

func part2() int {
	positions := util.LoadCommaSeparatedInts()
	target := mean(positions)
	targetCost := move(positions, target, func (d int) int { return gauss(abs(d))})
	t := target - 1
	for {
		tCost := move(positions, t, func (d int) int { return gauss(abs(d))})
		if tCost < targetCost {
			target = t
			targetCost = tCost
			t -= 1
		} else {
			break
		}

	}
	t = target + 1
	for {
		tCost := move(positions, t, func (d int) int { return gauss(abs(d))})
		if tCost < targetCost {
			target = t
			targetCost = tCost
			t += 1
		} else {
			break
		}
	}

	return targetCost
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}

