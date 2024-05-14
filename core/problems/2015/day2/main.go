package main

import (
	. "avent_of_code/core/geometry"
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

func parseInput(filePath string) []Cuboid[int] {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error opening file", err)
		panic(err)
	}

	defer file.Close()

	cuboids := []Cuboid[int]{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)

		splits := strings.Split(line, "x")
		cuboids = append(cuboids, *NewCuboid[int](atoi(splits[0]), atoi(splits[1]), atoi(splits[2])))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from file:", err)
	}

	return cuboids
}

func solvePart1(cuboids []Cuboid[int]) int {
	sum := 0

	for _, cuboid := range cuboids {
		sum += cuboid.GetSurfaceArea() + cuboid.GetAreaOfSmallestSide()
	}

	return sum
}

func solvePart2(cuboids []Cuboid[int]) int {
	sum := 0

	for _, cuboid := range cuboids {
		sum += cuboid.GetVolume() + cuboid.GetSmallestPerimeter()
	}

	return sum
}

func main() {
	cuboids := parseInput("input.txt")

	startTime := time.Now()
	solution := solvePart1(cuboids)
	elapsedTime := time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 1 is: %s\n", elapsedTime)

	startTime = time.Now()
	solution = solvePart2(cuboids)
	elapsedTime = time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 2 is: %s\n", elapsedTime)
}
