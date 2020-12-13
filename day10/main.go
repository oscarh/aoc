package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)


func readInput(filename string) []int {
	var adapters []int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		adapters = append(adapters, number)
	}
	return adapters
}

func checkChain(adapters []int) (int, int) {
	oneJoltDifferences := 0
	threeJoltDifferences := 0
	for index, jolts := range adapters {
		previousRating := 0
		if index > 0 {
			previousRating = adapters[index - 1]
		}
		diff := jolts - previousRating
		switch diff {
		case 1:
			oneJoltDifferences += 1
		case 3:
			threeJoltDifferences += 1
		}
	}
	return oneJoltDifferences, threeJoltDifferences + 1 // The + 1 is for the device
}

func findNumberOfPossibleChains(adapters []int) int {
	sequenceLengthToArrangements  := map[int]int{1: 1, 2: 1, 3: 2, 4: 4, 5: 7}
	previous := 0
	sequenceLength := 1
	possibleArrangements := 1
	for x := 0; x < len(adapters); x += 1 {
		current := adapters[x]
		if current == previous + 1 {
			sequenceLength += 1
		} else {
			possibleArrangements *= sequenceLengthToArrangements[sequenceLength]
			sequenceLength = 1
		}
		previous = current
	}
	possibleArrangements *= sequenceLengthToArrangements[sequenceLength]

	return possibleArrangements
}

func main() {
	fmt.Println("Day 10")
	adapters := readInput("day10/input.txt")
	sort.Ints(adapters)
	oneJoltDiffs, threeJoltDiffs := checkChain(adapters)
	log.Printf("One jolt differances (%d) multiplied with three jolt (%d) differances: %d\n",
		oneJoltDiffs, threeJoltDiffs, oneJoltDiffs * threeJoltDiffs)
	possibleChains := findNumberOfPossibleChains(adapters)
	log.Printf("Possible chains: %d\n", possibleChains)
}
