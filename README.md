# gocollection

![Build status](https://github.com/tztz/gocollection/actions/workflows/build.yml/badge.svg)

An API to handle sets providing methods from Set theory.

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
fmt.Println(intersectedSet.List()) // [apple banana]
```

Todo: more doc to come ...
