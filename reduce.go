package jslice

// Reduce reduce a slice
func Reduce[I any, O any](source []I, reducer func(prevValue O, currValue I, index int, list []I) O, initial O) O {
	prevElement := initial
	out := initial

	for i := range source {
		out = reducer(prevElement, source[i], i, source)
		prevElement = out
	}

	return out
}
