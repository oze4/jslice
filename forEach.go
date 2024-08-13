package jslice

// ForEach performs an action for each element in a slice.
// Note: ForEach returns void.
func ForEach[T any](s []T, fn func(int, T)) {
	for i, e := range s {
		fn(i, e)
	}
}
