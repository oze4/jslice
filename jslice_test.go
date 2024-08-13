package jslice_test

import (
	"fmt"
	"reflect"
	"strconv"
	"testing"

	"github.com/oze4/jslice"
)

func TestFilter(t *testing.T) {
	const EXPECT = 5
	a := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
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
		t.Fatalf("Expected type []out, but got : %s\n", reflect.TypeOf(r))
	}
	t.Logf("[Input Type = %s | Output Type = %s]\n", reflect.TypeOf(a), reflect.TypeOf(r))
}

func TestForEach(t *testing.T) {
	type foo struct {
		bar string
		baz int
	}

	foos := []*foo{{bar: "a"}, {bar: "b"}, {bar: "c"}, {bar: "d"}, {bar: "e"}}

	jslice.ForEach(foos, func(i int, f *foo) {
		f.baz = i
	})

	str := ""
	jslice.ForEach(foos, func(i int, e *foo) {
		str += fmt.Sprintf("%v, ", e)
	})

	t.Logf("%s\n", str)
}

func TestReduce(t *testing.T) {
	type anum struct {
		num int
	}
	arr := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	res := jslice.Reduce(arr, func(acc []anum, el int, i int, src []int) []anum {
		return append(acc, anum{num: el})
	}, []anum{})

	t.Log(res)
}

func TestPop(t *testing.T) {
	const (
		EXPECT_LEN        = 4
		EXPECT_POPPED_VAL = 5
	)

	arr := []int{1, 2, 3, 4, 5}
	lenBefore := len(arr)
	popped := jslice.Pop(&arr)
	lenAfter := len(arr)

	if popped != EXPECT_POPPED_VAL {
		t.Fatalf("Popped value incorrect. Expected=%d | Got=%d\n", EXPECT_POPPED_VAL, popped)
	}
	if lenAfter != EXPECT_LEN {
		t.Fatalf("Slice length incorrect. Expected=%d | Got=%d\n", EXPECT_LEN, lenAfter)
	}
	t.Logf("Popped=%d | Slice Length=(before=%d,after=%d)\n", popped, lenBefore, lenAfter)
}

func TestPopSliceOfStructs(t *testing.T) {
	const (
		EXPECT_LEN = 6
		EXPECT_VAL = 6
	)

	type foo struct {
		bar int
	}

	foos := []foo{{bar: 0}, {bar: 1}, {bar: 2}, {bar: 3}, {bar: 4}, {bar: 5}, {bar: 6}}
	lenBefore := len(foos)
	popped := jslice.Pop(&foos)
	lenAfter := len(foos)

	if popped.bar != EXPECT_VAL {
		t.Fatalf("Popped element incorrect. Expected=%d | Got=%d\n", EXPECT_VAL, popped.bar)
	}
	if lenAfter != EXPECT_LEN {
		t.Fatalf("Slice length incorrect. Expected=%d | Got=%d\n", EXPECT_LEN, lenAfter)
	}
	t.Logf("Popped=%v | Length=(before=%d,after=%d)\n", popped, lenBefore, lenAfter)
}

func TestPush(t *testing.T) {
	const EXPECT_LEN = 5
	arr := []int{1, 2, 3, 4}
	valToPush := 5
	jslice.Push(&arr, valToPush)
	if len(arr) != EXPECT_LEN {
		t.Fatalf("Slice length incorrect. Expect=%d | Got=%d\n", EXPECT_LEN, len(arr))
	}
	t.Logf("Pushed %d into %v\n", valToPush, arr)
}

func TestSome(t *testing.T) {
	arr := []int{1, 1, 1, 1, 2, 1, 1}
	res := jslice.Some(arr, func(i int, e int) bool {
		return e == 2
	})
	if res == false {
		t.Fatalf("Incorrect result. Expected = true | Got = false\n")
	}
	t.Logf("Some : %t\n", res)
}

func TestEvery(t *testing.T) {
	type Shipment struct {
		Source string
	}

	s := []Shipment{{Source: "New York"}, {Source: "New York"}, {Source: "New York"}}

	r := jslice.Every(s, func(i int, e Shipment) bool {
		return e.Source == "New York"
	})

	if r == false {
		t.Fatalf("Expect true | Got false\n")
	}
	t.Logf("%t\n", r)
}

// **************************************************************

func TestTest(t *testing.T) {
	type Number struct {
		Value int
	}

	s := []int{1, 2, 3}

	reducer := func(acc []Number, e int, i int, og []int) []Number {
		num := Number{Value: e}
		return append(acc, num)
	}

	r := jslice.Reduce(s, reducer, []Number{})

	t.Log(reflect.TypeOf(r))
}
