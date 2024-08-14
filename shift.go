package jslice

// Shift removes the first element from a slice and returns that element.
// Shift changes the lenght of an array.
func Shift[T any](s *[]T) T {
	first := (*s)[0]
	*s = (*s)[1:]
	return first
}
