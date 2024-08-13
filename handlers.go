package jslice

// FilterHandler handles filter operations.
type FilterHandler[T comparable] func(int, T) bool

// MapHandler handles map operations.
type MapHandler[T comparable, O any] func(int, T) O
