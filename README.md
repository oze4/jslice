# jslice

`jslice`, pronounced [JS-slice] provides generic, JavaScript-like, [array methods](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Array#instance_methods) for Go slices.

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

Appends an element to the end of a slice.

```go
s := []int{1,2,3}
i := 4
jslice.Push(&s, i)
// s == []int{1,2,3,4}
```

## Pop

Removes element from end of slice and returns the element that was removed.

```go
s := []int{1,2,3}
item := jslice.Pop(&s)
// s == []int{1,2}
// item == 3
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


<br />
<br />
<br />
<br />

------

<footer>
<small>matt oestreich</small>
</footer>

----