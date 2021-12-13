package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"strconv"
)

func convertToInt(strValues []string, base int) []int {
	values := []int{}
	for _, strval := range strValues {
		value, err := strconv.ParseInt(strval, base, 0)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Invalid input: %s", err.Error())
			os.Exit(1)
		}
		values = append(values, int(value))
	}
	return values
}

func LoadCommaSeparatedInts() []int {
	input := LoadInput()
	return convertToInt(strings.Split(input[0], ","), 10)
}

func LoadInput() []string {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s PUZLE_INPUT\n", os.Args[0])
		os.Exit(1)
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}

	return strings.Split(strings.TrimSpace(string(buf)), "\n")
}

func LoadXInts(base int) []int {
	return convertToInt(LoadInput(), base)
}

func LoadBinaryInts() []int {
	return LoadXInts(2)
}

func LoadInts() []int {
	return LoadXInts(10)
}
