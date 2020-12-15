package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	data := parsePasswordData()
	var partOneValid int
	var partTwoValid int
	for _, datum := range data {
		occurenceValid := checkLetterOccurence(datum.Letter, datum.Password, datum.Bounds.Lower, datum.Bounds.Upper)
		positionValid := checkPosition(datum.Letter, datum.Password, datum.Bounds.Lower, datum.Bounds.Upper)
		if occurenceValid {
			partOneValid++
		}

		if positionValid {
			partTwoValid++
		}
	}

	fmt.Printf("There are %d valid passwords by the old policy. \n", partOneValid)
	fmt.Printf("There are %d valid passwords by the new policy. \n", partTwoValid)
}

type bounds struct {
	Upper int
	Lower int
}

type passwordStruct struct {
	Letter   string
	Password string
	Bounds   bounds
}

func parsePasswordData() []passwordStruct {
	file, err := os.Open("input.txt")
	check(err)

	scanner := bufio.NewScanner(file)
	var data []passwordStruct
	for scanner.Scan() {
		passwordData := scanner.Text()
		splitData := strings.Split(passwordData, " ")
		lower, higher := parseNums(splitData[0])
		letter := parseLetter(splitData[1])

		data = append(data, passwordStruct{
			Password: splitData[2],
			Letter:   letter,
			Bounds: bounds{
				Lower: lower,
				Upper: higher,
			},
		})
	}

	return data
}

func parseNums(occurenceRange string) (lower, upper int) {
	nums := strings.Split(occurenceRange, "-")
	lower, _ = strconv.Atoi(nums[0])
	upper, _ = strconv.Atoi(nums[1])
	return lower, upper
}

func parseLetter(input string) string {
	split := strings.Split(input, ":")
	return split[0]
}

func checkPosition(letter, password string, first, second int) bool {
	letters := strings.Split(password, "")
	numLetters := len(letters)
	positionOne := first - 1
	positionTwo := second - 1
	var firstInstance string
	var lastInstance string
	if positionOne >= 0 && positionOne < numLetters {
		firstInstance = letters[positionOne]
	}
	if positionTwo < numLetters {
		lastInstance = letters[positionTwo]
	}

	if firstInstance == letter && lastInstance != letter {
		return true
	} else if firstInstance != letter && lastInstance == letter {
		return true
	} else {
		return false
	}
}

func checkLetterOccurence(letter, password string, lower, upper int) bool {
	var counter int
	letters := strings.Split(password, "")

	for _, l := range letters {
		if l == letter {
			counter++
		}
	}

	if counter >= lower && counter <= upper {
		return true
	}

	return false
}
