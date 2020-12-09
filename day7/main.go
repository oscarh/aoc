package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func readInput(filename string) map[string]map[string]int {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	rules := make(map[string]map[string]int)
	bagsInContainer := make(map[string]int)
	for scanner.Scan() {
		line := scanner.Text()
		const Delimiter = " bags contain "
		i := strings.Index(line, Delimiter)
		container := line[:i]
		contents := strings.Split(line[i + len(Delimiter):], ",")
		if len(contents) > 1 || contents[0] != "no other bags." {
			for _, bag := range contents {
				bag = strings.Trim(bag, " .")
				quantity, err := strconv.Atoi(bag[0:1])
				if err != nil {
					log.Fatal(err)
				}
				if quantity == 1 {
					bag = bag[0:len(bag) - 4]
				} else {
					bag = bag[0:len(bag) - 5]
				}
				bagsInContainer[bag[2:]] = quantity
			}
		}
		rules[container] = bagsInContainer
		bagsInContainer = make(map[string]int)
	}

	return rules
}

func canContain(bag string, targetBag string, rules map[string]map[string]int) bool {
	for bag := range rules[bag] {
		if bag == targetBag {
			return true
		} else {
			if canContain(bag, targetBag, rules) {
				return true
			} else {
				continue
			}
		}
	}
	return false
}

func mustContainHowMany(bag string, rules map[string]map[string]int) int {
	cointainedBags := 0

	for bagName, count := range rules[bag] {
		cointainedBags  += count
		cointainedBags  += count * mustContainHowMany(bagName, rules)
	}

	return cointainedBags
}

func main() {
	fmt.Println("Day 7")
	rules := readInput("day7/input.txt")
	bagsWhichCanContainShinyGold := 0

	for container, _ := range rules {
		if canContain(container, "shiny gold", rules) {
			bagsWhichCanContainShinyGold += 1
		}
	}
	fmt.Printf("%d bags can contain shiny gold bags \n", bagsWhichCanContainShinyGold)
	fmt.Printf("A \"shiny gold\" bag must contain %d bags\n", mustContainHowMany("shiny gold", rules))

}