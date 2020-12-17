package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

/*

required!
   byr (Birth Year)
   iyr (Issue Year)
   eyr (Expiration Year)
   hgt (Height)
   hcl (Hair Color)
   ecl (Eye Color)
   pid (Passport ID)
   cid (Country ID)
*/

var requiredKeys = [7]string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}

func isValid(passport map[string]string) bool {
	valid := true
	for _, key := range requiredKeys {
		if passport[key] == "" {
			valid = false
			break
		}
	}
	return valid
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal("oh no")
	}

	scanner := bufio.NewScanner(file)

	var passports []map[string]string

	inc := 1
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) != 0 {
			keyValPairs := strings.Split(line, " ")
			if len(passports) < inc {
				passports = append(passports, map[string]string{})
			}
			for _, pair := range keyValPairs {
				splitPair := strings.Split(pair, ":")
				key, value := splitPair[0], splitPair[1]

				passports[inc-1][key] = value
			}
		} else {
			inc++
		}
	}

	var numValid int
	for _, passport := range passports {
		if isValid(passport) {
			numValid++
		}
	}
	fmt.Printf("There are %d valid passports\n", numValid)
}
