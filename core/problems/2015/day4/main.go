package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"
	"time"
)

func getMD5Hash(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func solvePart1(input string) int {
	counter := 0
	hash := getMD5Hash(input + strconv.Itoa(counter))

	for !strings.HasPrefix(hash, "00000") {
		counter++
		hash = getMD5Hash(input + strconv.Itoa(counter))
	}

	return counter
}

func solvePart2(input string) int {
	counter := 0
	hash := getMD5Hash(input + strconv.Itoa(counter))

	for !strings.HasPrefix(hash, "000000") {
		counter++
		hash = getMD5Hash(input + strconv.Itoa(counter))
	}

	return counter
}

func main() {
	startTime := time.Now()
	solution := solvePart1("iwrupvqb")
	elapsedTime := time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 1 is: %s\n", elapsedTime)

	startTime = time.Now()
	solution = solvePart2("iwrupvqb")
	elapsedTime = time.Since(startTime)

	fmt.Printf("The solution for part 1 is %d\n", solution)
	fmt.Printf("Execution time for part 2 is: %s\n", elapsedTime)
}
