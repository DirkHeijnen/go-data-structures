package advent_of_code

import "golang.org/x/exp/constraints"

type Stack[T constraints.Ordered] struct {
	array Array[T]
}

// NewStack creates a new Stack instance
func NewStack[T constraints.Ordered]() *Stack[T] {
	return &Stack[T]{}
}

// Push adds an element to the top of the stack
func (stack *Stack[T]) Push(value T) {
	stack.array.PushBack(value) // * PushBack is done because popping from the back is O(1) and popping from the front is O(n) with slices in Go.
}

// Pop removes and returns the top element of the stack
func (stack *Stack[T]) Pop() T {
	if stack.IsEmpty() {
		panic("Pop from empty stack")
	}

	return stack.array.PopBack() // * PopBack is done because popping from the back is O(1) and popping from the front is O(n) with slices in Go.
}

// Peek returns the top element of the stack without removing it
func (stack *Stack[T]) Peek() T {
	if stack.IsEmpty() {
		panic("Peek from empty stack")
	}

	return stack.array.Read(stack.Size() - 1)
}

// Size returns the number of elements in the stack
func (stack *Stack[T]) Size() int {
	return stack.array.Length()
}

// IsEmpty checks if the stack is empty
func (stack *Stack[T]) IsEmpty() bool {
	return stack.Size() == 0
}
