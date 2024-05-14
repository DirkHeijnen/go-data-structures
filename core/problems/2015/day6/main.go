package main

import (
	. "avent_of_code/core/datastructures"
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Instruction struct {
	action string
	startX int
	startY int
	endX   int
	endY   int
}

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func parseInput(filePath string) []Instruction {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file", err)
		panic(err)
	}

	defer file.Close()

	result := []Instruction{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		splits := strings.Split(line, " ")

		action := ""
		startX := 0
		startY := 0
		endX := 0
		endY := 0

		if len(splits) == 5 { // turn of / turn on
			action = splits[1]
			startX = atoi(strings.Split(splits[2], ",")[0])
			startY = atoi(strings.Split(splits[2], ",")[1])
			endX = atoi(strings.Split(splits[4], ",")[0])
			endY = atoi(strings.Split(splits[4], ",")[1])
		} else { // toggle
			action = splits[0]
			startX = atoi(strings.Split(splits[1], ",")[0])
			startY = atoi(strings.Split(splits[1], ",")[1])
			endX = atoi(strings.Split(splits[3], ",")[0])
			endY = atoi(strings.Split(splits[3], ",")[1])
		}

		result = append(result, Instruction{
			action: action,
			startX: startX,
			startY: startY,
			endX:   endX,
			endY:   endY,
		})
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return result
}

func solvePart1(table *Table[bool], instructions []Instruction) int {
	// Reset table to all values on false.
	table.SetAll(false)

	for _, instruction := range instructions {
		for row := instruction.startY; row <= instruction.endY; row++ {
			for column := instruction.startX; column <= instruction.endX; column++ {

				if instruction.action == "on" {
					table.Set(row, column, true)
					continue
				}

				if instruction.action == "off" {
					table.Set(row, column, false)
					continue
				}

				if instruction.action == "toggle" {
					table.Set(row, column, !table.Get(row, column))
				}

			}
		}
	}

	count := 0

	for row := 0; row < table.GetRowCount(); row++ {
		for column := 0; column < table.GetColumnCount(); column++ {
			if table.Get(row, column) == true {
				count++
			}
		}
	}

	return count
}

func solvePart2(table *Table[int], instructions []Instruction) int {
	// Reset table to all values on false.
	table.SetAll(0)

	for _, instruction := range instructions {
		for row := instruction.startY; row <= instruction.endY; row++ {
			for column := instruction.startX; column <= instruction.endX; column++ {

				if instruction.action == "on" {
					table.Set(row, column, table.Get(row, column)+1)
					continue
				}

				if instruction.action == "off" && table.Get(row, column) > 0 {
					table.Set(row, column, table.Get(row, column)-1)
					continue
				}

				if instruction.action == "toggle" {
					table.Set(row, column, table.Get(row, column)+2)
				}

			}
		}
	}

	count := 0

	for row := 0; row < table.GetRowCount(); row++ {
		for column := 0; column < table.GetColumnCount(); column++ {
			count += table.Get(row, column)
		}
	}

	return count
}

func main() {
	instructions := parseInput("input.txt")

	startTime := time.Now()
	solution := solvePart1(NewTable[bool](1000, 1000), instructions)
	elapsedTime := time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 1 is: %s\n", elapsedTime)

	startTime = time.Now()
	solution = solvePart2(NewTable[int](1000, 1000), instructions)
	elapsedTime = time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 2 is: %s\n", elapsedTime)
}
