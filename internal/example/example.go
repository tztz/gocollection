package main

import (
	"fmt"
	"strings"

	"github.com/tztz/gocollection/pkg/collection/set"
)

func example() []string {
	// First set
	set1 := set.NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")
	set1.AddWithValue("brick", "red")

	// Second set
	set2 := set.NewWithValues[string, string]()
	set2.AddWithValue("apple", "green")
	set2.AddWithValue("banana", "brownish")
	set2.AddWithValue("mango", "green-orange")
	set2.AddWithValue("brick", "red")

	// Remove element from set
	set2.Remove("brick")

	// Calculate intersection of two sets
	intersectedSet := set1.Intersect(set2)

	// Filter set
	filteredSet := set1.Filter(func(elem string, value string) bool {
		return strings.Contains(elem, "c")
	})

	// Map set
	mappedSet := set1.Map(func(elem string, value string) (string, string) {
		return strings.ToUpper(elem), fmt.Sprintf("color: %v", strings.ToUpper(value))
	})

	// Freely map set
	type fruit struct {
		value string
	}
	freelyMappedSet := set.MapFree(set1, func(elem string, value string) (fruit, int) {
		newElem := fruit{value: value + " " + strings.ToUpper(elem)}
		newValue := len(elem)
		return newElem, newValue
	})

	// Map set to list
	type fruitSpec struct {
		name  string
		color string
	}
	list := set.MapToList(set1, func(elem string, value string) fruitSpec {
		return fruitSpec{name: elem, color: value}
	})

	// Reduce set
	reducedValue := set.Reduce(set1, func(elem string, value string, acc int) int {
		return acc + len(elem)
	}, 0)

	// Get one random element from set
	rndElement, rndValue, _ := set1.OneR()

	results := make([]string, 0)

	results = append(results, fmt.Sprintln(intersectedSet.GetElements()))                 // map[apple:green banana:brownish]
	results = append(results, fmt.Sprintln(intersectedSet.Size()))                        // 2
	results = append(results, fmt.Sprintln(intersectedSet.Contains("banana")))            // true
	results = append(results, fmt.Sprintln(intersectedSet.List()))                        // [apple banana]
	results = append(results, fmt.Sprintln(intersectedSet))                               // banana, apple
	results = append(results, fmt.Sprintln(intersectedSet.StringWithValues()))            // banana (brownish), apple (green)
	results = append(results, fmt.Sprintln(set1.Equals(set2)))                            // false
	results = append(results, fmt.Sprintln(intersectedSet.IsSubset(set1)))                // true
	results = append(results, fmt.Sprintln(filteredSet))                                  // cherry, brick
	results = append(results, fmt.Sprintln(mappedSet))                                    // BANANA, CHERRY, BRICK, APPLE
	results = append(results, fmt.Sprintln(mappedSet.StringWithValues()))                 // CHERRY (color: DARK RED), BRICK (color: RED), APPLE (color: RED), BANANA (color: YELLOW)
	results = append(results, fmt.Sprintln(freelyMappedSet))                              // map[{dark red CHERRY}:6 {red APPLE}:5 {red BRICK}:5 {yellow BANANA}:6]
	results = append(results, fmt.Sprintln(list))                                         // [{apple red} {banana yellow} {cherry dark red} {brick red}]
	results = append(results, fmt.Sprintln(reducedValue))                                 // 22
	results = append(results, fmt.Sprintf("elem: %v, value: %v\n", rndElement, rndValue)) // elem: banana, value: yellow

	// Clear set
	intersectedSet.Clear()

	results = append(results, fmt.Sprintln(intersectedSet.Size())) // 0

	return results
}

func main() {
	results := example()
	fmt.Println(strings.Join(results, ""))
}
