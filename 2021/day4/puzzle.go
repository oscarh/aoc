package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/oscarh/aoc/util"
)

type board struct {
	rows [][]int
}

func (b *board) findPosition(number int) (int, int) {
	for y, row := range b.rows {
		for x, value := range row {
			if value == number {
				return x, y
			}
		}
	}
	panic(fmt.Sprintf("dosen't contain %d\n%s", number, b))
	return 0, 0
}
func (b *board) sumNotDrawn(drawnNumbers map[int]bool) int {
	sum := 0
	for _, row := range b.rows {
		for _, value := range row {
			if !drawnNumbers[value] {
				sum += value
			}
		}
	}

	return sum
}

func (b *board) check(number int, drawnNumbers map[int]bool) (bool, int) {
	x, y := b.findPosition(number)
	bingo := true

	for x := 0; x < 5; x += 1 {
		if !drawnNumbers[b.rows[y][x]] {
			bingo = false
			break
		}
	}

	if bingo {
		return true, b.sumNotDrawn(drawnNumbers)
	}

	bingo = true
	for y := 0; y < 5; y += 1 {
		if !drawnNumbers[b.rows[y][x]] {
			bingo = false
			break
		}
	}

	if bingo {
		return true, b.sumNotDrawn(drawnNumbers)
	}

	return false, 0
}

func (b *board) String() string {
	str := strings.Builder{}
	for _, row := range b.rows {
		for _, column := range row {
			str.WriteString(fmt.Sprintf("%d ", column))
		}
		str.WriteString("\n")
	}
	str.WriteString("\n")

	return str.String()
}

func part1() int {
	_, boardLookup, numbers := setupGame()

	drawnNumbers := map[int]bool{}
	for _, num := range numbers {
		drawnNumbers[num] = true
		for _, board := range boardLookup[num] {
			if hasBingo, sumUnmarked := board.check(num, drawnNumbers); hasBingo {
				return sumUnmarked * num
			}
		}

	}

	panic("No one won")
	return 0
}

func part2() int {
	boards, boardLookup, numbers := setupGame()

	drawnNumbers := map[int]bool{}
	for _, num := range numbers {
		drawnNumbers[num] = true
		for _, board := range boardLookup[num] {
			if boards[board] {
				if hasBingo, sumUnmarked := board.check(num, drawnNumbers); hasBingo {
					if len(boards) == 1 {
						return sumUnmarked * num
					} else {
						delete(boards, board)
					}
				}
			}
		}

	}

	panic("No one won")
	return 0
}

func setupGame() (map[*board]bool, map[int][]*board, []int) {
	input := util.LoadInput()

	boards := map[*board]bool{}
	boardLookup := map[int][]*board{}

	boardNo := 0
	current := &board{}
	for _, numbers := range input[2:] {
		if numbers == "" {
			boards[current] = true
			current = &board{}
			boardNo += 1
			continue
		}

		row := []int{}
		for _, strValue := range strings.Fields(numbers) {
			num, err := strconv.Atoi(strValue)
			if err != nil {
				panic(fmt.Sprintf("Invalid input: %s, row: %s", strValue, numbers))
			}
			boardLookup[num] = append(boardLookup[num], current)
			row = append(row, num)
		}
		current.rows = append(current.rows, row)
	}

	numbers := []int{}
	for _, strValue := range strings.Split(input[0], ",") {
		num, err := strconv.Atoi(strValue)
		if err != nil {
			panic(fmt.Sprintf("Invalid input: %s", strValue))
		}
		numbers = append(numbers, num)
	}

	return boards, boardLookup, numbers
}

func main() {
	fmt.Printf("Part 1: %d\n", part1())
	fmt.Printf("Part 2: %d\n", part2())
}
