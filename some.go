package jslice

// Some tests that at least one element in the slice matches provided testFn.
func Some[T any](s []T, testFn func(i int, e T) bool) bool {
	for i, el := range s {
		res := testFn(i, el)
		if res {
			return true
		}
	}
	return false
}
