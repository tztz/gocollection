package set

import (
	"fmt"
	"slices"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestZeroValueOfSetIsAnEmptySet(t *testing.T) {
	// When
	var set1 tzSet[string, string]
	// Then
	assert.Equal(t, 0, set1.Size())
	assert.Equal(t, []string{}, set1.List())

	// When
	set2 := NewWithoutValues[string]()
	// Then
	assert.Equal(t, 0, set2.Size())
	assert.Equal(t, []string{}, set2.List())

	// When
	set3 := NewWithValues[string, string]()
	// Then
	assert.Equal(t, 0, set3.Size())
	assert.Equal(t, []string{}, set3.List())
}

func TestShouldCreateNewMapWithValues(t *testing.T) {
	// Given
	type myKeyType string
	type myValueType string

	// When
	var m map[myKeyType]myValueType
	// Then
	assert.Nil(t, m)

	// When
	m = createNewWithValues[myKeyType, myValueType]()
	// Then
	assert.NotNil(t, m)
	assert.Equal(t, 0, len(m))
	assert.Equal(t, myValueType(""), m[myKeyType("")])

	// When
	m[myKeyType("apple")] = myValueType("red")
	m[myKeyType("banana")] = myValueType("yellow")
	// Then
	assert.Equal(t, 2, len(m))
	assert.Equal(t, myValueType("red"), m[myKeyType("apple")])
	assert.Equal(t, myValueType("yellow"), m[myKeyType("banana")])

	// When
	clear(m)
	// Then
	assert.Equal(t, 0, len(m))
}

func TestShouldReturnElementsOfSet(t *testing.T) {
	// Given
	set := NewWithValues[string, string]()

	// When
	elements := set.GetElements()
	// Then
	assert.Equal(t, 0, set.Size())
	assert.Equal(t, 0, len(elements))

	// When
	set.AddWithValue("apple", "red")
	set.AddWithValue("banana", "yellow")
	set.AddWithValue("cherry", "dark red")

	// Then
	assert.Equal(t, 3, set.Size())
	assert.Equal(t, 3, len(elements))
	assert.Equal(t, "red", elements["apple"])
	assert.Equal(t, "yellow", elements["banana"])
	assert.Equal(t, "dark red", elements["cherry"])
}

func TestShouldReturnRandomInt64Numbers(t *testing.T) {
	// Given
	set := NewWithoutValues[string]()

	// Expect
	assert.Equal(t, int64(-1), set.randIndex())

	// When
	set.AddWithoutValue("apple")
	set.AddWithoutValue("banana")
	set.AddWithoutValue("cherry")
	set.AddWithoutValue("mango")
	set.AddWithoutValue("orange")
	set.AddWithoutValue("pear")
	set.AddWithoutValue("pineapple")
	set.AddWithoutValue("watermelon")
	set.AddWithoutValue("kiwi")
	set.AddWithoutValue("grape")
	set.AddWithoutValue("strawberry")
	set.AddWithoutValue("blueberry")
	set.AddWithoutValue("blackberry")
	set.AddWithoutValue("raspberry")
	set.AddWithoutValue("papaya")
	set.AddWithoutValue("guava")
	set.AddWithoutValue("lychee")
	set.AddWithoutValue("passion fruit")
	set.AddWithoutValue("dragon fruit")
	set.AddWithoutValue("star fruit")
	// and
	count := 100_000
	var randomNumbers = make([]int64, 0, count)
	for i := 0; i < count; i++ {
		randomNumbers = append(randomNumbers, set.randIndex())
	}

	// Then
	setSize := int64(set.Size())
	for i := 0; i < count; i++ {
		var randomNumber int64 = randomNumbers[i]
		assert.True(t, randomNumber >= 0)
		assert.True(t, randomNumber <= setSize-1)
	}
	/* This is theoretically flaky, but practically it should work
	// and
	numbersMap := map[int64]bool{}
	assert.Equal(t, 0, len(numbersMap))
	for i := 0; i < count; i++ {
		numbersMap[randomNumbers[i]] = true
	}
	assert.Equal(t, set.Size(), len(numbersMap))
	*/
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
	assert.Equal(t, "red", set.GetElements()["apple"])

	// When
	set.AddWithValue("banana", "yellow")
	// Then
	assert.Equal(t, 2, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
	assert.Equal(t, "red", set.GetElements()["apple"])
	assert.Equal(t, "yellow", set.GetElements()["banana"])
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
	assert.Equal(t, "red", set.GetElements()["apple"])
	assert.Equal(t, "yellow", set.GetElements()["banana"])

	// When
	set.AddWithValue("apple", "light red")
	set.AddWithValue("banana", "dark yellow")

	// Then
	assert.Equal(t, "light red", set.GetElements()["apple"])
	assert.Equal(t, "dark yellow", set.GetElements()["banana"])
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

	// When removing non-existing element
	set.Remove("brick")
	// Then nothing changes
	assert.Equal(t, 1, set.Size())
	assert.False(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
	assert.False(t, set.Contains("brick"))

	// When
	set.Remove("banana")
	// Then
	assert.Equal(t, 0, set.Size())
	assert.False(t, set.Contains("apple"))
	assert.False(t, set.Contains("banana"))

	// When removing non-existing element
	set.Remove("brick")
	// Then nothing changes
	assert.Equal(t, 0, set.Size())
	assert.False(t, set.Contains("apple"))
	assert.False(t, set.Contains("banana"))
	assert.False(t, set.Contains("brick"))
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
	assert.Equal(t, "red", set1.GetElements()["apple"])
	assert.Equal(t, "brownish", set1.GetElements()["banana"])
	assert.Equal(t, "glossy red", set1.GetElements()["cherry"])
	assert.Equal(t, "green-orange", set1.GetElements()["mango"])

	// and set2 remains unchanged
	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.GetElements()["banana"])
	assert.Equal(t, "glossy red", set2.GetElements()["cherry"])
	assert.Equal(t, "green-orange", set2.GetElements()["mango"])

	// When
	set1.AddAll(nil)
	// Then
	assert.Equal(t, 4, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.True(t, set1.Contains("mango"))
	assert.Equal(t, "red", set1.GetElements()["apple"])
	assert.Equal(t, "brownish", set1.GetElements()["banana"])
	assert.Equal(t, "glossy red", set1.GetElements()["cherry"])
	assert.Equal(t, "green-orange", set1.GetElements()["mango"])
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

	// When removing non-existing elements
	setOfAliens := NewWithoutValues[string]()
	setOfAliens.AddWithoutValue("alien1")
	setOfAliens.AddWithoutValue("alien2")
	set1.RemoveAll(setOfAliens)
	// Then nothing changes
	assert.Equal(t, 1, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.False(t, set1.Contains("banana"))
	assert.False(t, set1.Contains("cherry"))
	assert.False(t, set1.Contains("mango"))
	assert.False(t, set1.Contains("alien1"))
	assert.False(t, set1.Contains("alien2"))
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

func TestShouldCalculateSizeOfSet(t *testing.T) {
	// Given
	set := NewWithoutValues[string]()

	// Expect
	assert.Equal(t, 0, set.Size())

	// When
	set.AddWithoutValue("apple")
	set.AddWithoutValue("banana")
	set.AddWithoutValue("cherry")

	// Then
	assert.Equal(t, 3, set.Size())

	// When
	set.AddWithoutValue("mango")
	set.AddWithoutValue("pear")

	// Then
	assert.Equal(t, 5, set.Size())
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

// TODO: Add tests for Contains

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
	assert.Equal(t, "red", copiedSet.GetElements()["apple"])
	assert.Equal(t, "yellow", copiedSet.GetElements()["banana"])
	assert.Equal(t, "dark red", copiedSet.GetElements()["cherry"])

	// When
	copiedSet.AddWithValue("mango", "green-orange")
	// Then
	assert.Equal(t, 4, copiedSet.Size())
	assert.True(t, copiedSet.Contains("apple"))
	assert.True(t, copiedSet.Contains("banana"))
	assert.True(t, copiedSet.Contains("cherry"))
	assert.True(t, copiedSet.Contains("mango"))
	assert.Equal(t, "red", copiedSet.GetElements()["apple"])
	assert.Equal(t, "yellow", copiedSet.GetElements()["banana"])
	assert.Equal(t, "dark red", copiedSet.GetElements()["cherry"])
	assert.Equal(t, "green-orange", copiedSet.GetElements()["mango"])

	// and set remains unchanged
	assert.Equal(t, 3, set.Size())
	assert.True(t, set.Contains("apple"))
	assert.True(t, set.Contains("banana"))
	assert.True(t, set.Contains("cherry"))
	assert.False(t, set.Contains("mango"))
	assert.Equal(t, "red", set.GetElements()["apple"])
	assert.Equal(t, "yellow", set.GetElements()["banana"])
	assert.Equal(t, "dark red", set.GetElements()["cherry"])
	assert.Equal(t, "", set.GetElements()["mango"])
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

	// When intersecting with aliens
	setOfAliens := NewWithoutValues[string]()
	setOfAliens.AddWithoutValue("alien1")
	setOfAliens.AddWithoutValue("alien2")
	intersectedSet3 := set1.Intersect(setOfAliens)
	// Then the intersection is empty
	assert.Equal(t, 0, intersectedSet3.Size())
	assert.False(t, intersectedSet3.Contains("apple"))
	assert.False(t, intersectedSet3.Contains("banana"))
	assert.False(t, intersectedSet3.Contains("cherry"))
	assert.False(t, intersectedSet3.Contains("mango"))
	assert.False(t, intersectedSet3.Contains("alien1"))
	assert.False(t, intersectedSet3.Contains("alien2"))
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
	assert.Equal(t, "brownish", intersectedSet.GetElements()["banana"])
	assert.Equal(t, "glossy red", intersectedSet.GetElements()["cherry"])
	assert.Equal(t, "", intersectedSet.GetElements()["apple"])
	assert.Equal(t, "", intersectedSet.GetElements()["mango"])

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.Equal(t, "red", set1.GetElements()["apple"])
	assert.Equal(t, "yellow", set1.GetElements()["banana"])
	assert.Equal(t, "dark red", set1.GetElements()["cherry"])

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.GetElements()["banana"])
	assert.Equal(t, "glossy red", set2.GetElements()["cherry"])
	assert.Equal(t, "green-orange", set2.GetElements()["mango"])

	// When
	intersectedSet2 := set1.Intersect(nil)
	// Then
	assert.Equal(t, 0, intersectedSet2.Size())
	assert.False(t, intersectedSet2.Contains("apple"))
	assert.False(t, intersectedSet2.Contains("banana"))
	assert.False(t, intersectedSet2.Contains("cherry"))
	assert.False(t, intersectedSet2.Contains("mango"))
	assert.Equal(t, "", intersectedSet2.GetElements()["apple"])
	assert.Equal(t, "", intersectedSet2.GetElements()["banana"])
	assert.Equal(t, "", intersectedSet2.GetElements()["cherry"])
	assert.Equal(t, "", intersectedSet2.GetElements()["mango"])
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
	assert.Equal(t, "red", unitedSet.GetElements()["apple"])
	assert.Equal(t, "brownish", unitedSet.GetElements()["banana"])
	assert.Equal(t, "glossy red", unitedSet.GetElements()["cherry"])
	assert.Equal(t, "green-orange", unitedSet.GetElements()["mango"])

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.Equal(t, "red", set1.GetElements()["apple"])
	assert.Equal(t, "yellow", set1.GetElements()["banana"])
	assert.Equal(t, "dark red", set1.GetElements()["cherry"])

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.GetElements()["banana"])
	assert.Equal(t, "glossy red", set2.GetElements()["cherry"])
	assert.Equal(t, "green-orange", set2.GetElements()["mango"])

	// When
	unitedSet2 := set1.Unite(nil)
	// Then
	assert.Equal(t, 3, unitedSet2.Size())
	assert.True(t, unitedSet2.Contains("apple"))
	assert.True(t, unitedSet2.Contains("banana"))
	assert.True(t, unitedSet2.Contains("cherry"))
	assert.False(t, unitedSet2.Contains("mango"))
	assert.Equal(t, "red", unitedSet2.GetElements()["apple"])
	assert.Equal(t, "yellow", unitedSet2.GetElements()["banana"])
	assert.Equal(t, "dark red", unitedSet2.GetElements()["cherry"])
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
	assert.Equal(t, "red", symDifferenceSet.GetElements()["apple"])
	assert.Equal(t, "green-orange", symDifferenceSet.GetElements()["mango"])
	assert.Equal(t, "", symDifferenceSet.GetElements()["banana"])
	assert.Equal(t, "", symDifferenceSet.GetElements()["cherry"])

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.Equal(t, "red", set1.GetElements()["apple"])
	assert.Equal(t, "yellow", set1.GetElements()["banana"])
	assert.Equal(t, "dark red", set1.GetElements()["cherry"])

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.GetElements()["banana"])
	assert.Equal(t, "glossy red", set2.GetElements()["cherry"])
	assert.Equal(t, "green-orange", set2.GetElements()["mango"])

	// When
	symDifferenceSet2 := set1.UniteDisjunctively(nil)
	// Then
	assert.Equal(t, 3, symDifferenceSet2.Size())
	assert.True(t, symDifferenceSet2.Contains("apple"))
	assert.True(t, symDifferenceSet2.Contains("banana"))
	assert.True(t, symDifferenceSet2.Contains("cherry"))
	assert.False(t, symDifferenceSet2.Contains("mango"))
	assert.Equal(t, "red", symDifferenceSet2.GetElements()["apple"])
	assert.Equal(t, "yellow", symDifferenceSet2.GetElements()["banana"])
	assert.Equal(t, "dark red", symDifferenceSet2.GetElements()["cherry"])
	assert.Equal(t, "", symDifferenceSet2.GetElements()["mango"])
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

	// When subtracting non-existing elements
	setOfAliens := NewWithoutValues[string]()
	setOfAliens.AddWithoutValue("alien1")
	setOfAliens.AddWithoutValue("alien2")
	subtractedSet3 := set1.Subtract(setOfAliens)
	// Then nothing changes
	assert.Equal(t, 3, subtractedSet3.Size())
	assert.True(t, subtractedSet3.Contains("apple"))
	assert.True(t, subtractedSet3.Contains("banana"))
	assert.True(t, subtractedSet3.Contains("cherry"))
	assert.False(t, subtractedSet3.Contains("mango"))
	assert.False(t, subtractedSet3.Contains("alien1"))
	assert.False(t, subtractedSet3.Contains("alien2"))
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
	assert.Equal(t, "red", subtractedSet.GetElements()["apple"])

	// and set1 and set2 remain unchanged
	assert.Equal(t, 3, set1.Size())
	assert.True(t, set1.Contains("apple"))
	assert.True(t, set1.Contains("banana"))
	assert.True(t, set1.Contains("cherry"))
	assert.Equal(t, "red", set1.GetElements()["apple"])
	assert.Equal(t, "yellow", set1.GetElements()["banana"])
	assert.Equal(t, "dark red", set1.GetElements()["cherry"])

	assert.Equal(t, 3, set2.Size())
	assert.True(t, set2.Contains("banana"))
	assert.True(t, set2.Contains("cherry"))
	assert.True(t, set2.Contains("mango"))
	assert.Equal(t, "brownish", set2.GetElements()["banana"])
	assert.Equal(t, "glossy red", set2.GetElements()["cherry"])
	assert.Equal(t, "green-orange", set2.GetElements()["mango"])

	// When
	subtractedSet2 := set1.Subtract(nil)
	// Then
	assert.Equal(t, 3, subtractedSet2.Size())
	assert.True(t, subtractedSet2.Contains("apple"))
	assert.True(t, subtractedSet2.Contains("banana"))
	assert.True(t, subtractedSet2.Contains("cherry"))
	assert.False(t, subtractedSet2.Contains("mango"))
	assert.Equal(t, "red", subtractedSet2.GetElements()["apple"])
	assert.Equal(t, "yellow", subtractedSet2.GetElements()["banana"])
	assert.Equal(t, "dark red", subtractedSet2.GetElements()["cherry"])
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
	assert.Equal(t, "yellow", filteredSet.GetElements()["banana"])
	assert.Equal(t, "green-orange", filteredSet.GetElements()["mango"])
	assert.Equal(t, "orange", filteredSet.GetElements()["orange"])
	assert.Equal(t, "", filteredSet.GetElements()["apple"])
	assert.Equal(t, "", filteredSet.GetElements()["cherry"])
	assert.Equal(t, "", filteredSet.GetElements()["pear"])

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
	assert.Equal(t, "red", filteredSet2.GetElements()["apple"])
	assert.Equal(t, "yellow", filteredSet2.GetElements()["banana"])
	assert.Equal(t, "dark red", filteredSet2.GetElements()["cherry"])
	assert.Equal(t, "green-orange", filteredSet2.GetElements()["mango"])
	assert.Equal(t, "orange", filteredSet2.GetElements()["orange"])
	assert.Equal(t, "green", filteredSet2.GetElements()["pear"])
}

func TestShouldGetOneRandomElementFromSet(t *testing.T) {
	// Given
	set1 := NewWithoutValues[string]()
	set1.AddWithoutValue("apple")
	set1.AddWithoutValue("banana")
	set1.AddWithoutValue("cherry")
	set1.AddWithoutValue("mango")
	set1.AddWithoutValue("orange")
	set1.AddWithoutValue("pear")

	// When
	randomElement1, randomValue1, err1 := set1.OneR()

	// Then
	assert.True(t, set1.Contains(randomElement1))
	assert.Equal(t, internalEmptyValue, randomValue1)
	assert.Nil(t, err1)

	// Given
	set2 := NewWithValues[string, string]()
	set2.AddWithValue("apple", "red")
	set2.AddWithValue("banana", "yellow")
	set2.AddWithValue("cherry", "dark red")
	set2.AddWithValue("mango", "green-orange")
	set2.AddWithValue("orange", "orange")
	set2.AddWithValue("pear", "green")

	// When
	randomElement2, randomValue2, err2 := set2.OneR()

	// Then
	assert.True(t, set2.Contains(randomElement2))
	assert.Equal(t, randomValue2, set2.GetElements()[randomElement2])
	assert.Nil(t, err2)

	// Given
	set3 := NewWithoutValues[string]()

	// When
	randomElement3, randomValue3, err3 := set3.OneR()

	// Then
	assert.Equal(t, "", randomElement3)
	assert.Equal(t, internalEmptyValue, randomValue3)
	assert.NotNil(t, err3)
	assert.Equal(t, fmt.Errorf("cannot get a random element from set because it is empty"), err3)
}
