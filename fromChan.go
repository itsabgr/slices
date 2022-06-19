package slices

func FromChan[T any](c chan T) []T {
	if c == nil {
		return nil
	}
	var slice []T
	for {
		item, ok := <-c
		if !ok {
			break
		}
		slice = append(slice, item)
	}
	return slice
}
