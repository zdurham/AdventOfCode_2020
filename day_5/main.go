package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"sort"
	"strings"
)

func main() {
	file, _ := ioutil.ReadFile("input.txt")

	lines := strings.Split(string(file), "\n")

	var biggestSeatID float64
	var seatIDs []float64
	for _, line := range lines {
		seatID := GetSeatID(line)
		seatIDs = append(seatIDs, seatID)
		if seatID > biggestSeatID {
			biggestSeatID = seatID
		}
	}

	mySeatID := FindMissingSeatID(seatIDs)
	fmt.Printf("The biggest seatID is %v.\nMy seat number is %v\n", biggestSeatID, mySeatID)
}

func FindMissingSeatID(seatIDs []float64) (mySeatID float64) {
	sort.Float64s(seatIDs)

	for i := range seatIDs {
		// make sure we aren't at the front or back of the list
		if i != 0 || (i > 0 && i != (len(seatIDs)-1)) {
			curr := seatIDs[i]
			last := seatIDs[i-1]
			if curr-last == 2 {
				mySeatID = curr - 1
				break
			}
		}
	}
	return
}

func GetSeatID(letters string) float64 {
	firstSeven := strings.Join(strings.Split(letters, "")[0:7], "")
	lastThree := strings.Join(strings.Split(letters, "")[7:10], "")
	row := FindPosition(firstSeven, 127.0, 0.0)
	seat := FindPosition(lastThree, 7.0, 0)

	// seat ID is row times 8, plus the seat position
	return (row * 8) + seat
}

func FindPosition(letters string, upper, lower float64) float64 {
	var position float64
	for i, letterByte := range strings.Split(letters, "") {
		letter := string(letterByte)
		midPoint := findMidpoint(upper, lower)
		// lower half
		if letter == "F" || letter == "L" {
			upper = midPoint
		} else if letter == "B" || letter == "R" {
			// upper half
			lower = midPoint
		}

		if i == len(letters)-1 {
			if letter == "F" || letter == "L" {
				position = midPoint
			} else {
				position = midPoint + 1
			}
		}
	}
	return position
}

func findMidpoint(upper, lower float64) float64 {
	return math.Floor((upper + lower) / 2)
}
