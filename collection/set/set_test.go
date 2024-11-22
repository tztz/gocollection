package set

import (
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroValueOfSetIsAnEmptySet(t *testing.T) {
	// When
	var set Set[string, string]

	// Then
	assert.Equal(t, 0, set.Size())
	assert.Equal(t, []string{}, set.List())
}

func TestShouldCreateEmptySet(t *testing.T) {
	// When
	set1 := NewWithoutValues[string]()

	// Then
	assert.Equal(t, 0, set1.Size())
	assert.False(t, set1.Contains("does not exist"))
	assert.Equal(t, []string{}, set1.List())

	// When
	set2 := NewWithValues[string, string]()

	// Then
	assert.Equal(t, 0, set2.Size())
	assert.False(t, set2.Contains("does not exist"))
	assert.Equal(t, []string{}, set2.List())
}

func TestShouldAddItemWithoutValueToSet(t *testing.T) {
	// Given
	set := NewWithoutValues[string]()

	// Expect
	assert.Equal(t, 0, set.Size())

	// When
	set.AddWithoutValue("apple")
	// Then
	assert.Equal(t, 1, set.Size())
	assert.True(t, set.Contains("apple"))

	// When
	set.AddWithoutValue("banana")
	// Then
	assert.Equal(t, 2, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
}

func TestShouldAddItemWithValueToSet(t *testing.T) {
	// Given
	set := NewWithValues[string, string]()

	// Expect
	assert.Equal(t, 0, set.Size())

	// When
	set.AddWithValue("apple", "red")
	// Then
	assert.Equal(t, 1, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.Equal(t, "red", set.elements["apple"])

	// When
	set.AddWithValue("banana", "yellow")
	// Then
	assert.Equal(t, 2, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
	assert.Equal(t, "red", set.elements["apple"])
	assert.Equal(t, "yellow", set.elements["banana"])
}

func TestSetWithoutValuesShouldStayUnique(t *testing.T) {
	// Given
	set := NewWithoutValues[string]()

	// Expect
	assert.Equal(t, 0, set.Size())

	// When
	set.AddWithoutValue("apple")
	set.AddWithoutValue("apple")
	set.AddWithoutValue("apple")
	set.AddWithoutValue("banana")
	set.AddWithoutValue("banana")
	set.AddWithoutValue("banana")

	// Then
	assert.Equal(t, 2, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
}

func TestSetWithValuesShouldStayUnique(t *testing.T) {
	// Given
	set := NewWithValues[string, string]()

	// Expect
	assert.Equal(t, 0, set.Size())

	// When
	set.AddWithValue("apple", "red")
	set.AddWithValue("apple", "red")
	set.AddWithValue("banana", "yellow")
	set.AddWithValue("banana", "yellow")

	// Then
	assert.Equal(t, 2, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
	assert.Equal(t, "red", set.elements["apple"])
	assert.Equal(t, "yellow", set.elements["banana"])

	// When
	set.AddWithValue("apple", "light red")
	set.AddWithValue("banana", "dark yellow")

	// Then
	assert.Equal(t, "light red", set.elements["apple"])
	assert.Equal(t, "dark yellow", set.elements["banana"])
}

func TestShouldRemoveItemFromSet(t *testing.T) {
	// Given
	set := NewWithoutValues[string]()
	set.AddWithoutValue("apple")
	set.AddWithoutValue("banana")

	// Expect
	assert.Equal(t, 2, set.Size())

	// When
	set.Remove("apple")
	// Then
	assert.Equal(t, 1, set.Size())
	assert.False(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))

	// When
	set.Remove("banana")
	// Then
	assert.Equal(t, 0, set.Size())
	assert.False(t, set.Contains("apple"))
	assert.False(t, set.Contains("banana"))
}

func TestShouldListAllSetItems(t *testing.T) {
	// Given
	fruitSet := NewWithoutValues[string]()
	fruitSet.AddWithoutValue("apple")
	fruitSet.AddWithoutValue("banana")
	fruitSet.AddWithoutValue("cherry")
	// When
	items1 := fruitSet.List()
	// Then
	assert.ElementsMatch(t, []string{"apple", "banana", "cherry"}, items1)
	assert.Equal(t, 3, fruitSet.Size())

	// Given
	emptySet := NewWithoutValues[string]()
	// When
	items2 := emptySet.List()
	// Then
	assert.ElementsMatch(t, []string{}, items2)
	assert.Equal(t, 0, emptySet.Size())
}

func TestShouldClearSet(t *testing.T) {
	// Given
	set := NewWithoutValues[string]()
	set.AddWithoutValue("apple")
	set.AddWithoutValue("banana")
	set.AddWithoutValue("cherry")

	// When
	set.Clear()

	// Then
	assert.Equal(t, 0, set.Size())
	assert.Equal(t, []string{}, set.List())
	assert.False(t, set.Contains("apple"))
	assert.False(t, set.Contains("banana"))
	assert.False(t, set.Contains("cherry"))
}

func TestShouldGetCanonicalStringRepresentationOfSetWithoutValues(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")
	// When
	str1 := set1.String()
	// Then
	assert.Equal(t, 21, len(str1))
	splitStrings := strings.Split(str1, ", ")
	slices.Sort(splitStrings)
	assert.Equal(t, "apple, banana, cherry", strings.Join(splitStrings, ", "))

	// Given
	set2 := NewWithoutValues[string]()
	// When
	str2 := set2.String()
	// Then
	assert.Equal(t, "", str2)
}

func TestShouldGetCanonicalStringRepresentationOfSetWithValues(t *testing.T) {
	// Given
	set1 := NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")
	// When
	str1 := set1.String()
	// Then
	assert.Equal(t, 21, len(str1))
	splitStrings := strings.Split(str1, ", ")
	slices.Sort(splitStrings)
	assert.Equal(t, "apple, banana, cherry", strings.Join(splitStrings, ", "))

	// Given
	set2 := NewWithValues[string, string]()
	// When
	str2 := set2.String()
	// Then
	assert.Equal(t, "", str2)
}

func TestStringWithValuesShouldGetStringRepresentationOfSetWithoutValues(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")
	// When
	str1 := set1.StringWithValues()
	// Then
	assert.Equal(t, 36, len(str1))
	splitStrings := strings.Split(str1, ", ")
	slices.Sort(splitStrings)
	assert.Equal(t, "apple ({}), banana ({}), cherry ({})", strings.Join(splitStrings, ", "))

	// Given
	set2 := NewWithValues[string, string]()
	// When
	str2 := set2.StringWithValues()
	// Then
	assert.Equal(t, "", str2)
}

func TestStringWithValuesShouldGetStringRepresentationOfSetWithValues(t *testing.T) {
	// Given
	set1 := NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")
	// When
	str1 := set1.StringWithValues()
	// Then
	assert.Equal(t, 47, len(str1))
	splitStrings := strings.Split(str1, ", ")
	slices.Sort(splitStrings)
	assert.Equal(t, "apple (red), banana (yellow), cherry (dark red)", strings.Join(splitStrings, ", "))

	// Given
	set2 := NewWithValues[string, string]()
	// When
	str2 := set2.StringWithValues()
	// Then
	assert.Equal(t, "", str2)
}

func TestTwoSetsWithoutValuesAndWithSameContentShouldBeEqual(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")

	set2 := NewWithoutValues[string]()
	set2.AddWithoutValue("apple")
	set2.AddWithoutValue("banana")
	set2.AddWithoutValue("cherry")

	// Expect
	assert.True(t, set1.Equals(set2))

	// When adding an already existing item
	set1.AddWithoutValue("apple")
	// Then the sets remain equal
	assert.True(t, set1.Equals(set2))

	// When
	set1.AddWithoutValue("brick")
	// Then
	assert.False(t, set1.Equals(set2))

	// When
	set2.AddWithoutValue("brick")
	// Then
	assert.True(t, set1.Equals(set2))

	// When
	set2.AddWithoutValue("stone")
	// Then
	assert.False(t, set1.Equals(set2))

	// When
	set1.AddWithoutValue("bottle")
	// Then
	assert.False(t, set1.Equals(set2))
}

func TestTwoSetsWithValuesAndWithSameContentShouldBeEqual(t *testing.T) {
	// Given
	set1 := NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")

	set2 := NewWithValues[string, string]()
	set2.AddWithValue("apple", "green")
	set2.AddWithValue("banana", "brownish")
	set2.AddWithValue("cherry", "glossy red")

	// Expect
	assert.True(t, set1.Equals(set2))

	// When adding an already existing item
	set1.AddWithValue("apple", "red but with a hint of green")
	// Then the sets remain equal
	assert.True(t, set1.Equals(set2))

	// When
	set1.AddWithValue("brick", "brick red")
	// Then
	assert.False(t, set1.Equals(set2))

	// When
	set2.AddWithValue("brick", "brick red with a hint of black")
	// Then
	assert.True(t, set1.Equals(set2))

	// When
	set2.AddWithValue("stone", "grey")
	// Then
	assert.False(t, set1.Equals(set2))

	// When
	set1.AddWithValue("bottle", "green")
	// Then
	assert.False(t, set1.Equals(set2))
}

func TestShouldAddAllGivenItemsWithoutValuesToSet(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")

	set2 := NewWithoutValues[string]()
	set2.AddWithoutValue("banana")
	set2.AddWithoutValue("cherry")
	set2.AddWithoutValue("mango")

	// When
	set1.AddAll(set2)

	// Then
	assert.Equal(t, 4, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.True(t, set1.Contains("mango"))

	// and set2 remains unchanged
	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))

	// When
	set1.AddAll(nil)
	// Then
	assert.Equal(t, 4, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.True(t, set1.Contains("mango"))
}

func TestShouldAddAllGivenItemsWithValuesToSet(t *testing.T) {
	// Given
	set1 := NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")

	set2 := NewWithValues[string, string]()
	set2.AddWithValue("banana", "brownish")
	set2.AddWithValue("cherry", "glossy red")
	set2.AddWithValue("mango", "green-orange")

	// When
	set1.AddAll(set2)

	// Then
	assert.Equal(t, 4, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.True(t, set1.Contains("mango"))
	assert.Equal(t, "red", set1.elements["apple"])
	assert.Equal(t, "brownish", set1.elements["banana"])
	assert.Equal(t, "glossy red", set1.elements["cherry"])
	assert.Equal(t, "green-orange", set1.elements["mango"])

	// and set2 remains unchanged
	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.elements["banana"])
	assert.Equal(t, "glossy red", set2.elements["cherry"])
	assert.Equal(t, "green-orange", set2.elements["mango"])

	// When
	set1.AddAll(nil)
	// Then
	assert.Equal(t, 4, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.True(t, set1.Contains("mango"))
	assert.Equal(t, "red", set1.elements["apple"])
	assert.Equal(t, "brownish", set1.elements["banana"])
	assert.Equal(t, "glossy red", set1.elements["cherry"])
	assert.Equal(t, "green-orange", set1.elements["mango"])
}

func TestShouldRemoveAllGivenItemsFromSet(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")

	set2 := NewWithoutValues[string]()
	set2.AddWithoutValue("banana")
	set2.AddWithoutValue("cherry")
	set2.AddWithoutValue("mango")

	// When
	set1.RemoveAll(set2)

	// Then
	assert.Equal(t, 1, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.False(t, set1.Contains("banana"))
	assert.False(t, set1.Contains("cherry"))
	assert.False(t, set1.Contains("mango"))

	// and set2 remains unchanged
	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))

	// When
	set1.RemoveAll(nil)
	// Then
	assert.Equal(t, 1, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.False(t, set1.Contains("banana"))
	assert.False(t, set1.Contains("cherry"))
	assert.False(t, set1.Contains("mango"))
}

func TestShouldCopySetWithoutValues(t *testing.T) {
	// Given
	set := NewWithoutValues[string]()
	set.AddWithoutValue("apple")
	set.AddWithoutValue("banana")
	set.AddWithoutValue("cherry")

	// When
	copiedSet := set.Copy()

	// Then
	assert.Equal(t, 3, copiedSet.Size())
	assert.True(t, copiedSet.Contains("apple"))
	assert.True(t, copiedSet.Contains("banana"))
	assert.True(t, copiedSet.Contains("cherry"))

	// When
	copiedSet.AddWithoutValue("mango")
	// Then
	assert.Equal(t, 4, copiedSet.Size())
	assert.True(t, copiedSet.Contains("apple"))
	assert.True(t, copiedSet.Contains("banana"))
	assert.True(t, copiedSet.Contains("cherry"))
	assert.True(t, copiedSet.Contains("mango"))

	// and set remains unchanged
	assert.Equal(t, 3, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
	assert.True(t, set.Contains("cherry"))
	assert.False(t, set.Contains("mango"))
}

func TestShouldCopySetWithValues(t *testing.T) {
	// Given
	set := NewWithValues[string, string]()
	set.AddWithValue("apple", "red")
	set.AddWithValue("banana", "yellow")
	set.AddWithValue("cherry", "dark red")

	// When
	copiedSet := set.Copy()

	// Then
	assert.Equal(t, 3, copiedSet.Size())
	assert.True(t, copiedSet.Contains("apple"))
	assert.True(t, copiedSet.Contains("banana"))
	assert.True(t, copiedSet.Contains("cherry"))
	assert.Equal(t, "red", copiedSet.elements["apple"])
	assert.Equal(t, "yellow", copiedSet.elements["banana"])
	assert.Equal(t, "dark red", copiedSet.elements["cherry"])

	// When
	copiedSet.AddWithValue("mango", "green-orange")
	// Then
	assert.Equal(t, 4, copiedSet.Size())
	assert.True(t, copiedSet.Contains("apple"))
	assert.True(t, copiedSet.Contains("banana"))
	assert.True(t, copiedSet.Contains("cherry"))
	assert.True(t, copiedSet.Contains("mango"))
	assert.Equal(t, "red", copiedSet.elements["apple"])
	assert.Equal(t, "yellow", copiedSet.elements["banana"])
	assert.Equal(t, "dark red", copiedSet.elements["cherry"])
	assert.Equal(t, "green-orange", copiedSet.elements["mango"])

	// and set remains unchanged
	assert.Equal(t, 3, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
	assert.True(t, set.Contains("cherry"))
	assert.False(t, set.Contains("mango"))
	assert.Equal(t, "red", set.elements["apple"])
	assert.Equal(t, "yellow", set.elements["banana"])
	assert.Equal(t, "dark red", set.elements["cherry"])
	assert.Equal(t, "", set.elements["mango"])
}

func TestShouldIntersectTwoSetsWithoutValues(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")

	set2 := NewWithoutValues[string]()
	set2.AddWithoutValue("banana")
	set2.AddWithoutValue("cherry")
	set2.AddWithoutValue("mango")

	// When
	intersectedSet := set1.Intersect(set2)

	// Then
	assert.Equal(t, 2, intersectedSet.Size())
	assert.True(t, intersectedSet.Contains("banana"))
	assert.True(t, intersectedSet.Contains("cherry"))
	assert.False(t, intersectedSet.Contains("apple"))
	assert.False(t, intersectedSet.Contains("mango"))

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))

	// When
	intersectedSet2 := set1.Intersect(nil)
	// Then
	assert.Equal(t, 0, intersectedSet2.Size())
	assert.False(t, intersectedSet2.Contains("apple"))
	assert.False(t, intersectedSet2.Contains("banana"))
	assert.False(t, intersectedSet2.Contains("cherry"))
	assert.False(t, intersectedSet2.Contains("mango"))
}

func TestShouldIntersectTwoSetsWithValues(t *testing.T) {
	// Given
	set1 := NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")

	set2 := NewWithValues[string, string]()
	set2.AddWithValue("banana", "brownish")
	set2.AddWithValue("cherry", "glossy red")
	set2.AddWithValue("mango", "green-orange")

	// When
	intersectedSet := set1.Intersect(set2)

	// Then
	assert.Equal(t, 2, intersectedSet.Size())
	assert.True(t, intersectedSet.Contains("banana"))
	assert.True(t, intersectedSet.Contains("cherry"))
	assert.False(t, intersectedSet.Contains("apple"))
	assert.False(t, intersectedSet.Contains("mango"))
	assert.Equal(t, "brownish", intersectedSet.elements["banana"])
	assert.Equal(t, "glossy red", intersectedSet.elements["cherry"])
	assert.Equal(t, "", intersectedSet.elements["apple"])
	assert.Equal(t, "", intersectedSet.elements["mango"])

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.Equal(t, "red", set1.elements["apple"])
	assert.Equal(t, "yellow", set1.elements["banana"])
	assert.Equal(t, "dark red", set1.elements["cherry"])

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.elements["banana"])
	assert.Equal(t, "glossy red", set2.elements["cherry"])
	assert.Equal(t, "green-orange", set2.elements["mango"])

	// When
	intersectedSet2 := set1.Intersect(nil)
	// Then
	assert.Equal(t, 0, intersectedSet2.Size())
	assert.False(t, intersectedSet2.Contains("apple"))
	assert.False(t, intersectedSet2.Contains("banana"))
	assert.False(t, intersectedSet2.Contains("cherry"))
	assert.False(t, intersectedSet2.Contains("mango"))
	assert.Equal(t, "", intersectedSet2.elements["apple"])
	assert.Equal(t, "", intersectedSet2.elements["banana"])
	assert.Equal(t, "", intersectedSet2.elements["cherry"])
	assert.Equal(t, "", intersectedSet2.elements["mango"])
}

func TestShouldUniteTwoSetsWithoutValues(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")

	set2 := NewWithoutValues[string]()
	set2.AddWithoutValue("banana")
	set2.AddWithoutValue("cherry")
	set2.AddWithoutValue("mango")

	// When
	unitedSet := set1.Unite(set2)

	// Then
	assert.Equal(t, 4, unitedSet.Size())
	assert.True(t, unitedSet.Contains("apple"))
	assert.True(t, unitedSet.Contains("banana"))
	assert.True(t, unitedSet.Contains("cherry"))
	assert.True(t, unitedSet.Contains("mango"))

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))

	// When
	unitedSet2 := set1.Unite(nil)
	// Then
	assert.Equal(t, 3, unitedSet2.Size())
	assert.True(t, unitedSet2.Contains("apple"))
	assert.True(t, unitedSet2.Contains("banana"))
	assert.True(t, unitedSet2.Contains("cherry"))
	assert.False(t, unitedSet2.Contains("mango"))
}

func TestShouldUniteTwoSetsWithValues(t *testing.T) {
	// Given
	set1 := NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")

	set2 := NewWithValues[string, string]()
	set2.AddWithValue("banana", "brownish")
	set2.AddWithValue("cherry", "glossy red")
	set2.AddWithValue("mango", "green-orange")

	// When
	unitedSet := set1.Unite(set2)

	// Then
	assert.Equal(t, 4, unitedSet.Size())
	assert.True(t, unitedSet.Contains("apple"))
	assert.True(t, unitedSet.Contains("banana"))
	assert.True(t, unitedSet.Contains("cherry"))
	assert.True(t, unitedSet.Contains("mango"))
	assert.Equal(t, "red", unitedSet.elements["apple"])
	assert.Equal(t, "brownish", unitedSet.elements["banana"])
	assert.Equal(t, "glossy red", unitedSet.elements["cherry"])
	assert.Equal(t, "green-orange", unitedSet.elements["mango"])

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.Equal(t, "red", set1.elements["apple"])
	assert.Equal(t, "yellow", set1.elements["banana"])
	assert.Equal(t, "dark red", set1.elements["cherry"])

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.elements["banana"])
	assert.Equal(t, "glossy red", set2.elements["cherry"])
	assert.Equal(t, "green-orange", set2.elements["mango"])

	// When
	unitedSet2 := set1.Unite(nil)
	// Then
	assert.Equal(t, 3, unitedSet2.Size())
	assert.True(t, unitedSet2.Contains("apple"))
	assert.True(t, unitedSet2.Contains("banana"))
	assert.True(t, unitedSet2.Contains("cherry"))
	assert.False(t, unitedSet2.Contains("mango"))
	assert.Equal(t, "red", unitedSet2.elements["apple"])
	assert.Equal(t, "yellow", unitedSet2.elements["banana"])
	assert.Equal(t, "dark red", unitedSet2.elements["cherry"])
}

func TestShouldSubtractOneSetFromAnotherWithoutValues(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")

	set2 := NewWithoutValues[string]()
	set2.AddWithoutValue("banana")
	set2.AddWithoutValue("cherry")
	set2.AddWithoutValue("mango")

	// When
	subtractedSet := set1.Subtract(set2)

	// Then
	assert.Equal(t, 1, subtractedSet.Size())
	assert.True(t, subtractedSet.Contains("apple"))
	assert.False(t, subtractedSet.Contains("banana"))
	assert.False(t, subtractedSet.Contains("cherry"))
	assert.False(t, subtractedSet.Contains("mango"))

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))

	// When
	subtractedSet2 := set1.Subtract(nil)
	// Then
	assert.Equal(t, 3, subtractedSet2.Size())
	assert.True(t, subtractedSet2.Contains("apple"))
	assert.True(t, subtractedSet2.Contains("banana"))
	assert.True(t, subtractedSet2.Contains("cherry"))
	assert.False(t, subtractedSet2.Contains("mango"))
}

