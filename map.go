package jslice

// Map iterates over a slice, returning a new slice.
func Map[S ~[]I, I comparable, O any](s S, fn func(int, I) O) []O {
	o := []O{}
	for i, e := range s {
		r := fn(i, e)
		o = append(o, r)
	}
	return o
}
