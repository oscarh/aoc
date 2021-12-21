package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oscarh/aoc/util"
)

const width = 10
const height = 10

type grid struct {
	octopus [10][10]int
}

func loadGrid() grid {
	g := grid{}
	for y, row := range util.LoadInput() {
		for x, r := range row {
			v, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err.Error())
			}

			g.octopus[x][y] = v
		}
	}
	return g
}

func (g grid) String() string {
	b := strings.Builder{}
	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			b.WriteString(strconv.Itoa(g.octopus[x][y]))
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (g *grid) increaseOctopus(x, y int) {
	g.octopus[x][y] += 1
	if g.octopus[x][y] == 10 {
		// flash
		g.increaseAjacent(x, y)
	}
}

func (g *grid) increaseAjacent(x, y int) {
	if x > 0 {
		if y > 0 {
			g.increaseOctopus(x-1, y-1)
		}
		g.increaseOctopus(x-1, y)
		if y < height-1 {
			g.increaseOctopus(x-1, y+1)
		}
	}

	if y > 0 {
		g.increaseOctopus(x, y-1)
	}
	if y < height-1 {
		g.increaseOctopus(x, y+1)
	}

	if x < width-1 {
		if y > 0 {
			g.increaseOctopus(x+1, y-1)
		}
		g.increaseOctopus(x+1, y)
		if y < height-1 {
			g.increaseOctopus(x+1, y+1)
		}
	}
}

func (g *grid) step() int {
	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			g.increaseOctopus(x, y)
		}
	}

	flashes := 0
	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			if g.octopus[x][y] > 9 {
				g.octopus[x][y] = 0
				flashes += 1
			}
		}
	}
	return flashes
}

func part1() int {
	g := loadGrid()
	flashes := 0
	for x := 0; x < 100; x += 1 {
		flashes += g.step()
	}
	return flashes
}

func part2() int {
	g := loadGrid()
	x := 0
	for {
		x += 1
		if g.step() == 100 {
			return x
		}
	}
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
