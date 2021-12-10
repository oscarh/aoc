package util

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"strconv"
)

func LoadInts() []int {
	if len(os.Args) < 2 {
		log.Fatal(fmt.Sprintf("usage: %s PUZLE_INPUT", os.Args[0]))
	}

	buf, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		log.Fatal(err.Error())
	}

	values := []int{}
	for _, strval := range strings.Split(strings.TrimSpace(string(buf)), "\n") {
		value, err := strconv.Atoi(strval)
		if err != nil {
			log.Fatal("Invalid input: ", err.Error())
		}
		values = append(values, value)
	}

	return values
}
