package slices

func Sort[T any](fn func(a, b T) bool, array ...T) []T {
	var n = len(array)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if fn(array[j-1], array[j]) {
				array[j-1], array[j] = array[j], array[j-1]
			}
			j = j - 1
		}
	}
	return array
}
