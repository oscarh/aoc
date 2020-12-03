package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func checkSlope(slopeX int, slopeY int, rows []string) int {
	x := 0
	y := 0

	rowLength := len(rows[0])
	trees := 0

	for y < len(rows) - 1 {
		x += slopeX
		if x >= rowLength {
			x -= rowLength
		}
		y += slopeY

		if rows[y][x] == '#' {
			trees += 1
		}
	}
	return trees
}

func main() {
	fmt.Println("Day 3")

	file, err := os.Open("day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var rows []string
	for scanner.Scan() {
		rows = append(rows, scanner.Text())
	}

	trees := checkSlope(3, 1, rows)
	fmt.Printf("Trees with slope 3, 1: %d\n", trees)

	trees = checkSlope(1, 1, rows) *
		checkSlope(3, 1, rows) *
		checkSlope(5, 1, rows) *
		checkSlope(7, 1, rows) *
		checkSlope(1, 2, rows)

	fmt.Printf("Trees from multiple slopes: %d\n", trees)

}