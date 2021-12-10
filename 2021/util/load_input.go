package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"strconv"
)

func LoadInput() []string {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Sprintf("usage: %s PUZLE_INPUT", os.Args[0]))
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	return strings.Split(strings.TrimSpace(string(buf)), "\n")
}

func LoadXInts(base int) []int {
	values := []int{}
	for _, strval := range LoadInput() {
		value, err := strconv.ParseInt(strval, base, 0)
		if err != nil {
			log.Fatal("Invalid input: ", err.Error())
		}
		values = append(values, int(value))
	}

	return values
}

func LoadBinaryInts() []int {
	return LoadXInts(2)
}

func LoadInts() []int {
	return LoadXInts(10)
}
