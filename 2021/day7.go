package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func day7() {
	file, _ := os.Open("data/day7.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	positionsStr := strings.Split(scanner.Text(), ",")

	positions := []int{}

	for _, positionStr := range positionsStr {
		position, _ := strconv.Atoi(positionStr)
		positions = append(positions, position)
	}

	fmt.Println(positions, maxInArray(positions))

	fuelSpent := []int{}

	for i := 0; i < maxInArray(positions); i++ {
		totalFuelSpent := 0
		for _, position := range positions {
			fuelCost := int(math.Abs(float64(i) - float64(position)))
			totalFuelSpent += (fuelCost * (fuelCost + 1)) / 2
		}
		fuelSpent = append(fuelSpent, totalFuelSpent)
	}

	fmt.Println(fuelSpent, minInArray(fuelSpent))

}

func maxInArray(array []int) int {
	maxVal := -1
	// maxIdx := -1
	for _, val := range array {
		if maxVal < 0 {
			maxVal = val
			// maxIdx = idx
		} else {
			if val > maxVal {
				maxVal = val
				// maxIdx = idx
			}
		}
	}
	return maxVal
}

func minInArray(array []int) int {
	maxVal := -1
	// maxIdx := -1
	for _, val := range array {
		if maxVal < 0 {
			maxVal = val
			// maxIdx = idx
		} else {
			if val < maxVal {
				maxVal = val
				// maxIdx = idx
			}
		}
	}
	return maxVal
}
