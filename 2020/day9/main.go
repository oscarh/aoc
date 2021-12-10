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
	var data []int

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
		data = append(data, number)
	}
	return data
}

func isSumOfNumbers(number int, summands []int) bool {
	for _, a := range summands {
		for _, b := range summands {
			if a == b {
				continue
			}
			if a + b == number {
				return true
			}
		}
	}
	return false
}

func findEncodingError(data []int) int {
	for i := 25; i < len(data); i += 1 {
		number := data[i]
		if !isSumOfNumbers(number, data[i - 25:i]) {
			return number
		}
	}
	log.Fatal("Could not find invalid number")
	return -1
}

func findSequenceThatSumsTo(targetSum int, data []int) []int {
	for i := 0; i < len(data); i += 1 {
		sum := 0
		for j := i; j < len(data) && sum < targetSum; j += 1 {
			sum += data[j]
			if sum == targetSum {
				return data[i:j]
			}
		}
	}
	log.Fatal("No sequence found")
	return []int{}
}

func main() {
	fmt.Println("Day 9")
	data := readInput("day9/input.txt")
	encodingError := findEncodingError(data)
	log.Printf("First invalid number: %d\n", encodingError)
	sequence := findSequenceThatSumsTo(encodingError, data)
	sort.Ints(sequence)
	log.Printf("Encryption weakness: %d\n", sequence[0] + sequence[len(sequence) - 1])
}
