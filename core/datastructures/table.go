package advent_of_code

// A table is a 2D array / grid type of structure.
type Table[T any] struct {
	grid [][]T
}

// NewTable creates a new Table with the specified number of rows and columns
func NewTable[T any](rows, cols int) *Table[T] {
	grid := make([][]T, rows)
	for i := range grid {
		grid[i] = make([]T, cols)
	}
	return &Table[T]{grid: grid}
}

func (table *Table[T]) SetAll(value T) {
	for i := range table.grid {
		for j := range table.grid[i] {
			table.grid[i][j] = value
		}
	}
}

// GetRowCount returns the number of rows in the table
func (table *Table[T]) GetRowCount() int {
	return len(table.grid)
}

// GetColumnCount returns the number of columns in the table
func (table *Table[T]) GetColumnCount() int {
	if len(table.grid) == 0 {
		return 0
	}
	return len(table.grid[0])
}

// Set sets the value at the specified row and column
func (table *Table[T]) Set(row, col int, value T) {
	table.grid[row][col] = value
}

// Get retrieves the value at the specified row and column
func (table *Table[T]) Get(row, col int) T {
	return table.grid[row][col]
}
