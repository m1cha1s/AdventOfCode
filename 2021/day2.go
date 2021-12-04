package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func day2() {
	file, _ := os.Open("data/day2.txt")
	defer file.Close()

	pos := 0
	depth := 0
	aim := 0

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		full_command := scanner.Text()

		splited_command := strings.Split(full_command, " ")

		command := splited_command[0]
		command_args, _ := strconv.Atoi(splited_command[1])

		switch command {
		case "forward":
			pos += command_args
			depth += aim * command_args
		case "up":
			aim -= command_args
		case "down":
			aim += command_args
		}
	}
	fmt.Printf("pos: %d depth: %d result: %d \n", pos, depth, pos*depth)
}