func TestShouldSubtractOneSetFromAnotherWithValues(t *testing.T) {
	// Given
	set1 := NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")

	set2 := NewWithValues[string, string]()
	set2.AddWithValue("banana", "brownish")
	set2.AddWithValue("cherry", "glossy red")
	set2.AddWithValue("mango", "green-orange")

	// When
	subtractedSet := set1.Subtract(set2)

	// Then
	assert.Equal(t, 1, subtractedSet.Size())
	assert.True(t, subtractedSet.Contains("apple"))
	assert.False(t, subtractedSet.Contains("banana"))
	assert.False(t, subtractedSet.Contains("cherry"))
	assert.False(t, subtractedSet.Contains("mango"))
	assert.Equal(t, "red", subtractedSet.elements["apple"])

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.Equal(t, "red", set1.elements["apple"])
	assert.Equal(t, "yellow", set1.elements["banana"])
	assert.Equal(t, "dark red", set1.elements["cherry"])

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.elements["banana"])
	assert.Equal(t, "glossy red", set2.elements["cherry"])
	assert.Equal(t, "green-orange", set2.elements["mango"])

	// When
	subtractedSet2 := set1.Subtract(nil)
	// Then
	assert.Equal(t, 3, subtractedSet2.Size())
	assert.True(t, subtractedSet2.Contains("apple"))
	assert.True(t, subtractedSet2.Contains("banana"))
	assert.True(t, subtractedSet2.Contains("cherry"))
	assert.False(t, subtractedSet2.Contains("mango"))
	assert.Equal(t, "red", subtractedSet2.elements["apple"])
	assert.Equal(t, "yellow", subtractedSet2.elements["banana"])
	assert.Equal(t, "dark red", subtractedSet2.elements["cherry"])
}

