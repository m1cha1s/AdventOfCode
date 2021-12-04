package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board struct {
	numbers [][]int
	drawn   [][]bool
}

func day4() {
	file, _ := os.Open("data/day4.txt")
	defer file.Close()

	scanner := bufio.NewScanner(file)

	scanner.Scan()
	numbersToDraw := []int{}
	numbersToDrawStr := strings.Split(scanner.Text(), ",")
	for _, number := range numbersToDrawStr {
		number, _ := strconv.Atoi(number)
		numbersToDraw = append(numbersToDraw, number)
	}
	fmt.Println(numbersToDraw)

	boards := []Board{}

	scanner.Scan() // Skip the first spacing

	board := Board{}

	row := 0

	for scanner.Scan() {
		text := scanner.Text()
		// fmt.Println(text)
		if text == "" {
			boards = append(boards, board)
			board = Board{}
			row = 0
			continue
		}

		board.numbers = append(board.numbers, []int{})
		board.drawn = append(board.drawn, []bool{false, false, false, false, false})
		for _, splitText := range strings.Split(text, " ") {
			number, err := strconv.Atoi(splitText)
			if err != nil {
				continue
			}
			board.numbers[row] = append(board.numbers[row], number)
		}

		row++
	}

	won := []int{}
	lastScore := 0

	for _, number := range numbersToDraw {
		for _, board := range boards {
			for boardRowIdx, boardRow := range board.numbers {
				for boardCollumnIdx, boardCollumn := range boardRow {
					if boardCollumn == number {
						board.drawn[boardRowIdx][boardCollumnIdx] = true
					}
				}
			}
		}
		for boardIdx, board := range boards {
			alreadyWon := false
			for _, w := range won {
				if w == boardIdx {
					alreadyWon = true
				}
			}
			if alreadyWon {
				continue
			}
			if checkWin(board) {
				won = append(won, boardIdx)
				lastScore = calculateScore(board, number)
			}
		}
	}

	fmt.Println(won[len(won)-1], boards[won[len(won)-1]], lastScore)

}

func checkWin(board Board) bool {
	for _, boardRow := range board.drawn {
		allRow := true
		for _, boardCollumn := range boardRow {
			if !boardCollumn {
				allRow = false
			}
		}
		if allRow {
			return true
		}

	}
	for boardCollumnIdx := 0; boardCollumnIdx < len(board.drawn[0]); boardCollumnIdx++ {
		allColumn := true
		for _, boardRow := range board.drawn {
			if !boardRow[boardCollumnIdx] {
				allColumn = false
			}
		}
		if allColumn {
			return true
		}
	}
	return false
}

func calculateScore(board Board, winningNum int) int {
	sum := 0
	for rowIdx, row := range board.drawn {
		for collumnIdx, collumn := range row {
			if !collumn {
				sum += board.numbers[rowIdx][collumnIdx]
			}
		}
	}
	return sum * winningNum
}
