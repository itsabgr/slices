package slices

func Keys[K comparable, V any](_map map[K]V) []K {
	slice := make([]K, 0, len(_map))
	for k, _ := range _map {
		slice = append(slice, k)
	}
	return slice
}
