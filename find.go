package slices

func Find[T ~[]E, E any](condition Condition[T], slice ...T) (index int, t T) {
	for i, item := range slice {
		if condition(item) {
			return i, item
		}
	}
	return -1, t
}
