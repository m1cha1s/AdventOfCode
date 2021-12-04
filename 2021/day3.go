package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day3() {
	file, _ := os.Open("data/day3.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lines := make(map[int][]string)

	lineNum := 0
	for scanner.Scan() {
		splitted := strings.Split(scanner.Text(), "")
		lines[lineNum] = splitted
		lineNum++
	}
	// fmt.Println(lines)

	gammaString, epsilonString := calculateGammaEpsilonString(lines)

	gamma, _ := strconv.ParseInt(gammaString, 2, 64)
	epsilon, _ := strconv.ParseInt(epsilonString, 2, 64)

	fmt.Printf("Gamma: %d or 0b%b Epsilon: %d or 0b%b Power: %d\n", gamma, gamma, epsilon, epsilon, gamma*epsilon)

	filteredOxygen := make(map[int][]string)

	for key, value := range lines {
		filteredOxygen[key] = value
	}

	for bit := 0; bit < len(lines[0]); bit++ {

		val_sum := 0

		for _, val := range filteredOxygen {
			tmp, _ := strconv.Atoi(val[bit])
			val_sum += tmp
		}

		selected := float64(val_sum) / float64(len(filteredOxygen))
		selectedInt := 0

		if selected >= 0.5 {
			selectedInt = 1
		} else {
			selectedInt = 0
		}

		toDelete := []int{}

		for lineIndex, line := range filteredOxygen {
			if lineInt, _ := strconv.Atoi(line[bit]); lineInt != selectedInt {
				toDelete = append(toDelete, lineIndex)
			}
		}

		for _, index := range toDelete {
			delete(filteredOxygen, index)
		}

		if len(filteredOxygen) == 1 {
			break
		}
	}

	filteredCO2 := make(map[int][]string)

	for key, value := range lines {
		filteredCO2[key] = value
	}

	for bit := 0; bit < len(lines[0]); bit++ {
		// fmt.Println(bit)

		val_sum := 0

		for _, val := range filteredCO2 {
			tmp, _ := strconv.Atoi(val[bit])
			val_sum += tmp
		}

		selected := float64(val_sum) / float64(len(filteredCO2))
		selectedInt := 0

		if selected < 0.5 {
			selectedInt = 1
		} else {
			selectedInt = 0
		}

		toDelete := []int{}

		for lineIndex, line := range filteredCO2 {
			if lineInt, _ := strconv.Atoi(line[bit]); lineInt != selectedInt {
				toDelete = append(toDelete, lineIndex)
			}
		}

		for _, index := range toDelete {
			delete(filteredCO2, index)
		}

		if len(filteredCO2) == 1 {
			break
		}
	}

	var oxygenRating int64
	for _, val := range filteredOxygen {
		tmp, _ := strconv.ParseInt(strings.Join(val, ""), 2, 64)
		oxygenRating = tmp
	}

	var co2Rating int64
	for _, val := range filteredCO2 {
		tmp, _ := strconv.ParseInt(strings.Join(val, ""), 2, 64)
		co2Rating = tmp
	}

	fmt.Printf("Oxygen rating: %d CO2 rating: %d Life support rating %d \n", oxygenRating, co2Rating, oxygenRating*co2Rating)
}

func calculateGammaEpsilonString(lines map[int][]string) (string, string) {
	var ones []int
	for _, line := range lines {
		if len(ones) != len(line) {
			ones = make([]int, len(line))
		}
	}

	for _, line := range lines {
		for bitIndex, bit := range line {
			bit_int, _ := strconv.Atoi(bit)
			if bit_int == 1 {
				ones[bitIndex]++
			}
		}
	}

	gammaString := ""
	epsilonString := ""

	for _, onesCount := range ones {
		if onesCount > len(lines)-onesCount {
			gammaString += "1"
			epsilonString += "0"
		} else {
			gammaString += "0"
			epsilonString += "1"
		}
	}

	return gammaString, epsilonString
}
