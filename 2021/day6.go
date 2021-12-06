package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day6() {
	file, _ := os.Open("data/day6.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()

	colonyStrSplit := strings.Split(scanner.Text(), ",")

	colony := []int8{}

	for _, fishString := range colonyStrSplit {
		fish, _ := strconv.Atoi(fishString)
		colony = append(colony, int8(fish))
	}

	fmt.Println(simulateColony(colony, 256))

}

func simulateColony(colony []int8, days int) int {
	fishInterval := make([]int, 9)

	for _, fish := range colony {
		fishInterval[fish]++
	}

	for day := 0; day < days; day++ {
		fishInterval = leftRotation(fishInterval, 9, 1)
		fishInterval[6] += fishInterval[8]
	}

	fishCount := 0
	for _, count := range fishInterval {
		fishCount += count
	}
	return fishCount
}

func leftRotation(a []int, size int, rotation int) []int {

	var newArray []int
	for i := 0; i < rotation; i++ {
		newArray = a[1:size]
		newArray = append(newArray, a[0])
		a = newArray
	}
	return a
}
