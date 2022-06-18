package slices

import "math/rand"

func Shuffle[T ~[]E, E any](slice ...T) []T {
	rand.Shuffle(len(slice), func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})
	return slice
}
