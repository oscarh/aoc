package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oscarh/aoc/util"
)

type grid struct {
	maxx   int
	maxy   int
	values [][]bool
}

type coord struct {
	x int
	y int
}

type foldSpec struct {
	axis string
	pos  int
}

func (g *grid) String() string {
	b := strings.Builder{}
	for y := 0; y <= g.maxy; y += 1 {
		for x := 0; x <= g.maxx; x += 1 {
			if g.values[x][y] {
				b.WriteString("#")
			} else {
				b.WriteString(".")
			}
		}
		b.WriteString("\n")
	}
	return b.String()
}

func (g *grid) fold(f foldSpec) {
	if f.axis == "x" {
		// Fold left
		for x := 1; x <= (g.maxx - f.pos); x += 1 {
			for y := 0; y <= g.maxy; y += 1 {
				if g.values[f.pos+x][y] {
					g.values[f.pos-x][y] = true
				}
			}
		}
		g.maxx = f.pos - 1
	} else {
		// Fold up
		for y := 1; y <= (g.maxy - f.pos); y += 1 {
			for x := 0; x <= g.maxx; x += 1 {
				if g.values[x][f.pos+y] {
					g.values[x][f.pos-y] = true
				}
			}
		}
		g.maxy = f.pos - 1
	}
}

func (g *grid) countDots() int {
	count := 0
	for y := 0; y <= g.maxy; y += 1 {
		for x := 0; x <= g.maxx; x += 1 {
			if g.values[x][y] {
				count += 1
			}
		}
	}
	return count
}

func parseCoord(line string) coord {
	xy := strings.Split(line, ",")
	x, err := strconv.Atoi(xy[0])
	if err != nil {
		panic(fmt.Sprintf("Could not convert %s to int: %s", xy[0], err.Error()))
	}
	y, err := strconv.Atoi(xy[1])
	if err != nil {
		panic(fmt.Sprintf("Could not convert %s to int: %s", xy[1], err.Error()))
	}
	return coord{x, y}
}

func parseFold(line string) foldSpec {
	split := strings.Split(line, " ")
	axispos := strings.Split(split[2], "=")
	pos, err := strconv.Atoi(axispos[1])
	if err != nil {
		panic(err.Error())
	}

	return foldSpec{axispos[0], pos}
}

func loadInstructions() ([]coord, int, int, []foldSpec) {
	maxx := 0
	maxy := 0

	coords := []coord{}
	folds := []foldSpec{}
	for _, line := range util.LoadInput() {
		if strings.HasPrefix(line, "fold") {
			folds = append(folds, parseFold(line))
		} else if line != "" {
			c := parseCoord(line)
			coords = append(coords, c)

			if c.x > maxx {
				maxx = c.x
			}
			if c.y > maxy {
				maxy = c.y
			}
		}
	}

	return coords, maxx, maxy, folds
}

func newGrid(maxx, maxy int) *grid {
	g := grid{
		maxx: maxx,
		maxy: maxy,
	}

	g.values = make([][]bool, maxx+1)
	for x := range g.values {
		g.values[x] = make([]bool, maxy+1)
	}
	return &g
}

func loadInput() (*grid, []foldSpec) {
	coords, maxx, maxy, folds := loadInstructions()
	grid := newGrid(maxx, maxy)
	for _, coord := range coords {
		grid.values[coord.x][coord.y] = true
	}
	return grid, folds
}

func part1() int {
	g, folds := loadInput()
	g.fold(folds[0])

	return g.countDots()
}

func part2() string {
	g, folds := loadInput()
	for _, f := range folds {
		g.fold(f)
	}

	return fmt.Sprintln(g)
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2:\n%s\n", part2())
}
