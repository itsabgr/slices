package slices

func Group[T any, G comparable](fn func(T) G, slice ...T) map[G][]T {
	_map := make(map[G][]T)
	for _, item := range slice {
		g := fn(item)
		m, _ := _map[g]
		m = append(m, item)
		_map[g] = m
	}
	return _map
}
