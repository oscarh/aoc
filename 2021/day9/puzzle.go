package main

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/oscarh/aoc/util"
)

type grid struct {
	noRows int
	noCols int
	values [][]int
}

type point struct {
	x int
	y int
}

func (g *grid) points() map[point]bool {
	m := map[point]bool{}
	for x := 0; x < g.noCols; x += 1 {
		for y := 0; y < g.noRows; y += 1 {
			m[point{x: x, y: y}] = true
		}
	}

	return m
}

func adjacent(x, y int, remaining map[point]bool) []point {
	a := []point{}

	add := func(p point) {
		if remaining[p] {
			a = append(a, p)
			delete(remaining, p)
		}
	}

	add(point{x - 1, y})
	add(point{x + 1, y})
	add(point{x, y - 1})
	add(point{x, y + 1})

	return a
}

func (g *grid) basin(p point, remaining map[point]bool) int {
	size := 0
	if g.value(p.x, p.y) == 9 {
		return 0
	} else {
		size += 1
		for _, pa := range adjacent(p.x, p.y, remaining) {
			size += g.basin(pa, remaining)
		}
	}

	return size
}

func (g *grid) basins() []int {
	basins := []int{}

	remaining := g.points()
	for point := range remaining {
		// It seems the point is still there during iteration
		delete(remaining, point)
		basins = append(basins, g.basin(point, remaining))
	}

	return basins
}

func (g *grid) add(x, y, h int) {
	g.values[x][y] = h
}

func (g *grid) value(x, y int) int {
	return g.values[x][y]
}

func (g *grid) isLowPoint(x, y int) bool {
	value := g.values[x][y]
	if x > 0 {
		if value >= g.values[x-1][y] {
			return false
		}
	}
	if x < g.noCols-1 {
		if value >= g.values[x+1][y] {
			return false
		}
	}
	if y > 0 {
		if value >= g.values[x][y-1] {
			return false
		}
	}
	if y < g.noRows-1 {
		if value >= g.values[x][y+1] {
			return false
		}
	}

	return true
}

func newGrid(noRows, noCols int) grid {
	values := make([][]int, noCols)
	for x := range values {
		values[x] = make([]int, noRows)
	}
	g := grid{
		noRows: noRows,
		noCols: noCols,
		values: values,
	}

	return g
}

func loadGrid() grid {
	rows := util.LoadInput()
	g := newGrid(len(rows), len(rows[0]))
	for y, row := range rows {
		for x, value := range row {
			h, err := strconv.Atoi(string(value))
			if err != nil {
				panic(err)
			}
			g.add(x, y, h)
		}
	}
	return g
}

func part1() int {
	g := loadGrid()
	sum := 0

	for x, row := range g.values {
		for y, height := range row {
			if g.isLowPoint(x, y) {
				sum += 1 + height
			}
		}
	}
	return sum
}

func part2() int {
	g := loadGrid()
	basins := g.basins()
	sort.Ints(basins)
	l := len(basins)
	return basins[l-1] * basins[l-2] * basins[l-3]
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
