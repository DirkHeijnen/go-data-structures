package advent_of_code

import "testing"

func TestTable_SetGet(t *testing.T) {
	table := NewTable[int](3, 3)
	table.Set(1, 1, 5)

	if val := table.Get(1, 1); val != 5 {
		t.Errorf("Expected 5, got %d", val)
	}
}

func TestTable_SetAll(t *testing.T) {
	table := NewTable[int](2, 2)
	table.SetAll(7)

	for i := 0; i < table.GetRowCount(); i++ {
		for j := 0; j < table.GetColumnCount(); j++ {
			if val := table.Get(i, j); val != 7 {
				t.Errorf("Expected 7, got %d at (%d, %d)", val, i, j)
			}
		}
	}
}

func TestTable_Dimensions(t *testing.T) {
	rows, cols := 4, 5
	table := NewTable[float64](rows, cols)

	if table.GetRowCount() != rows {
		t.Errorf("Expected %d rows, got %d", rows, table.GetRowCount())
	}

	if table.GetColumnCount() != cols {
		t.Errorf("Expected %d columns, got %d", cols, table.GetColumnCount())
	}
}

func BenchmarkTable_SetAll(b *testing.B) {
	table := NewTable[int](100, 100)
	for i := 0; i < b.N; i++ {
		table.SetAll(i)
	}
}

func BenchmarkTable_Set(b *testing.B) {
	table := NewTable[int](100, 100)
	for i := 0; i < b.N; i++ {
		table.Set(i%100, i%100, i)
	}
}
