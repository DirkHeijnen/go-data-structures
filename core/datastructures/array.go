package advent_of_code

import (
	"fmt"
	"strings"
)

type Array struct {
	elements []interface{}
}

func New(values ...interface{}) *Array {
	array := &Array{}

	if len(values) > 0 {
		array.Push(values)
	}

	return array
}

// Array methods.
func (array *Array) Get(index int) interface{} {
	return array.elements[index]
}

func (array *Array) Remove(index int) {
	array.elements = append(array.elements[:index], array.elements[index+1:]...)
}

func (array *Array) Contains(values ...interface{}) bool {
	for _, searchValue := range values {
		found := false

		for index := 0; index < array.Size(); index++ {
			if array.elements[index] == searchValue {
				found = true
				break
			}
		}

		if !found {
			return false
		}
	}

	return true
}

func (array *Array) IndexOf(value interface{}) int {
	for index, item := range array.elements {
		if item == value {
			return index
		}
	}

	return -1
}

func (array *Array) Swap(index1 int, index2 int) {
	array.elements[index1], array.elements[index2] = array.elements[index2], array.elements[index1]
}

func (array *Array) Push(values ...interface{}) {
	array.elements = append(array.elements, values...)
}

func (array *Array) Shift(values ...interface{}) {
	array.elements = append(values, array.elements...)
}

func (array *Array) Pop() interface{} {
	value := array.Get(array.Size() - 1)
	array.elements = array.elements[:array.Size()-1]

	return value
}

func (array *Array) Unshift() interface{} {
	value := array.Get(0)
	array.elements = array.elements[1:] // *: Removes the first element with fast time complexity.

	return value
}

func (array *Array) Map(function func(index int, element interface{}) interface{}) {
	for index, element := range array.elements {
		array.elements[index] = function(index, element)
	}
}

func (array *Array) Filter(function func(index int, element interface{}) bool) *Array {
	newList := New()

	for index, item := range array.elements {
		if function(index, item) {
			newList.Push(item)
		}
	}

	return newList
}

// Container Methods
func (array *Array) IsEmpty() bool {
	return len(array.elements) == 0
}

func (array *Array) Size() int {
	return len(array.elements)
}

func (array *Array) Clear() {
	array.elements = []interface{}{}
}

func (array *Array) ToString() string {
	str := "ArrayList\n"
	values := []string{}
	for _, value := range array.elements[:array.Size()] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}
