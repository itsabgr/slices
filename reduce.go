package slices

func Reduce[T any, Reduced any](cast func(Reduced, T) Reduced, slice ...T) (r Reduced) {
	for _, item := range slice {
		r = cast(r, item)
	}
	return r
}
