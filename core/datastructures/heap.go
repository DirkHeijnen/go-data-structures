package advent_of_code

import "golang.org/x/exp/constraints"

type MaxHeap[T constraints.Ordered] struct {
	array Array[T]
}

type Heap[T constraints.Ordered] struct {
	array           Array[T]
	compareFunction func(T, T) bool
}

func NewMaxHeap[T constraints.Ordered]() *MaxHeap[T] {
	return &MaxHeap[T]{
		array: Array[T]{},
	}
}

func NewHeap[T constraints.Ordered](compareFunction func(T, T) bool) *Heap[T] {
	return &Heap[T]{
		array:           Array[T]{},
		compareFunction: compareFunction,
	}
}

func (heap *MaxHeap[T]) Insert(value T) {
	heap.array.PushBack(value)
	heap.HeapifyUp(heap.array.Length() - 1)
}

func (heap *Heap[T]) Insert(value T) {
	heap.array.PushBack(value)
	heap.HeapifyUp(heap.array.Length() - 1)
}

func (heap *MaxHeap[T]) Peek() T {
	if heap.IsEmpty() {
		panic("Cannot call Peek() on a empty MaxHeap")
	}

	return heap.array.Read(0)
}

func (heap *Heap[T]) Peek() T {
	if heap.IsEmpty() {
		panic("Cannot call Peek() on a empty Heap")
	}

	return heap.array.Read(0)
}

func (heap *MaxHeap[T]) Remove() T {
	if heap.IsEmpty() {
		panic("Cannot call Remove() on an empty Heap")
	}

	value := heap.array.Read(0)
	lastIndex := heap.array.Length() - 1

	heap.array.data[0] = heap.array.data[lastIndex]
	heap.array.data = heap.array.data[:lastIndex]
	heap.HeapifyDown(0)

	return value
}

func (heap *Heap[T]) Remove() T {
	if heap.IsEmpty() {
		panic("Cannot call Remove() on an empty Heap")
	}

	value := heap.array.Read(0)
	lastIndex := heap.array.Length() - 1

	heap.array.data[0] = heap.array.data[lastIndex]
	heap.array.data = heap.array.data[:lastIndex]
	heap.HeapifyDown(0)

	return value
}

func (heap *MaxHeap[T]) HeapifyUp(index int) {
	for index > 0 && heap.array.Read(parent(index)) < heap.array.Read(index) {
		heap.array.Swap(parent(index), index)
		index = parent(index)
	}
}

func (heap *Heap[T]) HeapifyUp(index int) {
	for index > 0 && heap.compareFunction(heap.array.Read(index), heap.array.Read(parent(index))) {
		heap.array.Swap(parent(index), index)
		index = parent(index)
	}
}

func (heap *MaxHeap[T]) HeapifyDown(index int) {
	lastIndex := heap.array.Length() - 1
	leftIndex := left(index)
	rightIndex := right(index)

	childToCompare := 0

	for leftIndex <= lastIndex {
		if leftIndex == lastIndex {
			childToCompare = leftIndex
		} else if heap.array.Read(leftIndex) > heap.array.Read(rightIndex) {
			childToCompare = leftIndex
		} else {
			childToCompare = rightIndex
		}

		if heap.array.Read(index) < heap.array.Read(childToCompare) {
			heap.array.Swap(index, childToCompare)
			index = childToCompare
			leftIndex = left(index)
			rightIndex = right(index)
		} else {
			return
		}
	}
}

func (heap *Heap[T]) HeapifyDown(index int) {
	lastIndex := heap.array.Length() - 1
	leftIndex := left(index)
	rightIndex := right(index)

	childToCompare := 0

	for leftIndex <= lastIndex {
		childToCompare = leftIndex
		if rightIndex <= lastIndex && heap.compareFunction(heap.array.Read(rightIndex), heap.array.Read(leftIndex)) {
			childToCompare = rightIndex
		}

		// Compare and swap if necessary
		if heap.compareFunction(heap.array.Read(childToCompare), heap.array.Read(index)) {
			heap.array.Swap(index, childToCompare)
			index = childToCompare
			leftIndex, rightIndex = left(index), right(index)
		} else {
			return
		}
	}
}

func (heap *MaxHeap[T]) Length() int {
	return heap.array.Length()
}

func (heap *Heap[T]) Length() int {
	return heap.array.Length()
}

func (heap *MaxHeap[T]) IsEmpty() bool {
	return heap.Length() == 0
}

func (heap *Heap[T]) IsEmpty() bool {
	return heap.Length() == 0
}
