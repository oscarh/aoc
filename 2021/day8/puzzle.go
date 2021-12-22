package main

import (
	"fmt"
	"strings"

	"github.com/oscarh/aoc/util"
)

type entry struct {
	patterns []string
	outputs  []string
}

func loadPatternAndOutputs() []entry {
	var entries []entry
	for _, line := range util.LoadInput() {
		split := strings.Split(line, " | ")
		entries = append(entries, entry{
			patterns: strings.Split(split[0], " "),
			outputs:  strings.Split(split[1], " "),
		})
	}
	return entries
}

func part1() int {
	count := 0
	for _, e := range loadPatternAndOutputs() {
		unique := map[int]bool{2: true, 3: true, 4: true, 7: true}
		for _, d := range e.outputs {
			if unique[len(d)] {
				count += 1
			}
		}
	}

	return count
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	//fmt.Printf("Part 2: %d\n", part2())
}
