package jslice

// Some tests that at least one element in the slice matches provided testFn.
func Some[T any](s []T, fn func(i int, e T) bool) bool {
	for i, el := range s {
		if fn(i, el) {
			return true
		}
	}
	return false
}
