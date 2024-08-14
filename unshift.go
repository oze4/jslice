package jslice

// Unshift adds an element to the front of a slice.
func Unshift[T any](s *[]T, e T) {
	*s = append([]T{e}, *s...)
}
