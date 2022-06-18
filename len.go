package slices

func Len[T any](slice ...[]T) int {
	length := 0
	for _, item := range slice {
		length += len(item)
	}
	return length
}
