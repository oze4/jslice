package jslice

// Reverse reverses an array by modifying it's contents in-place.
func Reverse[T any](s *[]T) {
	l := 0
	r := len(*s) - 1
	for l < r {
		(*s)[l], (*s)[r] = (*s)[r], (*s)[l]
		l++
		r--
	}
}
