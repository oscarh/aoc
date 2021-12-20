package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/oscarh/aoc/util"
)

type stack []rune

func (s *stack) pop() rune {
	n := len(*s) - 1
	c := (*s)[n]
	*s = (*s)[:n]
	return c
}

func (s *stack) push(c rune) {
	*s = append(*s, c)
}

func (s stack) String() string {
	b := strings.Builder{}
	for _, c := range s {
		b.WriteRune(c)
	}
	return b.String()
}

func opens(c rune) bool {
	switch c {
	case '{':
		return true
	case '(':
		return true
	case '<':
		return true
	case '[':
		return true
	}

	return false
}

func matches(o, c rune) bool {
	switch o {
	case '{':
		return c == '}'
	case '(':
		return c == ')'
	case '<':
		return c == '>'
	case '[':
		return c == ']'
	}

	return false
}

func pointsCorrupt(c rune) int {
	switch c {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}

	panic("Invalid rune")
}

func pointsMissing(s stack) int {
	p := 0
	for x := len(s) - 1; x >= 0; x -= 1 {
		p *= 5

		switch s[x] {
		case '(':
			p += 1
		case '[':
			p += 2
		case '{':
			p += 3
		case '<':
			p += 4
		}
	}

	return p
}

func part1() int {
	points := 0

LINES:
	for _, line := range util.LoadInput() {
		stack := stack{}
		for _, c := range line {
			if opens(c) {
				stack.push(c)
			} else {
				o := stack.pop()
				if !matches(o, c) {
					points += pointsCorrupt(c)
					continue LINES
				}
			}
		}
	}
	return points
}

func part2() int {
	points := []int{}

LINES:
	for _, line := range util.LoadInput() {
		stack := stack{}
		for _, c := range line {
			if opens(c) {
				stack.push(c)
			} else {
				o := stack.pop()
				if !matches(o, c) {
					continue LINES
				}
			}
		}

		points = append(points, pointsMissing(stack))
	}

	sort.Ints(points)

	return points[len(points)/2]
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
