package slices

func ToChan[T ~[]E, E any](slice ...T) <-chan T {
	c := make(chan T, len(slice))
	for _, item := range slice {
		c <- item
	}
	return c
}
