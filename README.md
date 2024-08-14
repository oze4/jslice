# jslice

`jslice`, pronounced JS-slice (jay-es-slice), provides generic, JavaScript-like [array methods](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array#instance_methods) for Go slices.

## Filter

Filters an array.

```go
s := []int{1,1,1,9,9,9}

r := jslice.Filter(s, func(index int, element int) bool {
  return element > 5
})
// r == []int{9,9,9}
```

```go
type TestResult struct {
  Result string 
}

s := []TestResult{{Result: "Fail", {Result: "Pass"}, {Result: "In-progress"}}}

fails := jslice.Filter(s, func(i int, e TestResult) bool {
  return e.Result == "Fail"
})
// fails == []TestResult{{Result: "Fail"}}
```

## Map

Maps over an array and returns a new array.

```go
type In struct {
  Foo int
}
type Out struct {
  Bar string
}

s := []In{{Foo: 1}, {Foo: 2}, {Foo: 3}, {Foo: 4}}

r := jslice.Map(s, func(i int, e In) Out {
  return Out{Bar: strconv.Itoa(e.Foo)}
})
// r == []Out{{Bar: "1"}, {Bar: "2"}, {Bar: "3"}, {Bar: "4"}}
```

## Reduce

Executes the provided reducer function and stores the result in an accumulator, which is returned at the time of completion.

```go
s := []int{54,43,32,21,10,9}

reducer := func(acc int, currEl int, currIdx int, originalSlice []int) int {
  return acc + currEl
}

sum := jslice.Reduce(s, reducer, 0)
// sum == 169
```

```go
type Number struct {
  Value int
}

s := []int{1,2,3}

reducer := func(acc []Number, e int, i int, og []int) []Number {
  num := Number{Value: e}
  return append(acc, num)
}

r := jslice.Reduce(s, reducer, []Number{})
// r == []Number{{Value: 1}, {Value: 2}, {Value: 3}}
```

## ForEach

Iterates over a slice calling the provided function on each element. Note: `ForEach` does not return anything.

```go
s := []int{1,2,3}

jslice.ForEach(s, func(i int, e int) {
  fmt.Printf("Element '%d' is at index '%d'.\n", i, e)
})
```

## Push

Appends an element to the end of a slice. `Push` modifies the original slice.

```go
s := []int{1,2,3}
i := 4
jslice.Push(&s, i)
// s == []int{1,2,3,4}
```

## Pop

Removes element from end of slice and returns the element that was removed. `Pop` modifies the original slice.

```go
s := []int{1,2,3}
item := jslice.Pop(&s)
// s == []int{1,2}
// item == 3
```

## Shift

Removes first element (index 0) from slice and returns the removed element. `Shift` modifies the orignal slice.

```go
s := []int{1,2,3,4}
i := jslice.Shift(&s)
// s == []int{2,3,4}
// i == 1
```

## Some

Tests that at least one element in the slice matches provided condition.

```go
s := []int{1,1,2,1,1}

r := jslice.Some(s, func(i int, e int) bool {
  return e == 2
})
// r == true
```

## Every

Tests that every element in slice meets provided condition. Returns false if at least one element does not meet condition.

```go
type Shipment struct {
  Source string
}

s := []Shipment{{Source: "New York"}, {Source: "New York"}, {Source: "New York"}}

r := jslice.Every(s, func(i int, e Shipment) bool {
  return e.Source == "New York"
})
// r == true
```

## Slice

Returns a copy of a portion of a slice. We return a slice from `start` up to, but not including, `end`. If `end > len(slice)` we default to `end = len(slice)`. 

```go
s := []int{1,2,3,4,5}
start := 0
end := 3
r := jslice.Slice(s, start, end)
// r == []int{1,2,3}
```

## Splice

Changes the contents of a slice by removing or replacing existing elements and/or adding new elements.

- If `deleteCount` and `replacementElements` both equal `0`, we just return the original slice without modifying anything.
- If `start` is greater than or equal to the length of the slice, no elements will be deleted, but the method will behave as an adding function.

<b>Remove `0` elements before index `2` and insert "`earth`" and "`mars`"</b>
```go
s := []string{"mercury", "venus", "jupiter", "saturn"}
jslice.Splice(&s, 2, 0, "earth", "mars")
// s == []string{"mercury", "venus", "earth", "mars", "jupiter", "saturn"}
```

**Remove `0` elements at index `0` and insert "`earth`" and "`mars`".**

```go
s := []string{"mercury", "venus", "jupiter", "saturn"}
jslice.Splice(&s, 0, 0, "earth", "mars")
// s == []string{"earth", "mars", "mercury", "venus", "jupiter", "saturn"}
```

**Remove `1` element at index `2`, and insert "`earth`" and "`mars`"**

```go
s := []string{"mercury", "venus", "jupiter", "saturn"}
jslice.Splice(&s, 2, 1, "earth", "mars")
// s == []string{"mercury", "venus", "earth", "mars", "saturn"}
```

**Remove `1` element at index `0` and insert "`earth`" and "`mars`"**

```go
s := []string{"mercury", "venus", "jupiter", "saturn"}
jslice.Splice(&s, 0, 1, "earth", "mars")
// s == []string{"earth", "mars", "venus", "jupiter", "saturn"}
```

**Remove `3` elements starting at index `1` and insert nothing**

```go
s := []string{"mercury", "venus", "jupiter", "saturn"}
jslice.Splice(&s, 1, 3)
// s == []string{"mercury"}
```

**If `start` + `deleteCount` is greater than or equal to slice length, we modify `deleteCount` to equal the length of the slice - `start`**

```go
s := []string{"mercury", "venus", "jupiter", "saturn"}
jslice.Splice(&s, 2, 100) // <- 100 greater than slice length
// s == []string{"mercury", "venus"}
```

**If `start` is greater than or equal to the length of the slice, no elements are removed, but the method is treated as an add function**

```go
s := []string{"mercury", "venus", "jupiter", "saturn"}
// Even though `deleteCount` == 1, nothing will be
// deleted because `start` >= length of slice.
jslice.Splice(&s, 100, 1, "earth", "mars") 
// s == []string{"mercury", "venus", "jupiter", "saturn", "earth", "mars"}
```

**Splice last element by removing `1` element at index `3` and inserting "`earth`" and "`mars`"**

```go
s := []string{"mercury", "venus", "jupiter", "saturn"}
jslice.Splice(&s, 3, 1, "earth", "mars")
// s == []string{"mercury", "venus", "jupiter", "earth", "mars"}
```

## Reverse

Modifies a slice in-place by reversing it's elements. If you do not want to modify the slice in-place, use `ToReversed` method.

```go
s := []int{1,2,3,4,5}
jslice.Reverse(&s)
// s == []int{5,4,3,2,1}
```

## ToReversed

Returns a copy of a slice in reversed order.

```go
s := []string{"foo", "bar", "baz"}
r := jslice.ToReversed(s)
// s == []string{"foo", "bar", "baz"}
// r == []string{"baz", "bar", "foo"}
```

## Unshift

Adds an element to the front of a slice.

```go
s := []int{2,3,4,5}
jslice.Unshift(&s, 1)
// s == []int{1,2,3,4,5}
```




<br />
<br />
<br />
<br />

------

<footer>
<small>matt oestreich</small>
</footer>

----