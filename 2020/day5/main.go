package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func readInput(filename string) []int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	rows := make([]int, 128)
	for r := range rows {
		rows[r] = r
	}

	columns := make([]int, 8)
	for c := range columns {
		columns[c] = c
	}

	var seatIds []int
	highestSeatId := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var row []int = rows
		var column []int = columns
		for _, c := range line[0:7] {
			if c == 'F' {
				row = row[ : len(row) / 2]
			} else if c == 'B' {
				row = row[len(row) / 2 : ]
			}
		}
		for _, c := range line[7:] {
			if c == 'L' {
				column = column[ : len(column) / 2]
			} else if c == 'R' {
				column = column[len(column) / 2 : ]
			}
		}
		if len(row) > 1 {
			log.Fatal("Failed to find row...")
		}
		if len(column) > 1 {
			log.Fatal("Failed to find column...")
		}
		seatId := row[0] * 8 + column[0]
		if seatId > highestSeatId {
			highestSeatId = seatId
		}
		seatIds = append(seatIds, seatId)
	}

	log.Printf("Highest seat ID found: %d\n", highestSeatId)
	return seatIds
}

func main() {
	fmt.Println("Day 5")
	seatIds := readInput("day5/input.txt")
	sort.Ints(seatIds)
	for i, id := range seatIds {
		if seatIds[i + 1] != id + 1 {
			log.Printf("My seat is: %d\n", id + 1)
			break
		}
	}
}
