package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type passport struct {
	byr *string
	iyr *string
	eyr *string
	hgt *string
	hcl *string
	ecl *string
	pid *string
	cid *string
}

func readInput(filename string) []passport {
	var passports []passport

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	p := passport{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			passports = append(passports, p)
			p = passport{}
		} else {
			for _, token := range strings.Split(line, " ") {
				field := strings.Split(token, ":")
				value := field[1]
				switch field[0] {
				case "byr":
					p.byr = &value
				case "iyr":
					p.iyr = &value
				case "eyr":
					p.eyr = &value
				case "hgt":
					p.hgt = &value
				case "hcl":
					p.hcl = &value
				case "ecl":
					p.ecl = &value
				case "pid":
					p.pid = &value
				case "cid":
					p.cid = &value
				}
			}
		}
	}

	return passports
}

func hasRequiredFields(p passport) bool {
	return p.byr  != nil &&
			p.iyr  != nil &&
			p.eyr  != nil &&
			p.hgt  != nil &&
			p.hcl  != nil &&
			p.ecl  != nil &&
			p.pid  != nil
}

func validByr(byr string) bool {
	if len(byr) == 4 {
		byr, err := strconv.Atoi(byr)
		if err != nil {
			log.Fatal(err)
		}
		return byr >= 1920 && byr <= 2002

	} else {
		return false
	}
}

func validIyr(iyr string) bool {
	if len(iyr) == 4 {
		iyr, err := strconv.Atoi(iyr)
		if err != nil {
			log.Fatal(err)
		}
		return iyr >= 2010 && iyr <= 2020

	} else {
		return false
	}
}

func validEyr(eyr string) bool {
	if len(eyr) == 4 {
		eyr, err := strconv.Atoi(eyr)
		if err != nil {
			log.Fatal(err)
		}
		return eyr >= 2020 && eyr <= 2030

	} else {
		return false
	}
}

func validHgt(hgt string) bool {
	if len(hgt) > 2 {
		unit := hgt[len(hgt) - 2:len(hgt)]
		height, err := strconv.Atoi(hgt[0:len(hgt) - 2])
		if err != nil {
			log.Fatal(err)
		}
		if unit == "cm" {
			return height >= 150 && height <= 193
		} else if unit == "in" {
			return height >= 59 && height <= 76
		} else {
			return false
		}
	} else {
		return false
	}
}

func validHcl(hcl string) bool {
	valid, err := regexp.MatchString("^#[0-9a-f]{6}$", hcl)
	if err != nil {
		log.Fatal(err)
	}
	return valid
}

func validEcl(ecl string) bool {
	return ecl == "amb" ||
			ecl == "blu" ||
			ecl == "brn" ||
			ecl == "gry" ||
			ecl == "grn" ||
			ecl == "hzl" ||
			ecl == "oth"
}

func validPid(pid string) bool {
	valid, err := regexp.MatchString("^[0-9]{9}$", pid)
	if err != nil {
		log.Fatal(err)
	}
	return valid
}

func isValidData(p passport) bool {
	return validByr(*p.byr) &&
			validIyr(*p.iyr) &&
			validEyr(*p.eyr) &&
			validHgt(*p.hgt) &&
			validHcl(*p.hcl) &&
			validEcl(*p.ecl) &&
			validPid(*p.pid)

}

func main() {
	fmt.Println("Day 4")
	passports := readInput("day4/input.txt")
	passportsWithRequiredFields := 0
	passportsWithValidData := 0
	for _, p := range passports {
		if hasRequiredFields(p) {
			passportsWithRequiredFields += 1
			if isValidData(p) {
				passportsWithValidData += 1
			}
		}
	}
	fmt.Printf("Passports with all required fields: %d\n", passportsWithRequiredFields)
	fmt.Printf("Passports with validData: %d\n", passportsWithValidData)
}