package main

import (
	"fmt"
	"strings"

	"github.com/tztz/gocollection/collection/set"
)

func main() {
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

	// One random element
	rndElement, rndValue, _ := set1.OneR()

	fmt.Println(intersectedSet.GetElements())                 // map[apple:green banana:brownish]
	fmt.Println(intersectedSet.Size())                        // 2
	fmt.Println(intersectedSet.Contains("banana"))            // true
	fmt.Println(intersectedSet.List())                        // [apple banana]
	fmt.Println(intersectedSet)                               // apple, banana
	fmt.Println(intersectedSet.StringWithValues())            // apple (green), banana (brownish)
	fmt.Println(set1.Equals(set2))                            // false
	fmt.Println(intersectedSet.IsSubset(set1))                // true
	fmt.Println(filteredSet)                                  // cherry, brick
	fmt.Println(mappedSet)                                    // APPLE, BANANA, CHERRY, BRICK
	fmt.Println(mappedSet.StringWithValues())                 // APPLE (color: RED), BANANA (color: YELLOW), CHERRY (color: DARK RED), BRICK (color: RED)
	fmt.Printf("elem: %v, value: %v\n", rndElement, rndValue) // elem: banana, value: yellow

	// Clear set
	intersectedSet.Clear()

	fmt.Println(intersectedSet.Size()) // 0
}