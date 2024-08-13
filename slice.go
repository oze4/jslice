package jslice

// Slice returns a copy of a portion of a slice.
//
// **Note**: `end` is not included in result!
// We select elements from `start` up to, but not including, `end`.
//
// If you provide an `end` value that is greater than the length
// of the provided slice, we default `end` to that slices length.
func Slice[T any](s []T, start, end uint) []T {
	if int(end) > len(s) {
		end = uint(len(s))
	}
	return s[start:end]
}
