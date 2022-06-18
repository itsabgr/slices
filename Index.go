package slices

func Index[T any](condition Condition[T], slice ...T) int {
	for i, item := range slice {
		if condition(item) {
			return i
		}
	}
	return -1
}
