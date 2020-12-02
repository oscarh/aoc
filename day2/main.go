package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Day 2")

	file, err := os.Open("day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	validPasswordsOldPolicy := 0
	validPasswordsNewPolicy := 0
	for scanner.Scan() {
		ruleAndPwd := strings.Split(scanner.Text(), ":")
		rule := ruleAndPwd[0]
		pwd := strings.TrimSpace(ruleAndPwd[1])

		minMaxChar := strings.Split(rule, " ")
		minMax := strings.Split(minMaxChar[0], "-")

		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			log.Fatal(err)
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			log.Fatal(err)
		}
		char := minMaxChar[1]

		count := strings.Count(pwd, char)
		if count >= min && count <= max {
			validPasswordsOldPolicy += 1
		}

		if pwd[min - 1] == char[0] && pwd[max - 1] != char[0] {
			validPasswordsNewPolicy += 1
		} else if pwd[max - 1] == char[0] && pwd[min - 1] != char[0] {
			validPasswordsNewPolicy += 1
		}
	}

	fmt.Printf("Valid passwords (false policy): %d\n", validPasswordsOldPolicy)
	fmt.Printf("Valid passwords (updated policy): %d\n", validPasswordsNewPolicy)
}
