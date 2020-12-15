package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const NUM_TO_GET = 2020

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var numbers []int
	for scanner.Scan() {
		text := scanner.Text()
		num, _ := strconv.Atoi(text)
		numbers = append(numbers, num)
	}
outer:
	for _, num := range numbers {
		for _, innerNum := range numbers {
			if innerNum != num {
				total := innerNum + num

				if total == NUM_TO_GET {
					multiplied := innerNum * num
					fmt.Printf("the answer is %d \n", multiplied)
					ioutil.WriteFile("answer.txt", []byte(strconv.Itoa(multiplied)), 0644)
					break outer
				}
			}
		}
	}
}
