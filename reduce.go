package jslice

// Reduce reduce a slice
func Reduce[I any, O any](source []I, reducer func(accumulator O, currentElement I, currentIndex int, originalSlice []I) O, initialValue O) O {
	prevElement := initialValue
	out := initialValue

	for i := range source {
		out = reducer(prevElement, source[i], i, source)
		prevElement = out
	}

	return out
}
