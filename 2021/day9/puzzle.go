package main

import (
	"fmt"
	"strconv"

	"github.com/oscarh/aoc/util"
)

type grid struct {
	noRows int
	noCols int
	values [][]int
}

func (g *grid) add(x, y, h int) {
	g.values[x][y] = h
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
	if y < g.noCols-1 {
		if value >= g.values[x][y+1] {
			return false
		}
	}

	return true
}

func newGrid(noRows, noCols int) grid {
	values := make([][]int, noRows)
	for x := range values {
		values[x] = make([]int, noCols)
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
	for x, row := range rows {
		for y, value := range row {
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

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	//fmt.Printf("Part 2: %d\n", part2())
}
