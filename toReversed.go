package jslice

// ToReversed reverses an array and returns a copy. It does not modify
// the original slice.
func ToReversed[T any](s []T) []T {
	size := len(s)
	out := make([]T, size)

	for i, v := range s {
		out[size-1-i] = v
	}

	return out
}
