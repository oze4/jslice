package jslice_test

import (
	"reflect"
	"strconv"
	"testing"

	"github.com/oze4/jslice"
)

func TestFilter(t *testing.T) {
	const EXPECT = 5
	a := []int{0,1,2,3,4,5,6,7,8,9}
	r := jslice.Filter(a, func(i int, e int) bool {
		return e >= 5
	})
	t.Logf("[Filtered Slice = %v | Filtered Slice Length = %d]\n", r, len(r))
	if len(r) != 5 {
		t.Fatalf("Expected = %d | Got = %d\n", EXPECT, len(r))
	}
}

func TestMap(t *testing.T) {
	type in struct {
		bar int
	}
	type out struct {
		baz string
	}

	a := []in{{bar: 0}, {bar: 1}, {bar: 2}, {bar: 3}, {bar: 4}, {bar: 5}}
	r := jslice.Map(a, func(i int, e in) out {
		barStr := strconv.Itoa(e.bar)
		return out{baz: barStr}
	})
	if reflect.TypeOf(r) != reflect.TypeOf([]out{}) {
		t.Fatalf("Expected type []out, but got : %s\n", reflect.TypeOf(r) )
	}
	t.Logf("[Input Type = %s | Output Type = %s]\n", reflect.TypeOf(a), reflect.TypeOf(r))
}
