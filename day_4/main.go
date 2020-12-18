package main

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"
)

var validators = []*regexp.Regexp{
	regexp.MustCompile(`(byr):(?:(19[2-9]\d|200[0-2])(?:\s|$))?`),
	regexp.MustCompile(`(iyr):(?:(201\d|2020)(?:\s|$))?`),
	regexp.MustCompile(`(eyr):(?:(202\d|2030)(?:\s|$))?`),
	regexp.MustCompile(`(hgt):(?:((?:1[5-8]\d|19[0-3])cm|(?:59|6\d|7[0-6])in)(?:\s|$))?`),
	regexp.MustCompile(`(hcl):(?:(#[\da-f]{6})(?:\s|$))?`),
	regexp.MustCompile(`(ecl):(?:(amb|blu|brn|gry|grn|hzl|oth)(?:\s|$))?`),
	regexp.MustCompile(`(pid):(?:(\d{9})(?:\s|$))?`),
}

func main() {
	file, _ := ioutil.ReadFile("input.txt")
	partOneValid, partTwoValid := 0, 0

	for _, passport := range strings.Split(strings.TrimSpace(string(file)), "\n\n") {
		// track validity using ints, if invalid make it 0, so nothing added to above sums
		validOne, validTwo := 1, 1

		for _, reg := range validators {
			match := reg.FindStringSubmatch(passport)

			if len(match) == 0 {
				// no match, both invalid cause they are missing out
				validOne, validTwo = 0, 0
			} else if match[2] == "" {
				validTwo = 0
			} else {
				fmt.Println(match)
			}
		}
		partOneValid, partTwoValid = partOneValid+validOne, partTwoValid+validTwo
	}

	fmt.Printf("There are %d valid passports by part 1.\nThere are %d valid passports by part 2 standards", partOneValid, partTwoValid)
}
