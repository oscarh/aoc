package main

import (
	"fmt"

	"github.com/oscarh/aoc/util"
)

const bits = 12

func part1() int {
	input := util.LoadBinaryInts()

	bitcount := [bits]int{}
	for _, value := range input {
		for x := 0; x < bits; x += 1 {
			bitcount[x] += value >> x & 0x1
		}
	}

	gamma := 0
	epsilon := 0
	for x := 0; x < bits; x += 1 {
		if bitcount[x] > len(input) / 2 {
			gamma |= 1 << x
		} else {
			epsilon |= 1 << x
		}
	}


	return gamma * epsilon
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
}
