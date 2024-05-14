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

func atoi(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
}

func parseInput(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file", err)
		panic(err)
	}

	defer file.Close()

	result := ""

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = strings.TrimSpace(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return result
}

func move(x, y int, direction rune) (int, int) {
	switch direction {
	case '<':
		x--
	case '>':
		x++
	case '^':
		y--
	case 'v':
		y++
	}
	return x, y
}

func deliver(houses *HashMap[string, int], x, y int) {
	key := strconv.Itoa(x) + "," + strconv.Itoa(y)
	value, exists := houses.Get(key)

	if exists {
		houses.Set(key, value+1)
	} else {
		houses.Set(key, 1)
	}
}

func solvePart1(instructions string) int {
	houseMap := NewMap[string, int]()
	santaX, santaY := 0, 0

	// Set initial house with one present
	deliver(houseMap, santaX, santaY)

	for _, direction := range instructions {
		santaX, santaY = move(santaX, santaY, direction)
		deliver(houseMap, santaX, santaY)
	}

	return houseMap.Size()
}

func solvePart2(instructions string) int {
	houseMap := NewMap[string, int]()
	santaX, santaY := 0, 0
	robotX, robotY := 0, 0

	// Set initial house with two presents
	deliver(houseMap, santaX, santaY)
	deliver(houseMap, robotX, robotY)

	for index, direction := range instructions {
		if index%2 == 0 {
			santaX, santaY = move(santaX, santaY, direction)
			deliver(houseMap, santaX, santaY)
		} else {
			robotX, robotY = move(robotX, robotY, direction)
			deliver(houseMap, robotX, robotY)
		}
	}

	return houseMap.Size()
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
