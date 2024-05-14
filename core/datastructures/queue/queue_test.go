package advent_of_code

import (
	"testing"
)

// TestEnqueueDequeueFIFO tests the FIFO behavior of the queue
func TestEnqueueDequeueFIFO(t *testing.T) {
	queue := Queue[int]{}

	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)

	// Dequeue should return items in the order 1, 2, 3
	if value := queue.Dequeue(); value != 1 {
		t.Errorf("Expected 1, got %d", value)
	}
	if value := queue.Dequeue(); value != 2 {
		t.Errorf("Expected 2, got %d", value)
	}
	if value := queue.Dequeue(); value != 3 {
		t.Errorf("Expected 3, got %d", value)
	}
}

// TestEmptyQueueDequeue tests the behavior when dequeuing from an empty queue
func TestEmptyQueueDequeue(t *testing.T) {
	queue := Queue[int]{}

	defer func() {
		if r := recover(); r != nil {
			if errMsg, ok := r.(string); ok && errMsg != "Cannot call Dequeue() on an empty Queue" {
				t.Errorf("Unexpected panic message: %v", r)
			}
		} else {
			t.Errorf("The code did not panic")
		}
	}()

	// This should cause a panic
	queue.Dequeue()
}

// TestIsEmpty checks the IsEmpty method for both empty and non-empty queues
func TestIsEmpty(t *testing.T) {
	emptyQueue := Queue[int]{}
	if !emptyQueue.IsEmpty() {
		t.Errorf("Expected empty queue to be empty, got not empty")
	}
}

// TestLength checks the Length method for various queue states
func TestLength(t *testing.T) {
	queue := Queue[int]{}

	if queue.Length() != 0 {
		t.Errorf("Expected length 0 for new queue, got %d", queue.Length())
	}

	queue.Enqueue(1)
	if queue.Length() != 1 {
		t.Errorf("Expected length 1 after enqueue, got %d", queue.Length())
	}

	queue.Enqueue(2)
	if queue.Length() != 2 {
		t.Errorf("Expected length 2 after enqueue, got %d", queue.Length())
	}

	queue.Dequeue()
	if queue.Length() != 1 {
		t.Errorf("Expected length 1 after dequeue, got %d", queue.Length())
	}
}

// TestPeekFIFO checks the Peek method for a FIFO queue
func TestPeek(t *testing.T) {
	queue := Queue[int]{}
	queue.Enqueue(1)
	queue.Enqueue(2)

	// Peek should return the first element (1) without removing it
	if val := queue.Peek(); val != 1 {
		t.Errorf("Expected Peek to return 1, got %d", val)
	}

	// Queue length should remain the same after Peek
	if length := queue.Length(); length != 2 {
		t.Errorf("Expected queue length to remain 2, got %d", length)
	}
}

// TestPeekEmptyQueue checks the behavior of Peek on an empty queue
func TestPeekEmptyQueue(t *testing.T) {
	queue := Queue[int]{}

	defer func() {
		if r := recover(); r != nil {
			if errMsg, ok := r.(string); ok && errMsg != "Cannot call Peek() on an empty Queue" {
				t.Errorf("Unexpected panic message: %v", r)
			}
		} else {
			t.Errorf("The code did not panic on Peek with an empty queue")
		}
	}()

	// This should cause a panic
	queue.Peek()
}

// TestEnqueueWithMaxSize tests the Enqueue method with a maximum size set
func TestEnqueueWithMaxSize(t *testing.T) {
	queue := Queue[int]{maxSize: 2}

	queue.Enqueue(1)
	if queue.IsFull() {
		t.Errorf("Queue should not be full after one enqueue")
	}

	queue.Enqueue(2)
	if !queue.IsFull() {
		t.Errorf("Queue should be full after two enqueues")
	}

	defer func() {
		if r := recover(); r != nil {
			if errMsg, ok := r.(string); ok && errMsg != "Cannot call Enqueue() on a full Queue" {
				t.Errorf("Unexpected panic message: %v", r)
			}
		} else {
			t.Errorf("Enqueue did not panic on a full queue")
		}
	}()

	queue.Enqueue(3) // This should panic
}

