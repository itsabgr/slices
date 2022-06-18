package slices

func Fill[T ~[]E, E any](item T, slice ...T) []T {
	for i := range slice {
		slice[i] = item
	}
	return slice
}
