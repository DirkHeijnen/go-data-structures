package main

import (
	. "avent_of_code/core/datastructures"
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func parseInput(filePath string) []string {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file", err)
		panic(err)
	}

	defer file.Close()

	result := []string{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, strings.TrimSpace(line))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return result
}

func isNiceStringPart1(text string) bool {
	inputString := NewString(text)

	// Count distinct vowels (aeiou)
	if inputString.GetVowelCount() < 3 {
		return false
	}

	// Count substrings of 2+ consecutive characters (xx, yy, zzz)
	if inputString.HighestConsecutiveCharacterCount() < 2 {
		return false
	}

	// Does not contain ab, cd, pq or xy.
	if inputString.ContainsAny("ab", "cd", "pq", "xy") {
		return false
	}

	return true
}

func isNiceStringPart2(text string) bool {
	condition1 := false
	condition2 := false

	// -1 because I'm considering pairs
	// -2 because first index stops when there is room for another pair with j.
	for i := 0; i < len(text)-1-2; i++ {
		for j := i + 2; j < len(text)-1; j++ {
			if text[i:i+2] == text[j:j+2] {
				condition1 = true
			}
		}
	}

	for i := 0; i < len(text)-2; i++ {
		if text[i] == text[i+2] {
			condition2 = true
		}
	}

	return condition1 && condition2
}

func solvePart1(inputs []string) int {
	counter := 0

	for _, input := range inputs {
		if isNiceStringPart1(input) {
			counter++
		}
	}

	return counter
}

func solvePart2(inputs []string) int {
	counter := 0

	for _, input := range inputs {
		if isNiceStringPart2(input) {
			counter++
		}
	}

	return counter
}

func main() {
	inputs := parseInput("input.txt")

	startTime := time.Now()
	solution := solvePart1(inputs)
	elapsedTime := time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 1 is: %s\n", elapsedTime)

	startTime = time.Now()
	solution = solvePart2(inputs)
	elapsedTime = time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 2 is: %s\n", elapsedTime)
}
