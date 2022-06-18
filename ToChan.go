package slices

func ToChan[T any](slice ...T) <-chan T {
	c := make(chan T, len(slice))
	for _, item := range slice {
		c <- item
	}
	return c
}