// TestEnqueueWithoutMaxSize tests the Enqueue method without a maximum size
func TestEnqueueWithoutMaxSize(t *testing.T) {
	queue := Queue[int]{}

	for i := 0; i < 1000; i++ {
		queue.Enqueue(i)
		if queue.IsFull() {
			t.Errorf("Unlimited size queue should never be full")
		}
	}

	if queue.Length() != 1000 {
		t.Errorf("Expected length 1000, got %d", queue.Length())
	}
}

// TestIsFull tests the IsFull method
func TestIsFull(t *testing.T) {
	queue := Queue[int]{maxSize: 1}

	if queue.IsFull() {
		t.Errorf("New queue should not be full")
	}

	queue.Enqueue(1)
	if !queue.IsFull() {
		t.Errorf("Queue should be full after one enqueue")
	}
}

// TestEnqueue tests the Enqueue methods
func TestEnqueue(t *testing.T) {
	queue := Queue[int]{} // No priority function

	queue.Enqueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)

	// The queue should maintain FIFO order: 3, 1, 2
	expectedOrder := []int{3, 1, 2}
	for i, expected := range expectedOrder {
		if val := queue.Dequeue(); val != expected {
			t.Errorf("At index %d, expected %d, got %d", i, expected, val)
		}
	}
}

func TestPriorityQueue_EnqueueDequeue(t *testing.T) {
	options := PriorityQueueOptions[int]{
		PriorityFunction: func(a, b int) bool { return a < b },
	}
	pq := NewPriorityQueue[int](options)

	pq.Enqueue(5)
	pq.Enqueue(1)
	pq.Enqueue(3)

	if val := pq.Dequeue(); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	if val := pq.Dequeue(); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}

	if val := pq.Dequeue(); val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
}

func TestPriorityQueue_Peek(t *testing.T) {
	options := PriorityQueueOptions[int]{
		PriorityFunction: func(a, b int) bool { return a > b },
	}
	pq := NewPriorityQueue[int](options)

	pq.Enqueue(2)
	pq.Enqueue(5)
	pq.Enqueue(3)

	if val := pq.Peek(); val != 5 {
		t.Errorf("Expected Peek to return 5, got %d", val)
	}
}

// Benchmark for Enqueue operation in a Queue
func BenchmarkQueueEnqueue(b *testing.B) {
	queue := Queue[int]{}

	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
}

// Benchmark for Dequeue operation in a Queue
func BenchmarkQueueDequeue(b *testing.B) {
	queue := Queue[int]{}
	for i := 0; i < b.N; i++ {
		queue.Enqueue(i)
	}
	b.ResetTimer() // Start timing only the dequeue operation
	for !queue.IsEmpty() {
		queue.Dequeue()
	}
}

// BenchmarkPriorityQueueEnqueue benchmarks the Enqueue operation
func BenchmarkPriorityQueueEnqueue(b *testing.B) {
	pq := NewPriorityQueue[int](PriorityQueueOptions[int]{PriorityFunction: func(a, b int) bool { return a < b }})

	for i := 0; i < b.N; i++ {
		pq.Enqueue(i)
	}
}

// BenchmarkPriorityQueueDequeue benchmarks the Dequeue operation
func BenchmarkPriorityQueueDequeue(b *testing.B) {
	pq := NewPriorityQueue[int](PriorityQueueOptions[int]{PriorityFunction: func(a, b int) bool { return a < b }})

	for i := 0; i < b.N; i++ {
		pq.Enqueue(i)
	}
	b.ResetTimer()
	for !pq.IsEmpty() {
		pq.Dequeue()
	}
}
