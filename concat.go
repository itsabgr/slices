package slices

func Concat[T ~[]E, E any](slice ...[]T) []T {
	length := Len(slice)
	if length == 0 {
		return nil
	}
	result := make([]T, 0, length)
	for _, item := range slice {
		result = append(result, item...)
	}
	return result
}
