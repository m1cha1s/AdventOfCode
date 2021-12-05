package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const MAX_SIZE int = 1000

type ventMap struct {
	Map [][]int
}

func day5() {
	file, _ := os.Open("data/day5.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	Map := MakeMap(scanner)
	// DumpMap(Map)

	fmt.Println(FindDanger(Map))
}

func MakeMap(scanner *bufio.Scanner) *ventMap {
	vMap := ventMap{}

	for y := 0; y < MAX_SIZE; y++ {
		vMap.Map = append(vMap.Map, []int{})
		for x := 0; x < MAX_SIZE; x++ {
			vMap.Map[y] = append(vMap.Map[y], 0)
		}
	}

	for scanner.Scan() {
		text := scanner.Text()

		beStr := strings.Split(text, " -> ")

		bStr := strings.Split(beStr[0], ",")
		eStr := strings.Split(beStr[1], ",")

		bx, _ := strconv.Atoi(bStr[0])
		by, _ := strconv.Atoi(bStr[1])

		ex, _ := strconv.Atoi(eStr[0])
		ey, _ := strconv.Atoi(eStr[1])

		if bx == ex {
			start := 0
			stop := 0
			if by > ey {
				start = ey
				stop = by
			} else {
				start = by
				stop = ey
			}
			for y := start; y <= stop; y++ {
				vMap.Map[y][bx]++
			}
		} else if by == ey {
			start := 0
			stop := 0
			if bx > ex {
				start = ex
				stop = bx
			} else {
				start = bx
				stop = ex
			}
			for x := start; x <= stop; x++ {
				vMap.Map[by][x]++
			}
		} else {
			dirX := (ex - bx) / int(math.Abs(float64(ex-bx)))
			dirY := (ey - by) / int(math.Abs(float64(ey-by)))

			x := bx
			y := by

			for {
				vMap.Map[y][x]++
				if x == ex && y == ey {
					break
				}
				x += dirX
				y += dirY
			}
		}
	}

	return &vMap
}

func DumpMap(Map *ventMap) {
	for _, row := range Map.Map {
		for _, pos := range row {
			if pos != 0 {
				fmt.Printf("%d ", pos)
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println("")
	}
}

func FindDanger(Map *ventMap) int {
	dangerCount := 0
	for _, row := range Map.Map {
		for _, pos := range row {
			if pos > 1 {
				dangerCount++
			}
		}
	}
	return dangerCount
}
