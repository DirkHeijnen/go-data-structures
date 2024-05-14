package advent_of_code

import (
	"testing"
)

func TestEmptyArrayIsLengthZero(t *testing.T) {
	array := Array[int]{}

	if array.Length() != 0 {
		t.Errorf("Got %d, wanted 0", array.Length())
	}
}

func TestEmptyArrayIsEmpty(t *testing.T) {
	array := Array[int]{}

	if array.IsEmpty() == false {
		t.Errorf("Got %t, wanted true", array.IsEmpty())
	}
}

func TestNonEmptyArrayIsNotEmpty(t *testing.T) {
	array := Array[int]{}
	array.PushBack(1)

	if array.IsEmpty() == true {
		t.Errorf("Got %t, wanted false", array.IsEmpty())
	}
}

func TestEmptyArrayAfterOnePushBackHasLength1(t *testing.T) {
	array := Array[int]{}
	array.PushBack(1)

	if array.Length() != 1 {
		t.Errorf("Got %d, wanted 1", array.Length())
	}
}

func TestEmptyArrayAfterTwoPushesHasLength2(t *testing.T) {
	array := Array[int]{}
	array.PushBack(1)
	array.PushBack(2)

	if array.Length() != 2 {
		t.Errorf("Got %d, wanted 2", array.Length())
	}
}

func TestEmptyArrayAfterPushingTwoItemsHasLength2(t *testing.T) {
	array := Array[int]{}
	array.PushBack(1, 2)

	if array.Length() != 2 {
		t.Errorf("Got %d, wanted 2", array.Length())
	}
}

func TestEmptyArrayAfterPushingSliceHasLengthOfSlice(t *testing.T) {
	array := Array[int]{}
	slice := []int{}

	slice = append(slice, 1, 2)
	array.PushBack(slice...)

	if array.Length() != 2 {
		t.Errorf("Got %d, wanted 2", array.Length())
	}
}

func TestEmptyArrayAfterPushingBackTwoItemsHasCorrectOrder(t *testing.T) {
	array := Array[int]{}

	array.PushBack(1, 2)
	array.PushBack(3, 4)

	if array.Read(0) != 1 || array.Read(1) != 2 || array.Read(2) != 3 || array.Read(3) != 4 {
		t.Errorf("Got %+v, wanted [1, 2, 3, 4]", array)
	}
}

func TestEmptyArrayAfterPushingFrontHasCorrectOrder(t *testing.T) {
	array := Array[int]{}

	array.PushFront(1, 2)
	array.PushFront(3, 4)

	if array.Read(0) != 3 || array.Read(1) != 4 || array.Read(2) != 1 || array.Read(3) != 2 {
		t.Errorf("Got %+v, wanted [3, 4, 1, 2]", array)
	}
}

func TestMapWorksWithMultiplyBy2(t *testing.T) {
	array := Array[int]{}

	array.PushBack(2, 4)
	array.Map(func(index int, item int) int { return item * 2 })

	if array.Read(0) != 4 || array.Read(1) != 8 {
		t.Errorf("Got %+v, wanted [4, 8]", array)
	}
}

func TestMapWorksWithStringDoubling(t *testing.T) {
	array := Array[string]{}
	array.PushBack("hello", "world")
	array.Map(func(index int, item string) string { return item + item })

	if array.Read(0) != "hellohello" || array.Read(1) != "worldworld" {
		t.Errorf("Got %+v, wanted [hellohello, worldworld]", array)
	}
}

func TestIndexOfWorksWithFoundInteger(t *testing.T) {
	array := Array[int]{}

	array.PushBack(2, 4)
	index := array.IndexOf(4)

	if index != 1 {
		t.Errorf("Got %+v, wanted 1", index)
	}
}

func TestIndexOfWorksWithNotFoundInteger(t *testing.T) {
	array := Array[int]{}

	array.PushBack(2, 4)
	index := array.IndexOf(8)

	if index != -1 {
		t.Errorf("Got %+v, wanted -1", index)
	}
}

func TestMinWorks(t *testing.T) {
	array := Array[int]{}

	array.PushBack(8, 3, 4, 12, 1, 7)
	minValue := array.Min()

	if minValue != 1 {
		t.Errorf("Got %+v, wanted 1", minValue)
	}
}

func TestMaxWorks(t *testing.T) {
	array := Array[int]{}

	array.PushBack(8, 3, 4, 12, 1, 7)
	maxValue := array.Max()

	if maxValue != 12 {
		t.Errorf("Got %+v, wanted 12", maxValue)
	}
}

func TestFilter(t *testing.T) {
	// Arrange
	array := Array[int]{}
	array.PushBack(8, 3, 4, 12, 1, 7)

	// Act
	newArray := array.Filter(func(index int, item int) bool { return item < 8 })

	// Assert
	if newArray.Length() != 4 {
		t.Errorf("Got %+v, wanted 4", newArray.Length())
	}

	if newArray.Read(0) != 3 || newArray.Read(1) != 4 || newArray.Read(2) != 1 || newArray.Read(3) != 7 {
		t.Errorf("Got %+v, wanted [3, 4, 1, 7]", newArray)
	}

	if array.Length() != 6 {
		t.Errorf("Got %+v, wanted 6", array.Length())
	}

	if array.Read(0) != 8 || array.Read(1) != 3 || array.Read(2) != 4 || array.Read(3) != 12 || array.Read(4) != 1 || array.Read(5) != 7 {
		t.Errorf("Got %+v, wanted [8, 3, 4, 12, 1, 7]", array.data)
	}
}

func TestArray_Remove(t *testing.T) {
	array := Array[int]{data: []int{1, 2, 3, 4, 5}}

	// Remove the element at index 2 (value 3)
	array.Remove(2)

	expected := []int{1, 2, 4, 5}
	if len(array.data) != len(expected) {
		t.Errorf("Expected length %d, got %d", len(expected), len(array.data))
	}

	for i, v := range expected {
		if array.data[i] != v {
			t.Errorf("At index %d, expected %d, got %d", i, v, array.data[i])
		}
	}
}

func TestArray_PopFront(t *testing.T) {
	array := Array[int]{data: []int{1, 2, 3, 4, 5}}

	value := array.PopFront()
	if value != 1 {
		t.Errorf("PopFront: expected 1, got %d", value)
	}

	expected := []int{2, 3, 4, 5}
	if !slicesEqual(array.data, expected) {
		t.Errorf("PopFront: expected array %v, got %v", expected, array.data)
	}
}

func TestArray_PopBack(t *testing.T) {
	array := Array[int]{data: []int{1, 2, 3, 4, 5}}

	value := array.PopBack()
	if value != 5 {
		t.Errorf("PopBack: expected 5, got %d", value)
	}

	expected := []int{1, 2, 3, 4}
	if !slicesEqual(array.data, expected) {
		t.Errorf("PopBack: expected array %v, got %v", expected, array.data)
	}
}

// Helper function to compare slices
func slicesEqual[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
