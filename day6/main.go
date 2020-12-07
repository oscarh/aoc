package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readInput(filename string) (int, int) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	sumPositiveAnswers := 0
	sumAllPositiveInGroup := 0
	answers := make(map[int32]int)
	scanner := bufio.NewScanner(file)
	groupSize := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sumPositiveAnswers += len(answers)
			for _, count := range answers {
				if count == groupSize {
					sumAllPositiveInGroup += 1
				}
			}
			answers = make(map[int32]int)
			groupSize = 0
		} else {
			for _, c := range line {
				answers[c] += 1
			}
			groupSize += 1
		}
	}

	for _, count := range answers {
		if count == groupSize {
			sumAllPositiveInGroup += 1
		}
	}
	return sumPositiveAnswers + len(answers), sumAllPositiveInGroup
}

func main() {
	fmt.Println("Day 6")
	sumPositiveAnswers, sumAllPositiveInGroup := readInput("day6/input.txt")
	log.Printf("Sum of positive answers: %d\n", sumPositiveAnswers)
	log.Printf("Sum of positive answers from all in the group: %d\n", sumAllPositiveInGroup)
}
