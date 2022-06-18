package slices

func FromMap[K comparable, V any](_map map[K]V) []Entry[K, V] {
	entries := make([]Entry[K, V], 0, len(_map))
	for k, v := range _map {
		entries = append(entries, Entry[K, V]{k, v})
	}
	return entries
}
