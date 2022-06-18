package slices

func Count[T any](counter func(T) int, slice ...T) int {
	t := 0
	for _, item := range slice {
		t += counter(item)
	}
	return t
}
