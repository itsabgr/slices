package slices

func Filter[T any](condition Condition[T], slice ...T) []T {
	result := make([]T, 0, len(slice))
	for _, item := range slice {
		if condition(item) {
			result = append(result, item)
		}
	}
	return result
}
