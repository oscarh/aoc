package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func part1(entries map[int]bool) {
	for e1 := range entries {
		for e2 := range entries {
			if e1 + e2 == 2020 {
				fmt.Printf("%d + %d == 2020\n", e1, e2)
				fmt.Printf("%d * %d == %d\n", e1, e2, e1 * e2)
				return
			}
		}
	}
}

func part2(entries map[int]bool) {
	for e1 := range entries {
		for e2 := range entries {
			for e3 := range entries {
				if e1 + e2 + e3 == 2020 {
					fmt.Printf("%d + %d + %d == 2020\n", e1, e2, e3)
					fmt.Printf("%d * %d * %d == %d\n", e1, e2, e3, e1 * e2 * e3)
					return
				}
			}
		}
	}
}

func main() {
	file, err := os.Open("day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	entries := make(map[int]bool)
	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		entries[i] = true
	}

	part1(entries)
	part2(entries)
}
