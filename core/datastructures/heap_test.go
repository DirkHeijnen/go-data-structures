package advent_of_code

import (
	"testing"
)

func TestMinHeap_InsertAndRemove(t *testing.T) {
	heap := MinHeap[int]{}

	valuesToAdd := []int{5, 3, 10, 1, 4}
	expectedOrder := []int{1, 3, 4, 5, 10} // MinHeap should sort these

	for _, v := range valuesToAdd {
		heap.Insert(v)
	}

	for _, expected := range expectedOrder {
		if val := heap.Remove(); val != expected {
			t.Errorf("Expected %d, got %d", expected, val)
		}
	}

	// The heap should be empty now
	if heap.Length() != 0 {
		t.Errorf("Heap should be empty, got size %d", heap.Length())
	}
}

func TestMinHeap_RemoveFromEmpty(t *testing.T) {
	heap := MinHeap[int]{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic when removing from an empty heap")
		}
	}()

	heap.Remove() // This should panic
}

func TestMaxHeap_InsertAndRemove(t *testing.T) {
	heap := MaxHeap[int]{}

	valuesToAdd := []int{5, 3, 10, 1, 4}
	expectedOrder := []int{10, 5, 4, 3, 1} // MaxHeap should sort these

	for _, v := range valuesToAdd {
		heap.Insert(v)
	}

	for _, expected := range expectedOrder {
		if val := heap.Remove(); val != expected {
			t.Errorf("Expected %d, got %d", expected, val)
		}
	}

	if heap.Length() != 0 {
		t.Errorf("Heap should be empty, got size %d", heap.Length())
	}
}

func TestMaxHeap_RemoveFromEmpty(t *testing.T) {
	heap := MaxHeap[int]{}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Expected a panic when removing from an empty heap")
		}
	}()

	heap.Remove() // This should panic
}

// Test for heap with custom comparison function
func TestHeap_CustomComparison(t *testing.T) {
	minCompareFunc := func(a, b int) bool { return a < b }
	maxCompareFunc := func(a, b int) bool { return a > b }

	minHeap := Heap[int]{compareFunction: minCompareFunc}
	maxHeap := Heap[int]{compareFunction: maxCompareFunc}

	valuesToAdd := []int{5, 3, 10, 1, 4}
	for _, v := range valuesToAdd {
		minHeap.Insert(v)
		maxHeap.Insert(v)
	}

	// Test MinHeap
	expectedMinOrder := []int{1, 3, 4, 5, 10}
	for _, expected := range expectedMinOrder {
		if val := minHeap.Remove(); val != expected {
			t.Errorf("MinHeap: Expected %d, got %d", expected, val)
		}
	}

	// Test MaxHeap
	expectedMaxOrder := []int{10, 5, 4, 3, 1}
	for _, expected := range expectedMaxOrder {
		if val := maxHeap.Remove(); val != expected {
			t.Errorf("MaxHeap: Expected %d, got %d", expected, val)
		}
	}
}

func BenchmarkMinHeap_Insert(b *testing.B) {
	var heap MinHeap[int]
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
}

func BenchmarkMinHeap_Remove(b *testing.B) {
	var heap MinHeap[int]
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Remove()
	}
}

func BenchmarkMaxHeap_Insert(b *testing.B) {
	var heap MaxHeap[int]
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
}

func BenchmarkMaxHeap_Remove(b *testing.B) {
	var heap MaxHeap[int]
	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Remove()
	}
}

// Benchmark for heap insert operation
func BenchmarkHeap_Insert(b *testing.B) {
	compareFunc := func(a, b int) bool { return a < b } // MinHeap comparison
	heap := Heap[int]{compareFunction: compareFunc}

	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
}

// Benchmark for heap remove operation
func BenchmarkHeap_Remove(b *testing.B) {
	compareFunc := func(a, b int) bool { return a < b } // MinHeap comparison
	heap := Heap[int]{compareFunction: compareFunc}

	for i := 0; i < b.N; i++ {
		heap.Insert(i)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		heap.Remove()
	}
}
