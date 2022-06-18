package slices

func At[T any](index int, slice ...T) (t T) {
	if len(slice) > index {
		t = slice[index]
	}
	return t
}
