package slices

func FromChan[T ~chan E, E any](c T) []E {
	if c == nil {
		return nil
	}
	var slice []E
	for item := range c {
		slice = append(slice, item)
	}
	return slice
}
