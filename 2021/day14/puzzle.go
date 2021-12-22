package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/oscarh/aoc/util"
)

type pair struct {
	first  rune
	second rune
}

type element struct {
	r    rune
	next *element
}

type template struct {
	head *element
	tail *element
}

type cacheEntry struct {
	polymer string
	steps   int
}

type charCount map[rune]int
type countCache map[cacheEntry]charCount

type insertionRules map[pair]rune

func (t *template) add(r rune) {
	e := &element{r: r, next: nil}

	if t.head == nil && t.tail == nil {
		t.head = e
		t.tail = e
	} else {
		t.tail.next = e
		t.tail = e
	}
}

func (t template) String() string {
	b := strings.Builder{}
	c := t.head
	for c != nil {
		b.WriteRune(c.r)
		c = c.next
	}
	return b.String()
}

func (t *template) expand(pi insertionRules) {
	f := t.head
	s := f.next

	for s != nil {
		r := pi[pair{f.r, s.r}]
		f.next = &element{r: r, next: s}

		f = s
		s = s.next
	}
}

func (t *template) toSlice() []rune {
	slice := []rune{}
	c := t.head
	for c != nil {
		slice = append(slice, c.r)
		c = c.next
	}
	return slice
}

func merge(tot, add charCount) {
	for r, c := range add {
		tot[r] += c
	}
}

func expandAndCountCached(part []rune, pi insertionRules, steps int, cache countCache) charCount {
	if steps == 0 {
		return charCount{}
	}

	cacheKey := cacheEntry{string(part), steps}
	if counts, ok := cache[cacheKey]; ok {
		return counts
	}

	counts := charCount{}

	for x := 0; x < len(part)-1; x += 1 {
		first := part[x]
		second := part[x+1]
		between := pi[pair{first, second}]
		counts[between] += 1

		merge(counts, expandAndCountCached([]rune{first, between, second}, pi, steps-1, cache))
	}

	cache[cacheKey] = counts

	return counts
}

func (t *template) expandAndCount(pi insertionRules, steps int) charCount {
	counts := t.countElements()
	cache := countCache{}
	merge(counts, expandAndCountCached(t.toSlice(), pi, steps, cache))
	fmt.Println(counts)
	return counts
}

func (t *template) countElements() map[rune]int {
	counts := map[rune]int{}
	c := t.head
	for c != nil {
		counts[c.r] += 1
		c = c.next
	}
	fmt.Println(counts)
	return counts
}

func parseTemplate(line string) template {
	t := template{}
	for _, r := range line {
		t.add(r)
	}
	return t
}

func parsePairInsertions(lines []string) insertionRules {
	rules := insertionRules{}
	for _, insertion := range lines {
		ins := strings.Split(insertion, " -> ")

		rules[pair{rune(ins[0][0]), rune(ins[0][1])}] = rune(ins[1][0])
	}
	return rules
}

func loadInput() (template, insertionRules) {
	input := util.LoadInput()
	t := parseTemplate(input[0])
	pi := parsePairInsertions(input[2:])

	return t, pi
}

func part2() int {
	t, pi := loadInput()

	counts := []int{}
	for _, count := range t.expandAndCount(pi, 40) {
		counts = append(counts, count)
	}

	sort.Ints(counts)

	return counts[len(counts)-1] - counts[0]
}

func part1() int {
	t, pi := loadInput()

	for x := 0; x < 10; x += 1 {
		t.expand(pi)
	}

	counts := []int{}
	for _, count := range t.countElements() {
		counts = append(counts, count)
	}

	sort.Ints(counts)

	return counts[len(counts)-1] - counts[0]
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
