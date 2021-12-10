package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const(
	mask = iota
	mem = iota
)

type instruction struct {
	operation int
	value int
}

func readInput(filename string) []instruction {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var instructions []instruction
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value := strings.TrimSpace(strings.Split(line, "=")[1])
		if line[0:5] == "mask" {

			instructions = append(instructions, instruction{mask, value})
		}
	}
}

func main() {
	fmt.Println("Day 14")
	instructions := readInput("day14/input.txt")
	runInitializationCode(instructions)
}
