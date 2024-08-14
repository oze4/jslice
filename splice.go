package jslice

// Splice changes the contents of a slice by removing or replacing existing elements,
// and/or adding new elements.
func Splice[T any](s *[]T, start uint, deleteCount uint, replacementItems ...T) {
	if deleteCount == 0 && len(replacementItems) == 0 {
		return
	}

	// If start >= len(*s) no elements will be deleted, but the method will behave as
	// an adding function.
	if start >= uint(len(*s)) {
		if len(replacementItems) == 0 {
			return
		}
		*s = append(*s, replacementItems...)
		return
	}

	// If the "end" (start+deleteCount) is greater than the length of the slice, limit
	// the delete count to the length of the slice - start. Otherwise we get index out
	// of bounds error.
	if start+deleteCount >= uint(len(*s)) {
		deleteCount = uint(len(*s)) - start
	}

	*s = append((*s)[0:start], append(replacementItems, (*s)[start+deleteCount:]...)...)
	return
}