func TestShouldCreateDisjunctiveUnionOfTwoSetsWithoutValues(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")

	set2 := NewWithoutValues[string]()
	set2.AddWithoutValue("banana")
	set2.AddWithoutValue("cherry")
	set2.AddWithoutValue("mango")

	// When
	symDifferenceSet := set1.UniteDisjunctively(set2)

	// Then
	assert.Equal(t, 2, symDifferenceSet.Size())
	assert.True(t, symDifferenceSet.Contains("apple"))
	assert.True(t, symDifferenceSet.Contains("mango"))
	assert.False(t, symDifferenceSet.Contains("banana"))
	assert.False(t, symDifferenceSet.Contains("cherry"))

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))

	// When
	symDifferenceSet2 := set1.UniteDisjunctively(nil)
	// Then
	assert.Equal(t, 3, symDifferenceSet2.Size())
	assert.True(t, symDifferenceSet2.Contains("apple"))
	assert.True(t, symDifferenceSet2.Contains("banana"))
	assert.True(t, symDifferenceSet2.Contains("cherry"))
	assert.False(t, symDifferenceSet2.Contains("mango"))
}

func TestShouldCreateDisjunctiveUnionOfTwoSetsWithValues(t *testing.T) {
	// Given
	set1 := NewWithValues[string, string]()
	set1.AddWithValue("apple", "red")
	set1.AddWithValue("banana", "yellow")
	set1.AddWithValue("cherry", "dark red")

	set2 := NewWithValues[string, string]()
	set2.AddWithValue("banana", "brownish")
	set2.AddWithValue("cherry", "glossy red")
	set2.AddWithValue("mango", "green-orange")

	// When
	symDifferenceSet := set1.UniteDisjunctively(set2)

	// Then
	assert.Equal(t, 2, symDifferenceSet.Size())
	assert.True(t, symDifferenceSet.Contains("apple"))
	assert.True(t, symDifferenceSet.Contains("mango"))
	assert.False(t, symDifferenceSet.Contains("banana"))
	assert.False(t, symDifferenceSet.Contains("cherry"))
	assert.Equal(t, "red", symDifferenceSet.elements["apple"])
	assert.Equal(t, "green-orange", symDifferenceSet.elements["mango"])
	assert.Equal(t, "", symDifferenceSet.elements["banana"])
	assert.Equal(t, "", symDifferenceSet.elements["cherry"])

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.Equal(t, "red", set1.elements["apple"])
	assert.Equal(t, "yellow", set1.elements["banana"])
	assert.Equal(t, "dark red", set1.elements["cherry"])

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.elements["banana"])
	assert.Equal(t, "glossy red", set2.elements["cherry"])
	assert.Equal(t, "green-orange", set2.elements["mango"])

	// When
	symDifferenceSet2 := set1.UniteDisjunctively(nil)
	// Then
	assert.Equal(t, 3, symDifferenceSet2.Size())
	assert.True(t, symDifferenceSet2.Contains("apple"))
	assert.True(t, symDifferenceSet2.Contains("banana"))
	assert.True(t, symDifferenceSet2.Contains("cherry"))
	assert.False(t, symDifferenceSet2.Contains("mango"))
	assert.Equal(t, "red", symDifferenceSet2.elements["apple"])
	assert.Equal(t, "yellow", symDifferenceSet2.elements["banana"])
	assert.Equal(t, "dark red", symDifferenceSet2.elements["cherry"])
	assert.Equal(t, "", symDifferenceSet2.elements["mango"])
}

