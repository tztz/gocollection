package main

import (
	"fmt"
	"strings"

	"github.com/tztz/gocollection/pkg/collection/set"
)

func example() {
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

	// Remove element
	set2.Remove("brick")

	// Calculate intersection
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

	// Map to list
	type fruitSpec struct {
		name  string
		color string
	}
	list := set.MapToList(set1, func(elem string, value string) fruitSpec {
		return fruitSpec{name: elem, color: value}
	})

	// One random element
	rndElement, rndValue, _ := set1.OneR()

	fmt.Println(intersectedSet.GetElements())                 // map[apple:green banana:brownish]
	fmt.Println(intersectedSet.Size())                        // 2
	fmt.Println(intersectedSet.Contains("banana"))            // true
	fmt.Println(intersectedSet.List())                        // [apple banana]
	fmt.Println(intersectedSet)                               // banana, apple
	fmt.Println(intersectedSet.StringWithValues())            // banana (brownish), apple (green)
	fmt.Println(set1.Equals(set2))                            // false
	fmt.Println(intersectedSet.IsSubset(set1))                // true
	fmt.Println(filteredSet)                                  // cherry, brick
	fmt.Println(mappedSet)                                    // BANANA, CHERRY, BRICK, APPLE
	fmt.Println(mappedSet.StringWithValues())                 // CHERRY (color: DARK RED), BRICK (color: RED), APPLE (color: RED), BANANA (color: YELLOW)
	fmt.Println(freelyMappedSet)                              // map[{dark red CHERRY}:6 {red APPLE}:5 {red BRICK}:5 {yellow BANANA}:6]
	fmt.Println(list)                                         // [{apple red} {banana yellow} {cherry dark red} {brick red}]
	fmt.Printf("elem: %v, value: %v\n", rndElement, rndValue) // elem: banana, value: yellow

	// Clear set
	intersectedSet.Clear()

	fmt.Println(intersectedSet.Size()) // 0
}

func main() {
	example()
}
