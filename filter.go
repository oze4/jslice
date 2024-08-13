package jslice

// Filter filters a slice based upon conditions defined in FilterHandler
func Filter[S ~[]E, E comparable](s S, fn func(int, E) bool) S {
	r := S{}
	for i, e := range s {
		if fn(i, e) {
			r = append(r, e)
		}
	}
	return r
}
