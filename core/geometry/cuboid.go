package advent_of_code

import "golang.org/x/exp/constraints"

type Cuboid[T constraints.Signed] struct {
	width  T
	length T
	height T
}

func NewCuboid[T constraints.Signed](length T, width T, height T) *Cuboid[T] {
	return &Cuboid[T]{
		length: length,
		width:  width,
		height: height,
	}
}

func (cuboid *Cuboid[T]) GetSurfaceArea() T {
	return (2 * cuboid.length * cuboid.width) + (2 * cuboid.width * cuboid.height) + (2 * cuboid.height * cuboid.length)
}

func (cuboid *Cuboid[T]) GetAreaOfSmallestSide() T {
	side1 := cuboid.length * cuboid.width
	side2 := cuboid.width * cuboid.height
	side3 := cuboid.height * cuboid.length

	if side1 < side2 && side1 < side3 {
		return side1
	} else if side2 < side3 {
		return side2
	} else {
		return side3
	}
}

func (cuboid *Cuboid[T]) GetSmallestPerimeter() T {
	// Calculate the perimeters of each of the three faces
	perimeter1 := 2 * (cuboid.length + cuboid.width)  // Perimeter of face with sides Length and Width
	perimeter2 := 2 * (cuboid.width + cuboid.height)  // Perimeter of face with sides Width and Height
	perimeter3 := 2 * (cuboid.height + cuboid.length) // Perimeter of face with sides Height and Length

	// Return the smallest of the three perimeters
	return min(perimeter1, min(perimeter2, perimeter3))
}

func (cuboid *Cuboid[T]) GetVolume() T {
	return cuboid.height * cuboid.width * cuboid.length
}

func (cuboid *Cuboid[T]) IsCube() bool {
	return cuboid.length == cuboid.height && cuboid.length == cuboid.width && cuboid.width == cuboid.height
}

func (this *Cuboid[T]) IsEqual(other *Cuboid[T]) bool {
	return this.length == other.length && this.width == other.width && this.height == other.height
}
