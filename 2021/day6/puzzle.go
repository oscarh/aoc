package main

import (
	"fmt"

	"github.com/oscarh/aoc/util"
)

func bruteForce(fish []int) int {
	day := 0
	for day < 80 {
		for i, counter := range fish {
			if counter == 0 {
				fish = append(fish, 8)
				counter = 6
			} else {
				counter -= 1
			}
			fish[i] = counter
		}
		day += 1
	}
	return len(fish)
}

func part1() int {
	fish := util.LoadCommaSeparatedInts()
	return bruteForce(fish)
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}

