package jslice

// Push appends an element to a slice.
func Push[T any](s *[]T, element T) {
	*s = append(*s, element)
}
