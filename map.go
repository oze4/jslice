package jslice

// Map iterates over a slice, returning a new slice.
func Map[I any, O any](s []I, fn func(int, I) O) []O {
	o := []O{}
	for i, e := range s {
		o = append(o, fn(i, e))
	}
	return o
}
