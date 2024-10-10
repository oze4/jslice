package jslice

// Reduce reduces a slice
func Reduce[InputType any, OutputType any](source []InputType, reducer func(acc OutputType, currEl InputType, currIdx int, ogSlice []InputType) OutputType, initVal OutputType) OutputType {
	prevElement := initVal
	out := initVal

	for i := range source {
		out = reducer(prevElement, source[i], i, source)
		prevElement = out
	}

	return out
}
