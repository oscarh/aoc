package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

type instruction struct {
	operation uint8
	argument int
}

func abs(x int ) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
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
		argument, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatal(err)
		}
		instructions = append(instructions, instruction{operation: line[0], argument: argument})
	}

	return instructions
}

func calculateNewDirection(direction int32, turnDirection int32, degrees int)  int32 {
	for steps := degrees / 90; steps > 0; steps -= 1 {
		if turnDirection == 'L' {
			switch direction {
			case 'N':
				direction = 'W'
			case 'E':
				direction = 'N'
			case 'S':
				direction = 'E'
			case 'W':
				direction = 'S'
			}
		} else {
			switch direction {
			case 'N':
				direction = 'E'
			case 'E':
				direction = 'S'
			case 'S':
				direction = 'W'
			case 'W':
				direction = 'N'
			}
		}
	}

	return direction
}

func getDistance(instructions []instruction) int {
	north := 0
	east := 0
	direction := 'E'

	for _, ins := range instructions {
		switch ins.operation {
		case 'N':
			north += ins.argument
		case 'S':
			north -= ins.argument
		case 'E':
			east += ins.argument
		case 'W':
			east -= ins.argument
		case 'L':
			direction = calculateNewDirection(direction, 'L', ins.argument)
		case 'R':
			direction = calculateNewDirection(direction, 'R', ins.argument)
		case 'F':
			switch direction {
			case 'N':
				north += ins.argument
			case 'E':
				east += ins.argument
			case 'S':
				north -= ins.argument
			case 'W':
				east -= ins.argument
			}
		}
	}

	return abs(north) + abs(east)
}

func rotateWaypoint(north int, east int, degrees int) (int, int) {
	for steps := degrees / 90; steps > 0; steps -= 1 {
		newNorth := -east
		newEast := north

		north = newNorth
		east = newEast
	}
	return north, east
}

func getDistancePart2(instructions []instruction) int {
	north := 0
	east := 0
	waypointNorth := 1
	waypointEast := 10

	for _, ins := range instructions {
		switch ins.operation {
		case 'N':
			waypointNorth += ins.argument
		case 'S':
			waypointNorth -= ins.argument
		case 'E':
			waypointEast += ins.argument
		case 'W':
			waypointEast -= ins.argument
		case 'L':
			waypointNorth, waypointEast = rotateWaypoint(waypointNorth, waypointEast, 360 - ins.argument)
		case 'R':
			waypointNorth, waypointEast = rotateWaypoint(waypointNorth, waypointEast, ins.argument)
		case 'F':
			for x := ins.argument; x > 0; x-- {
				north += waypointNorth
				east += waypointEast
			}
		}
	}

	return abs(north) + abs(east)
}

func main() {
	fmt.Println("Day 12")
	instructions := readInput("day12/input.txt")
	log.Printf("Manhattan distance after following the guessed instructions: %d\n", getDistance(instructions))
	log.Printf("Manhattan distance after following the real instructions: %d\n", getDistancePart2(instructions))
}