func TestShouldBeSubset(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")

	set2 := NewWithoutValues[string]()
	set2.AddWithoutValue("banana")
	set2.AddWithoutValue("cherry")

	// Then
	assert.True(t, set2.IsSubset(set1))
	assert.False(t, set1.IsSubset(set2))

	// When
	set2.AddWithoutValue("apple")
	// Then
	assert.True(t, set2.IsSubset(set1))
	assert.True(t, set1.IsSubset(set2))

	// When
	set2.AddWithoutValue("mango")
	// Then
	assert.False(t, set2.IsSubset(set1))
	assert.True(t, set1.IsSubset(set2))

	// When
	set1.AddWithoutValue("brick")
	// Then
	assert.False(t, set2.IsSubset(set1))
	assert.False(t, set1.IsSubset(set2))

	// Expect
	assert.False(t, set1.IsSubset(nil))
	assert.False(t, set2.IsSubset(nil))

	// Given
	emptySet := NewWithoutValues[string]()
	// Expect
	assert.True(t, emptySet.IsSubset(emptySet))
	assert.True(t, emptySet.IsSubset(nil))
}

func TestShouldFilterSetWithoutValues(t *testing.T) {
	// Given
	set := NewWithoutValues[string]()
	set.AddWithoutValue("apple")
	set.AddWithoutValue("banana")
	set.AddWithoutValue("cherry")
	set.AddWithoutValue("mango")
	set.AddWithoutValue("orange")
	set.AddWithoutValue("pear")

	// When
	filteredSet := set.Filter(func(elem string) bool {
		return strings.Contains(elem, "an")
	})

	// Then
	assert.Equal(t, 3, filteredSet.Size())
	assert.True(t, filteredSet.Contains("banana"))
	assert.True(t, filteredSet.Contains("mango"))
	assert.True(t, filteredSet.Contains("orange"))
	assert.False(t, filteredSet.Contains("apple"))
	assert.False(t, filteredSet.Contains("cherry"))
	assert.False(t, filteredSet.Contains("pear"))

	// When the filter function is nil
	filteredSet2 := set.Filter(nil)
	// Then all elements are included (b/c there is no filter)
	assert.Equal(t, 6, filteredSet2.Size())
	assert.True(t, filteredSet2.Contains("apple"))
	assert.True(t, filteredSet2.Contains("banana"))
	assert.True(t, filteredSet2.Contains("cherry"))
	assert.True(t, filteredSet2.Contains("mango"))
	assert.True(t, filteredSet2.Contains("orange"))
	assert.True(t, filteredSet2.Contains("pear"))
}

