package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const(
	acc = iota
	jmp = iota
	nop = iota
)

func operationName(operation int) string {
	switch operation {
	case acc:
		return "ACC"
	case jmp:
		return "JMP"
	case nop:
		return "NOP"
	}

	return "INVALID INSTRUCTION"
}

type instruction struct {
	operation int
	argument int
}

func parseOperation(operation string) int {
	if operation == "acc" {
		return acc
	} else if operation == "jmp" {
		return jmp
	} else if operation == "nop" {
		return nop
	}
	log.Fatalf("Invalid operation: %s", operation)
	return -1
}

func readInput(filename string) []instruction {
	var instructions []instruction

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		operation := parseOperation(line[0:3])
		argument, err := strconv.Atoi(line[4:])
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, instruction{operation: operation, argument: argument})
	}

	return instructions
}

func run(bootCode []instruction) (int, map[int]bool) {
	executedInstructions := make(map[int]bool)
	pc := 0
	accumulator := 0

	for pc < len(bootCode) && !executedInstructions[pc] {
		executedInstructions[pc] = true
		i := bootCode[pc]

		switch i.operation {
		case acc:
			accumulator += i.argument
			pc += 1
		case jmp:
			pc += i.argument
		case nop:
			pc += 1
		}
	}
	if pc < len(bootCode) {
		log.Println("Terminated due to infinite loop")
	} else {
		log.Println("Terminated successfully")
	}
	return accumulator, executedInstructions
}


func findterminationInstructions(bootCode []instruction, possibleEnds map[int]bool) bool {
	numPossibleEnds := len(possibleEnds)
	for index, instr := range bootCode {
		if instr.operation == jmp {
			if possibleEnds[index + instr.argument] {
				possibleEnds[index] = true
				for tmpIndex := index - 1; bootCode[tmpIndex].operation != jmp; tmpIndex-- {
					possibleEnds[tmpIndex] = true
				}
			}
		}
	}
	return len(possibleEnds) > numPossibleEnds
}

func findChangeableInstruction(bootCode []instruction, executed map[int]bool, targets map[int]bool) (int, int) {
	for instructionIndex, _ := range executed {
		operation := bootCode[instructionIndex].operation
		argument := bootCode[instructionIndex].argument
		if operation == jmp && targets[instructionIndex + 1] {
			// if a JMP can be changed to a NOP, we'll get to a target instruction
			return instructionIndex, nop
		} else if operation == nop && targets[instructionIndex + argument]  {
			// if a NOP in changed to a JMP, we'll jump to a target instruction
			return instructionIndex, jmp
		}
	}
	log.Fatal("Could not find changeable index")
	return -1, nop
}

func main() {
	fmt.Println("Day 8")
	bootCode := readInput("day8/input.txt")
	acc, executedInstructions := run(bootCode)
	log.Printf("Accumulator value before infiniate loop: %d\n", acc)
	possibleEnds := map[int]bool{len(bootCode): true}
	for findterminationInstructions(bootCode, possibleEnds) {
	}
	index, changeTo := findChangeableInstruction(bootCode, executedInstructions, possibleEnds)
	log.Printf("Changeing %s to %s at index: %d\n", operationName(bootCode[index].operation), operationName(changeTo), index)
	bootCode[index].operation = changeTo
	acc, _ = run(bootCode)
	log.Printf("Accumulator value after patch: %d\n", acc)
}
