package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

type Instruction struct {
	up   bool
	down bool
}

func parseInput(filePath string) []Instruction {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file", err)
		panic(err)
	}

	defer file.Close()

	instructions := []Instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		for _, char := range line {
			if char == '(' {
				instructions = append(instructions, Instruction{up: true, down: false})
			} else {
				instructions = append(instructions, Instruction{up: false, down: true})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return instructions
}

func solvePart1(instructions []Instruction) int {
	counter := 0

	for _, instruction := range instructions {
		if instruction.up {
			counter++
		} else {
			counter--
		}
	}

	return counter
}

func solvePart2(instructions []Instruction) int {
	counter := 0

	for index, instruction := range instructions {
		if instruction.up {
			counter++
		} else {
			counter--
		}

		if counter < 0 {
			return index + 1
		}
	}

	panic("Did not reach a basement floor")
}

func main() {
	instructions := parseInput("input.txt")

	startTime := time.Now()
	solution := solvePart1(instructions)
	elapsedTime := time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 1 is: %s\n", elapsedTime)

	startTime = time.Now()
	solution = solvePart2(instructions)
	elapsedTime = time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 2 is: %s\n", elapsedTime)
}
