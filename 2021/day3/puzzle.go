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
		if bitcount[x] > len(input)/2 {
			gamma |= 1 << x
		} else {
			epsilon |= 1 << x
		}
	}

	return gamma * epsilon
}

func part2() int {
	oxygen := util.LoadBinaryInts()
	co2 := util.LoadBinaryInts()

	for x := bits - 1; x > -1; x -= 1 {
		bitcount := 0
		for _, value := range oxygen {
			bitcount += value >> x & 0x1
		}

		criteria := 0
		if bitcount*2 >= len(oxygen) {
			criteria = 1
		}

		if len(oxygen) > 1 {
			n := 0
			for _, value := range oxygen {
				if value>>x&0x1 == criteria {
					oxygen[n] = value
					n += 1
				}

			}
			oxygen = oxygen[:n]
		}

		bitcount = 0
		for _, value := range co2 {
			bitcount += value >> x & 0x1
		}

		criteria = 0
		if bitcount*2 < len(co2) {
			criteria = 1
		}

		if len(co2) > 1 {
			n := 0
			for _, value := range co2 {
				if value>>x&0x1 == criteria {
					co2[n] = value
					n += 1
				}

			}
			co2 = co2[:n]
		}
	}

	return oxygen[0] * co2[0]
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
