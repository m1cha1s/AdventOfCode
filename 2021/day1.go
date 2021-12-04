package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func day1() {
	file, _ := os.Open("data/day1.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	sums := []int{}

	data := []int{}

	for scanner.Scan() {
		lineInt, _ := strconv.Atoi(scanner.Text())
		data = append(data, lineInt)
	}

	for i := 0; i < len(data)-3; i++ {
		sum := 0

		for j := 0; j < 3; j++ {
			sum += data[i+j]
		}

		sums = append(sums, sum)
	}

	prev := sums[0]

	smaller := 0

	for i := 1; i < len(sums); i++ {
		if sums[i] > prev {
			smaller++
		}
		prev = sums[i]

	}

	fmt.Printf("Sonar readout: %d\n", smaller+1)
}
