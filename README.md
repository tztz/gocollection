# gocollection

![Build status](https://github.com/tztz/gocollection/actions/workflows/build.yml/badge.svg)

A Go library for handling sets.

## Set

An API to handle sets providing methods from set theory.

- A `Set` is a collection of unique elements having the same type `T`.
- Values of type `V` can be associated with the elements - but don't have to.
- A `Set` is actually a `map[T]V` with keys of type `T` and values of type `V`, where the values are just associated data.
- If you don't need values, you can omit them in the `Set` to save memory. Then it's just a set of elements (like a set of labels).
- A `Set` can, of course, be empty.
- The zero value of a `Set` is an empty set.

### Example

```go
// First set, add elements
set1 := set.NewWithValues[string, string]()
set1.AddWithValue("apple", "red")
set1.AddWithValue("banana", "yellow")
set1.AddWithValue("cherry", "dark red")
set1.AddWithValue("brick", "red")

// Second set, add elements
set2 := set.NewWithValues[string, string]()
set2.AddWithValue("apple", "green")
set2.AddWithValue("banana", "brownish")
set2.AddWithValue("mango", "green-orange")
set2.AddWithValue("brick", "red")

// Remove element
set2.Remove("brick")

// Calculate intersection
intersectedSet := set1.Intersect(set2)

fmt.Println(intersectedSet.List()) // [banana apple]
```

For more examples have a look at the [example.go](/internal/example/example.go) file.

### Methods

#### Operating on same set

- GetElements
- AddWithValue
- AddWithoutValue
- Remove
- AddAll
- RemoveAll
- Clear

#### Informative

- Size
- List
- String
- StringWithValues

#### Conditional checks

- Contains
- ContainsAny
- Equals
- IsSubset

#### Creating new set

- Copy
- Intersect
- Unite
- UniteDisjunctively
- Subtract
- Filter
- Map

#### Random

- OneR

### Functions

- MapFree
- MapToList
- Reduce
