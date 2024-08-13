package jslice

// Every returns true/false based on every element in slice
// meeting provided condition.
func Every[T any](s []T, fn func(int, T) bool) bool {
    for i, e := range s {
        if r := fn(i, e); r == false {
            return false
        }
    }
    return true
}