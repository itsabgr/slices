package slices

func Clone[T any](slice ...T) []T {
	clone := make([]T, len(slice))
	for i, item := range slice {
		clone[i] = item
	}
	return clone
}
