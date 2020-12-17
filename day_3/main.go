package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

type Steps struct {
	Right    int
	Down     int
	X        int
	NumTrees int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	steps := []Steps{
		{
			X:        0,
			Right:    1,
			Down:     1,
			NumTrees: 0,
		},
		{
			X:        0,
			Right:    3,
			Down:     1,
			NumTrees: 0,
		},
		{
			X:        0,
			Right:    5,
			Down:     1,
			NumTrees: 0,
		},
		{
			X:        0,
			Right:    7,
			Down:     1,
			NumTrees: 0,
		},
		{
			X:        0,
			Right:    1,
			Down:     2,
			NumTrees: 0,
		},
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lineNum = 1
	for scanner.Scan() {
		line := scanner.Text()

		for i := range steps {
			if (steps[i].Down == 2 && lineNum%2 != 0) || steps[i].Down == 1 {
				if line[steps[i].X%len(line)] == '#' {
					steps[i].NumTrees++
				}
				steps[i].X += steps[i].Right
			}
		}
		lineNum++
	}

	multiplied := steps[0].NumTrees
	for i, step := range steps {
		fmt.Printf("The number of trees was %d\n", step.NumTrees)
		if i > 0 {
			multiplied = multiplied * step.NumTrees
		}
	}

	fmt.Printf("The multipled count is %d\n", multiplied)
}
