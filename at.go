package jslice

import "math"

// At takes an integer value and returns the item at that index,
// allowing for positive and negative integers. Negative integers
// count back from the last item in the array.
//
// If the provided index is negative but it's absolute value
// is greater than the length of the array, we return the first
// item (index 0) in the array.
func At[T any](s []T, index int) T {
	size := len(s)
	if int(math.Abs(float64(index))) > size {
		return s[0]
	}
	if index < 0 {
		return s[size+index]
	}
	return s[index]
}