func TestShouldFilterSetWithValues(t *testing.T) {
	// Given
	set := NewWithValues[string, string]()
	set.AddWithValue("apple", "red")
	set.AddWithValue("banana", "yellow")
	set.AddWithValue("cherry", "dark red")
	set.AddWithValue("mango", "green-orange")
	set.AddWithValue("orange", "orange")
	set.AddWithValue("pear", "green")

	// When
	filteredSet := set.Filter(func(elem string) bool {
		return strings.Contains(elem, "an")
	})

	// Then
	assert.Equal(t, 3, filteredSet.Size())
	assert.True(t, filteredSet.Contains("banana"))
	assert.True(t, filteredSet.Contains("mango"))
	assert.True(t, filteredSet.Contains("orange"))
	assert.False(t, filteredSet.Contains("apple"))
	assert.False(t, filteredSet.Contains("cherry"))
	assert.False(t, filteredSet.Contains("pear"))
	assert.Equal(t, "yellow", filteredSet.elements["banana"])
	assert.Equal(t, "green-orange", filteredSet.elements["mango"])
	assert.Equal(t, "orange", filteredSet.elements["orange"])
	assert.Equal(t, "", filteredSet.elements["apple"])
	assert.Equal(t, "", filteredSet.elements["cherry"])
	assert.Equal(t, "", filteredSet.elements["pear"])

	// When the filter function is nil
	filteredSet2 := set.Filter(nil)
	// Then all elements are included (b/c there is no filter)
	assert.Equal(t, 6, filteredSet2.Size())
	assert.True(t, filteredSet2.Contains("apple"))
	assert.True(t, filteredSet2.Contains("banana"))
	assert.True(t, filteredSet2.Contains("cherry"))
	assert.True(t, filteredSet2.Contains("mango"))
	assert.True(t, filteredSet2.Contains("orange"))
	assert.True(t, filteredSet2.Contains("pear"))
	assert.Equal(t, "red", filteredSet2.elements["apple"])
	assert.Equal(t, "yellow", filteredSet2.elements["banana"])
	assert.Equal(t, "dark red", filteredSet2.elements["cherry"])
	assert.Equal(t, "green-orange", filteredSet2.elements["mango"])
	assert.Equal(t, "orange", filteredSet2.elements["orange"])
	assert.Equal(t, "green", filteredSet2.elements["pear"])
}
