package jslice

// Pop removes an item from the end of a slice and returns it.
func Pop[T any](s *[]T) T {
	size := len(*s)
	last := (*s)[size-1]
	*s = (*s)[:size-1]
	return last
}
