package advent_of_code

import (
	"strings"
)

type String struct {
	value string
}

// NewStack creates a new Stack instance
func NewString(value string) *String {
	return &String{value: value}
}

func (str *String) StartsWith(value string) {
	strings.HasPrefix(str.value, value)
}

func (str *String) EndWith(value string) {
	strings.HasSuffix(str.value, value)
}

func (str *String) Contains(value string) {
	strings.Contains(str.value, value)
}

func (str *String) ContainsAll(values ...string) bool {
	for _, value := range values {
		if !strings.Contains(str.value, value) {
			return false
		}
	}

	return true
}

func (str *String) ContainsAny(values ...string) bool {
	for _, value := range values {
		if strings.Contains(str.value, value) {
			return true
		}
	}

	return false
}

func (str *String) GetVowelCount() int {
	count := 0
	for _, char := range str.value {
		if char == 'a' || char == 'A' || char == 'e' || char == 'E' || char == 'i' || char == 'I' || char == 'o' || char == 'O' || char == 'u' || char == 'U' {
			count++
		}
	}

	return count
}

func (str *String) GetDistinctVowelCount() int {
	vowels := map[rune]bool{
		'a': false,
		'e': false,
		'i': false,
		'o': false,
		'u': false,
	}

	for _, char := range str.value {
		if _, exists := vowels[char]; exists {
			vowels[char] = true
		}
	}

	count := 0
	for _, found := range vowels {
		if found {
			count++
		}
	}

	return count
}

func (str *String) HighestConsecutiveCharacterCount() int {
	if len(str.value) == 0 {
		return 0
	}

	maxCount := 1
	currentCount := 1

	for i := 1; i < len(str.value); i++ {
		if str.value[i] == str.value[i-1] {
			currentCount++
			if currentCount > maxCount {
				maxCount = currentCount
			}
		} else {
			currentCount = 1
		}
	}
	return maxCount
}
