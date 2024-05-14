package advent_of_code

import "testing"

func TestStack_PushPop(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)

	if val := stack.Pop(); val != 3 {
		t.Errorf("Expected 3, got %d", val)
	}
	if val := stack.Pop(); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	if val := stack.Pop(); val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}
}

func TestStack_Peek(t *testing.T) {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)

	if val := stack.Peek(); val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}
	if stack.Size() != 2 {
		t.Errorf("Expected size 2, got %d", stack.Size())
	}
}

func BenchmarkStack_Push(b *testing.B) {
	stack := NewStack[int]()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
}

func BenchmarkStack_Pop(b *testing.B) {
	stack := NewStack[int]()
	for i := 0; i < b.N; i++ {
		stack.Push(i)
	}
	b.ResetTimer()
	for !stack.IsEmpty() {
		stack.Pop()
	}
}
