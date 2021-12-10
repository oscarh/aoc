package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oscarh/aoc/util"
)

func command(cmd string) (string, int) {
	parts := strings.Split(cmd, " ")
	val, _ := strconv.Atoi(parts[1])
	return parts[0], val
}

func part1() int {
	h := 0
	v := 0

	for _, step := range util.LoadInput() {
		op, val := command(step)
		switch op {
		case "forward":
			h += val
		case "down":
			v += val
		case "up":
			v -= val
		}
	}

	return h * v
}

func part2() int {
	h := 0
	v := 0
	aim := 0

	for _, step := range util.LoadInput() {
		op, val := command(step)
		switch op {
		case "forward":
			h += val
			v += aim * val
		case "down":
			aim += val
		case "up":
			aim -= val
		}
	}

	return h * v
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
