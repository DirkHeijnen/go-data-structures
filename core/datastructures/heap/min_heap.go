package advent_of_code

import (
	. "avent_of_code/core/datastructures"
	"fmt"
	"strings"
)

type MinHeap struct {
	array Array
}

func NewMinHeap() *MinHeap {
	return &MinHeap{
		array: Array{},
	}
}

// MinHeap Methods.
func (heap *MinHeap) Insert(value interface{}) {
	heap.array.Push(value)
	heap.HeapifyUp(heap.array.Size() - 1)
}

func (heap *MinHeap) Remove() interface{} {
	if heap.IsEmpty() {
		panic("Cannot call Remove() on a empty Heap")
	}

	value := heap.array.Get(0)
	lastIndex := heap.array.Size() - 1

	heap.array.elements[0] = heap.array.elements[lastIndex]
	heap.array.elements = heap.array.elements[:lastIndex]
	heap.HeapifyDown(0)

	return value
}

func (heap *MinHeap) Peek() interface{} {
	if heap.IsEmpty() {
		panic("Cannot call Peek() on a empty MinHeap")
	}

	return heap.array.Get(0)
}

func (heap *MinHeap) HeapifyUp(index int) {
	for heap.array.Get(parent(index)) > heap.array.Get(index) {
		heap.array.Swap(parent(index), index)
		index = parent(index)
	}
}

func (heap *MinHeap) HeapifyDown(index int) {
	lastIndex := heap.array.Size() - 1
	leftIndex := left(index)
	rightIndex := right(index)

	childToCompare := 0

	for leftIndex <= lastIndex {
		if leftIndex == lastIndex {
			childToCompare = leftIndex
		} else if heap.array.Get(leftIndex) < heap.array.Get(rightIndex) {
			childToCompare = leftIndex
		} else {
			childToCompare = rightIndex
		}

		if heap.array.Get(index) > heap.array.Get(childToCompare) {
			heap.array.Swap(index, childToCompare)
			index = childToCompare
			leftIndex = left(index)
			rightIndex = right(index)
		} else {
			return
		}
	}
}

// Container Methods
func (heap *MinHeap) IsEmpty() bool {
	return heap.array.IsEmpty()
}

func (heap *MinHeap) Size() int {
	return heap.array.Size()
}

func (heap *MinHeap) Clear() {
	heap.array.Clear()
}

func (heap *MinHeap) ToString() string {
	str := "MinHeap\n"
	values := []string{}
	for _, value := range heap.array.elements[:heap.array.Size()] {
		values = append(values, fmt.Sprintf("%v", value))
	}
	str += strings.Join(values, ", ")
	return str
}

// get parent index
func parent(i int) int { return (i - 1) / 2 }

// get left child index
func left(i int) int { return 2*i + 1 }

// get right child index
func right(i int) int { return 2*i + 2 }
