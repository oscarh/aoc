package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func readInput(filename string) (int, []int, []int) {
	var busses []int
	var bussesAndGaps []int

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	timestamp, err := strconv.Atoi(scanner.Text())
	if err != nil {
		log.Fatal(err)
	}

	scanner.Scan()
	for _, bus := range strings.Split(scanner.Text(), ",") {
		if bus == "x" {
			bussesAndGaps = append(bussesAndGaps, 0)
			continue
		} else {
			bus, err := strconv.Atoi(bus)
			if err != nil {
				log.Fatal(err)
			}
			busses = append(busses, bus)
			bussesAndGaps = append(bussesAndGaps, bus)
		}
	}

	return timestamp, busses, bussesAndGaps
}

func nextBus(timestamp int, busses []int) (int, int) {
	var nextBusId int
	nextBusInMinutes := math.MaxInt32

	for _, bus := range busses {
		minutesSinceBus := timestamp % bus
		minutesUntilBus := bus - minutesSinceBus
		if minutesUntilBus < nextBusInMinutes {
			nextBusId = bus
			nextBusInMinutes = minutesUntilBus
		}
	}

	return nextBusId, nextBusInMinutes
}

func maxId(busses []int) (int, int) {
	currentMax := 0
	var index int
	for i, bus := range busses {
		if bus > currentMax {
			currentMax = bus
			index = i
		}
	}

	return index, currentMax
}

func timestampWhereBussesDepartInSequence(busses []int) int {
	maxIndex, id := maxId(busses)

	timestamp := id
	OUTER:
	for {
		INNER:
		for i := 0; i < len(busses); i += 1 {
			bus := busses[i]
			if bus == 0 {
				continue INNER
			}  else if (timestamp - maxIndex + i) % bus != 0 {
				nextTimestamp := timestamp + id
				if nextTimestamp < timestamp {
					panic("Timstamp wrapped")
				} else {
					timestamp += id
				}
				continue OUTER
			}
		}
		return timestamp - maxIndex
	}
}

func main() {
	fmt.Println("Day 13")
	timestamp, bussesWithoutX, bussesWithGaps := readInput("day13/input.txt")
	id, inMinutes := nextBus(timestamp, bussesWithoutX)
	log.Printf("Bus %d comes in %d minutes (%d)\n", id, inMinutes, id * inMinutes)
	log.Printf("Busses depart in sequance at: %d\n", timestampWhereBussesDepartInSequence(bussesWithGaps))
}