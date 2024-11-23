# gocollection

![Build status](https://github.com/tztz/gocollection/actions/workflows/build.yml/badge.svg)

## Set

An API to handle sets providing methods from set theory.

- A `Set` is a collection of unique elements having the same type `T`.
- Values of type `V` can be associated with the elements - but don't have to.
- A `Set` is actually a `map[T]V` with keys of type `T` and values of type `V`, where the values are just associated data.
- If you don't need values, you can omit them in the `Set` to save memory. Then it's just a set of elements like a set of labels.
- A `Set` can, of course, be empty.
- The zero value of a `Set` is an empty set.

Example

```go
// First set
set1 := NewWithValues[string, string]()
set1.AddWithValue("apple", "red")
set1.AddWithValue("banana", "yellow")
set1.AddWithValue("cherry", "dark red")

// Second set
set2 := NewWithValues[string, string]()
set2.AddWithValue("apple", "green")
set2.AddWithValue("banana", "brownish")
set2.AddWithValue("mango", "green-orange")

// Calculate intersection
intersectedSet := set1.Intersect(set2)

// Result
fmt.Println(intersectedSet.Size()) // 2
fmt.Println(intersectedSet.List()) // [banana apple]
```

Todo: more doc to come ...
