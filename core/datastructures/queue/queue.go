package advent_of_code

import "golang.org/x/exp/constraints"

// Queue represents a generic FIFO (First-In-First-Out) queue structure.
type Queue[T constraints.Ordered] struct {
	array   Array[T]
	maxSize int // 0 for unlimited size.
}

// QueueOptions give you the ability to set the functionality of the queue.
type QueueOptions[T constraints.Ordered] struct {
	MaxSize int
}

// PriorityQueue represents a generic priority queue where elements are sorted based on a priority function.
type PriorityQueue[T constraints.Ordered] struct {
	heap             Heap[T]
	maxSize          int // 0 for unlimited size.
	priorityFunction func(T, T) bool
}

// PriorityQueueOptions give you the ability to set the functionality of the queue.
type PriorityQueueOptions[T constraints.Ordered] struct {
	MaxSize          int
	PriorityFunction func(T, T) bool
}

// NewPriorityQueue creates a new instance of a priority queue with the specified options.
// The options include maxSize.
func NewQueue[T constraints.Ordered](opts ...QueueOptions[T]) *Queue[T] {
	var opt QueueOptions[T]
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = QueueOptions[T]{MaxSize: 0}
	}

	return &Queue[T]{
		array:   *&Array[T]{},
		maxSize: opt.MaxSize,
	}
}

// NewPriorityQueue creates a new instance of a priority queue with the specified options.
// The options can include maxSize and a custom priority function.
func NewPriorityQueue[T constraints.Ordered](opts ...PriorityQueueOptions[T]) *PriorityQueue[T] {
	var opt PriorityQueueOptions[T]
	if len(opts) > 0 {
		opt = opts[0]
	} else {
		opt = PriorityQueueOptions[T]{MaxSize: 0, PriorityFunction: nil}
	}

	return &PriorityQueue[T]{
		heap:             *NewHeap(opt.PriorityFunction),
		maxSize:          opt.MaxSize,
		priorityFunction: opt.PriorityFunction,
	}
}

// Enqueue adds a new element to the end of the queue.
// It panics if the queue has reached its maximum size.
func (queue *Queue[T]) Enqueue(value T) {
	if queue.maxSize != 0 && queue.array.Length() >= queue.maxSize {
		panic("Cannot call Enqueue() on a full Queue")
	}

	queue.array.PushBack(value)
}

// Enqueue adds a new element to the priority queue in a position based on its priority.
// It panics if the priority queue has reached its maximum size.
func (queue *PriorityQueue[T]) Enqueue(value T) {
	if queue.maxSize != 0 && queue.heap.Length() >= queue.maxSize {
		panic("Cannot call Enqueue() on a full PriorityQueue")
	}

	queue.heap.Insert(value)
}

// Dequeue removes and returns the element at the front of the queue.
// It panics if the queue is empty.
func (queue *Queue[T]) Dequeue() T {
	if queue.IsEmpty() {
		panic("Cannot call Dequeue() on an empty Queue")
	}

	return queue.array.PopFront()
}

// Dequeue removes and returns the element with the highest priority from the priority queue.
// It panics if the priority queue is empty.
func (queue *PriorityQueue[T]) Dequeue() T {
	if queue.IsEmpty() {
		panic("Cannot call Dequeue() on an empty PriorityQueue")
	}

	return queue.heap.Remove()
}

// IsFull checks if the queue has reached its maximum size.
func (queue *Queue[T]) IsFull() bool {
	if queue.maxSize == 0 {
		return false
	}

	return queue.Length() >= queue.maxSize
}

// IsFull checks if the queue has reached its maximum size.
func (queue *PriorityQueue[T]) IsFull() bool {
	if queue.maxSize == 0 {
		return false
	}

	return queue.heap.Length() >= queue.maxSize
}

// Peek returns the element at the front of the queue without removing it.
// It panics if the queue is empty.
func (queue *Queue[T]) Peek() T {
	if queue.IsEmpty() {
		panic("Cannot call Peek() on an empty Queue")
	}

	return queue.array.Read(0)
}

// Peek returns the element at the front of the queue without removing it.
// It panics if the queue is empty.
func (queue *PriorityQueue[T]) Peek() T {
	if queue.IsEmpty() {
		panic("Cannot call Peek() on an empty PriorityQueue")
	}

	return queue.heap.Peek()
}

// IsEmpty checks if the queue is empty.
func (queue *Queue[T]) IsEmpty() bool {
	return queue.array.IsEmpty()
}

// IsEmpty checks if the queue is empty.
func (queue *PriorityQueue[T]) IsEmpty() bool {
	return queue.heap.IsEmpty()
}

// Length returns the number of elements in the queue.
func (queue *Queue[T]) Length() int {
	return queue.array.Length()
}

// Length returns the number of elements in the queue.
func (queue *PriorityQueue[T]) Length() int {
	return queue.heap.Length()
}
