package jslice

// Map iterates over a slice, returning a new slice.
func Map[S ~[]E, E comparable, O any](s S, fn MapHandler[E, O]) []O {
	o := []O{}
	for i, e := range s {
		r := fn(i, e)
		o = append(o, r)
	}
	return o
}
