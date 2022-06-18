package slices

func Map[From any, To any](cast func(From) To, slice ...From) []To {
	v := make([]To, len(slice))
	for i, item := range slice {
		v[i] = cast(item)
	}
	return v
}
