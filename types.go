package slices

type Condition[T any] func(a T) bool
type Entry[K any, V any] struct {
	Key K
	Val V
}
