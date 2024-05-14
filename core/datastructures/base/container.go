package containers

type Container interface {
	IsEmpty() bool
	Size() int
	Clear()
	ToString()
}
