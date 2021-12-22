package main

import (
	"fmt"

	"github.com/oscarh/aoc/util"
)

func countIncreased(values []int) int {
	noIncreased := 0
	previousValue := values[0]
	for _, value := range values[1:] {
		if value > previousValue {
			noIncreased += 1
		}
		previousValue = value
	}
	return noIncreased
}

func part1() int {
	return countIncreased(util.LoadInts())
}

func part2() int {
	rawValues := util.LoadInts()
	groupedValues := []int{}

	for pos, value := range rawValues {
		groupedValues = append(groupedValues, value)

		if pos > 0 {
			groupedValues[pos-1] += value
		}

		if pos > 1 {
			groupedValues[pos-2] += value
		}
	}

	return countIncreased(groupedValues)
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
