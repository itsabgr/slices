package slices

func Values[K comparable, V any](_map map[K]V) []V {
	slice := make([]V, 0, len(_map))
	for _, v := range _map {
		slice = append(slice, v)
	}
	return slice
}
