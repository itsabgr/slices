package slices

func Len[T ~[]E, E any](slice ...[]T) int {
	length := 0
	for _, item := range slice {
		length += len(item)
	}
	return length
}
