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

func TestSlice(t *testing.T) {
	const (
		EXPECT_OG_LEN     = 5
		EXPECT_SLICED_LEN = 3
	)
	s := []int{1, 2, 3, 4, 5}
	r := jslice.Slice(s, 0, 3)
	if len(s) != EXPECT_OG_LEN {
		t.Fatalf("Expect original slice length to be = %d | Got = %d\n", EXPECT_OG_LEN, len(s))
	}
	if len(r) != EXPECT_SLICED_LEN {
		t.Fatalf("Expect sliced length to be = %d | Got = %d\n", EXPECT_SLICED_LEN, len(r))
	}
	t.Log(s)
	t.Log(r)
}

func TestShift(t *testing.T) {
	const (
		EXPECT_OG_LEN     = 3
		EXPECT_RESULT_VAL = 1
	)
	s := []int{1, 2, 3, 4}
	r := jslice.Shift(&s)
	if len(s) != EXPECT_OG_LEN {
		t.Fatalf("Expected original slice length to now be = %d | Got = %d\n", EXPECT_OG_LEN, len(s))
	}
	if r != EXPECT_RESULT_VAL {
		t.Fatalf("Expected result value to be = %d | Got = %d\n", EXPECT_RESULT_VAL, r)
	}
	t.Log(s)
	t.Log(r)
}

func TestSplice_StartIndex2_Delete0_Insert2(t *testing.T) {
	EXPECT := []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 2, 0, "earth", "mars")

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestSplice_StartIndex0_Delete0_Insert2(t *testing.T) {
	EXPECT := []string{"earth", "mars", "mercury", "venus", "jupiter", "saturn"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 0, 0, "earth", "mars")

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestSplice_StartIndex2_Delete1_Insert2(t *testing.T) {
	EXPECT := []string{"mercury", "venus", "earth", "mars", "saturn"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 2, 1, "earth", "mars")

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestSplice_StartIndex0_Delete1_Insert2(t *testing.T) {
	EXPECT := []string{"earth", "mars", "venus", "jupiter", "saturn"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 0, 1, "earth", "mars")

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestSplice_StartIndex1_DeleteCount0_Insert0(t *testing.T) {
	EXPECT := []string{"mercury"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 1, 3)

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestSplice_DeleteCountGreaterThanSliceLen(t *testing.T) {
	EXPECT := []string{"mercury", "venus"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 2, 100)

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestSplice_StartIndexGreaterThanSliceLen_WithoutReplacementItems(t *testing.T) {
	EXPECT := []string{"mercury", "venus", "jupiter", "saturn"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 100, 100)

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestSplice_StartIndexGreaterThanSliceLen_WithReplacementItems(t *testing.T) {
	EXPECT := []string{"mercury", "venus", "jupiter", "saturn", "earth", "mars"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 100, 1, "earth", "mars")

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestSplice_SpliceLastElement(t *testing.T) {
	EXPECT := []string{"mercury", "venus", "jupiter", "earth", "mars"}
	s := []string{"mercury", "venus", "jupiter", "saturn"}

	jslice.Splice(&s, 3, 1, "earth", "mars")

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	// use jslice to help with testing ;)
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestReverse_1(t *testing.T) {
	EXPECT := []int{5, 4, 3, 2, 1}
	s := []int{1, 2, 3, 4, 5}
	jslice.Reverse(&s)
	if len(EXPECT) != len(s) {
		t.Fatalf("Expected length of = %d | Got %d\n", len(EXPECT), len(s))
	}
	jslice.ForEach(s, func(i int, e int) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestReverse_2(t *testing.T) {
	type foo struct {
		bar int
	}

	EXPECT := []foo{{bar: 5}, {bar: 4}, {bar: 3}, {bar: 2}, {bar: 1}}
	s := []foo{{bar: 1}, {bar: 2}, {bar: 3}, {bar: 4}, {bar: 5}}

	jslice.Reverse(&s)

	if len(EXPECT) != len(s) {
		t.Fatalf("Expected length of = %d | Got %d\n", len(EXPECT), len(s))
	}
	jslice.ForEach(s, func(i int, e foo) {
		if EXPECT[i].bar != e.bar {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestToReversed(t *testing.T) {
	type foo struct {
		bar int
	}

	EXPECT_REVERSED := []foo{{bar: 5}, {bar: 4}, {bar: 3}, {bar: 2}, {bar: 1}}
	EXPECT_OG := []foo{{bar: 1}, {bar: 2}, {bar: 3}, {bar: 4}, {bar: 5}}

	s := []foo{{bar: 1}, {bar: 2}, {bar: 3}, {bar: 4}, {bar: 5}}
	r := jslice.ToReversed(s)

	if len(EXPECT_REVERSED) != len(r) {
		t.Fatalf("Expected length of = %d | Got %d\n", len(EXPECT_REVERSED), len(r))
	}
	jslice.ForEach(r, func(i int, e foo) {
		if EXPECT_REVERSED[i].bar != e.bar {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT_REVERSED, s)
		}
	})
	if len(EXPECT_OG) != len(s) {
		t.Fatalf("Expected length of = %d | Got %d\n", len(EXPECT_OG), len(s))
	}
	jslice.ForEach(s, func(i int, e foo) {
		if EXPECT_OG[i].bar != e.bar {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT_OG, s)
		}
	})
	t.Log(s)
	t.Log(r)
}

func TestToReversed_2(t *testing.T) {
	EXPECT_REVERSED := []string{"baz", "bar", "foo"}
	EXPECT_OG := []string{"foo", "bar", "baz"}

	s := []string{"foo", "bar", "baz"}
	r := jslice.ToReversed(s)

	if len(EXPECT_REVERSED) != len(r) {
		t.Fatalf("Expected length of = %d | Got %d\n", len(EXPECT_REVERSED), len(r))
	}
	jslice.ForEach(r, func(i int, e string) {
		if EXPECT_REVERSED[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT_REVERSED, s)
		}
	})
	if len(EXPECT_OG) != len(s) {
		t.Fatalf("Expected length of = %d | Got %d\n", len(EXPECT_OG), len(s))
	}
	jslice.ForEach(s, func(i int, e string) {
		if EXPECT_OG[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT_OG, s)
		}
	})
	t.Log(s)
	t.Log(r)
}

func TestUnshift(t *testing.T) {
	EXPECT := []int{1, 2, 3, 4, 5}

	s := []int{2, 3, 4, 5}
	jslice.Unshift(&s, 1)

	if len(EXPECT) != len(s) {
		t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
	}
	jslice.ForEach(s, func(i int, e int) {
		if EXPECT[i] != e {
			t.Fatalf("\nExpected\t= %v\nGot\t\t= %v\n", EXPECT, s)
		}
	})
	t.Log(s)
}

func TestAt_NegativeIndex(t *testing.T) {
	EXPECT_VAL := 5
	s := []int{1, 2, 3, 4, 5}
	r := jslice.At(s, -1)
	if r != EXPECT_VAL {
		t.Fatalf("Expected=%d | Got=%d\n", EXPECT_VAL, r)
	}
	t.Log(r)
}

func TestAt_PositiveIndex(t *testing.T) {
	EXPECT_VAL := 2
	s := []int{1, 2, 3, 4, 5}
	r := jslice.At(s, 1)
	if r != EXPECT_VAL {
		t.Fatalf("Expected=%d | Got=%d\n", EXPECT_VAL, r)
	}
	t.Log(r)
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
