package slices

func ToMap[T ~[]E, E any](slice ...T) map[int]T {
	_map := make(map[int]T, len(slice))
	for i, item := range slice {
		_map[i] = item
	}
	return _map
}
