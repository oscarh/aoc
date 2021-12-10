package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) [][]int32 {
	var grid [][]int32

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		seats := make([]int32, len(line))
		for i, char := range line {
			seats[i] = char
		}
		grid = append(grid, seats)
	}

	return grid
}

func max(a int, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func min(a int, b int) int {
	if a > b {
		return b
	} else {
		return a
	}
}

func occupiedAdjacent(grid [][]int32, row int, pos int) int {
	gridWidth := len(grid[0])
	gridHeight := len(grid)

	occupied := 0
	for i := max(0, row - 1); i <= min(row + 1, gridHeight - 1); i += 1 {
		for j := max(0, pos - 1); j <= min(pos + 1, gridWidth - 1); j += 1 {
			if i != row || j != pos {
				if grid[i][j] == '#' {
					occupied += 1
				}
			}
		}
	}
	return occupied
}

func applyRule1(grid [][]int32) ([][]int32, int) {
	newGrid := make([][]int32, len(grid))
	changes := 0

	for i, row := range grid {
		newGrid[i] = make([]int32, len(row))
		for j, square := range row {
			switch square {
			case 'L':
				if occupiedAdjacent(grid, i, j) == 0 {
					changes += 1
					newGrid[i][j] = '#'
				} else {
					newGrid[i][j] = 'L'
				}
			case '#':
				if occupiedAdjacent(grid, i, j) >= 4 {
					changes += 1
					newGrid[i][j] = 'L'
				} else {
					newGrid[i][j] = '#'
				}
			case '.':
				newGrid[i][j] = '.'
			}
		}
	}
	return newGrid, changes
}

//func canSeeOccupiedSeat(grid [][]int32, row int, pos int) bool {
//	round := 1
//
//	for deltaRow := -1; deltaRow <=1; deltaRow += 1 {
//		currentRow := row + deltaRow * round
//		if currentRow > 0 && currentRow < len(grid) {
//			for deltaPos := -1; deltaPos <=1; deltaPos += 1 {
//				if deltaRow != 0 || deltaPos != 0 {
//					currentPos := pos + deltaPos*round
//					if currentPos > 0 && currentPos < len(grid[0]) {
//
//					}
//				}
//			}
//		}
//		round += 1
//	}
//}

const(
	occupied = iota
	free = iota
	unknown = iota
)

func getState(grid [][]int32, row int, pos int) int {
	if row < 0 || row >= len(grid) || pos < 0 || pos >= len(grid[0]){
		// Outside of grid, free
		return free
	} else {
		switch grid[row][pos] {
		case '.':
			return unknown
		case '#':
			return occupied
		case 'L':
			return free
		}
	}
	return unknown
}

func canSeeXOccupiedSeat(grid [][]int32, row int, pos int, expectedOccupiedSeats int) bool {
	states := []int{unknown, unknown, unknown, unknown, unknown, unknown, unknown, unknown}
	deltaRow := []int{-1, -1,  0,  1,  1,  1,  0, -1}
	deltaPos := []int{0 ,  1,  1,  1,  0, -1, -1, -1}

	occupiedSeats := 0

	distance := 1
	for occupiedSeats < expectedOccupiedSeats {
		directionsTested := 0
		for direction, state := range states {
			if state == unknown {
				directionsTested += 1

				rowToInspect := row + deltaRow[direction] * distance
				posToInspect := pos + deltaPos[direction] * distance
				newState := getState(grid, rowToInspect, posToInspect)
				states[direction] = newState
				if newState == occupied {
					occupiedSeats += 1
					if occupiedSeats == expectedOccupiedSeats {
						break
					}
				}
			}
		}
		if directionsTested == 0 {
			break
		}
		distance += 1
	}

	return occupiedSeats == expectedOccupiedSeats

}

func applyRule2(grid [][]int32) ([][]int32, int) {
	newGrid := make([][]int32, len(grid))
	changes := 0

	for i, row := range grid {
		newGrid[i] = make([]int32, len(row))
		for j, square := range row {
			switch square {
			case 'L':
				if !canSeeXOccupiedSeat(grid, i, j, 1) {
					newGrid[i][j] = '#'
					changes += 1
				} else {
					newGrid[i][j] = 'L'
				}
			case '#':
				if canSeeXOccupiedSeat(grid, i, j, 5) {
					newGrid[i][j] = 'L'
					changes += 1
				} else {
					newGrid[i][j] = '#'
				}
			case '.':
				newGrid[i][j] = '.'
			}
		}
	}

	return newGrid, changes
}

func countOccupied(grid [][]int32) int {
	occupied := 0
	for _, row := range grid {
		for _, pos := range row {
			if pos == '#' {
				occupied += 1
			}
		}
	}
	return occupied
}

func print(grid [][]int32) {
	for _, row := range grid {
		for _, pos := range row {
			fmt.Printf(" %c ", pos)
		}
		fmt.Printf("\n")
	}

	fmt.Printf("\n\n")
}

func _runRule1UntilStable(grid [][]int32, rounds int) ([][]int32, int) {
	newGrid, changes := applyRule1(grid)
	rounds += 1
	if changes != 0 {
		return _runRule1UntilStable(newGrid, rounds)
	} else {
		return newGrid, rounds
	}
}

func runRule1UntilStable(grid [][]int32) ([][]int32, int) {
	return _runRule1UntilStable(grid, 0)
}

func _runRule2UntilStable(grid [][]int32, rounds int) ([][]int32, int) {
	newGrid, changes := applyRule2(grid)
	rounds += 1
	if changes != 0 {
		return _runRule2UntilStable(newGrid, rounds)
	} else {
		return newGrid, rounds
	}
}

func runRule2UntilStable(grid [][]int32) ([][]int32, int) {
	return _runRule2UntilStable(grid, 0)
}

func main() {
	fmt.Println("Day 11")
	grid := readInput("day11/input.txt")
	stableGrid1, runs := runRule1UntilStable(grid)

	log.Printf("Rule 1: Occupied seats when stable (after %d runs): %d\n", runs, countOccupied(stableGrid1))

	stableGrid2, runs := runRule2UntilStable(grid)
	log.Printf("Rule 2: Occupied seats when stable (after %d runs): %d\n", runs, countOccupied(stableGrid2))
}
