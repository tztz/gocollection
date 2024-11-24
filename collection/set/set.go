// An API to handle sets providing methods from set theory.
package set

import (
	"fmt"
	"strings"
)

type internalEmptyType struct{}
type FilterFunc[T comparable] func(T) bool

// Set is a collection of unique elements having the same type T.
// Values of type V can be associated with the elements - but don't have to.
// A Set is actually a map[T]V with keys of type T and values of type V, where the values are just associated data.
// If you don't need values, you can omit them in the Set to save memory. Then it's just a set of elements like a set of labels.
// A Set can, of course, be empty.
// The zero value of a Set is an empty set.
type Set[T comparable, V any] struct {
	elements map[T]V
}

// NewWithValues creates a new, empty set that can contain elements of type T having values of type V (like a map).
func NewWithValues[T comparable, V any]() *Set[T, V] {
	return &Set[T, V]{
		elements: createNewWithValues[T, V](),
	}
}

// NewWithoutValues creates a new, empty set that can contain elements of type T (like a set of labels).
func NewWithoutValues[T comparable]() *Set[T, internalEmptyType] {
	return &Set[T, internalEmptyType]{
		elements: createNewWithValues[T, internalEmptyType](),
	}
}

func createNewWithValues[T comparable, V any]() map[T]V {
	return make(map[T]V)
}

// AddWithValue adds an element with an associated value to the set.
func (s *Set[T, V]) AddWithValue(element T, value V) {
	s.elements[element] = value
}

// AddWithoutValue adds an element (without an associated value) to the set.
func (s *Set[T, V]) AddWithoutValue(element T) {
	var empty V
	s.elements[element] = empty
}

// Remove removes an element from the set.
func (s *Set[T, V]) Remove(element T) {
	delete(s.elements, element)
}

// AddAll adds all elements (including the value) from otherSet to this set.
// If otherSet is nil, nothing happens.
// If an element already exists in this set, the value is overwritten with the value from otherSet.
// The otherSet remains unchanged.
func (s *Set[T, V]) AddAll(otherSet *Set[T, V]) {
	if otherSet == nil {
		return
	}
	for elem, value := range otherSet.elements {
		s.elements[elem] = value
	}
}

// RemoveAll removes all elements from otherSet from this set.
// If otherSet is nil, nothing happens.
// The otherSet remains unchanged.
func (s *Set[T, V]) RemoveAll(otherSet *Set[T, V]) {
	if otherSet == nil {
		return
	}
	for elem := range otherSet.elements {
		delete(s.elements, elem)
	}
}

// Clear removes all elements from the set.
func (s *Set[T, V]) Clear() {
	s.elements = createNewWithValues[T, V]()
}

// Size returns the number of elements in the set.
func (s *Set[T, V]) Size() int {
	return len(s.elements)
}

// List returns all elements (without values) of the set as a slice.
// The returned slice is a copy, changes to that copy do not interfere with the original set.
func (s *Set[T, V]) List() []T {
	elements := make([]T, 0, s.Size())
	for elem := range s.elements {
		elements = append(elements, elem)
	}
	return elements
}

// Contains checks whether or not the given element exists in the set (ignoring the value).
// Returns true if the element is in the set, false otherwise.
// The value associated with the element is not considered, i.e. it doesn't matter whether
// the given element's value is different from the element's value in this set.
func (s *Set[T, V]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Equals checks if this set is equal to otherSet ignoring the values.
// Returns true if both sets are of equal size and contain the same elements (ignoring the values), false otherwise.
func (s *Set[T, V]) Equals(otherSet *Set[T, V]) bool {
	if s.Size() != otherSet.Size() {
		return false
	}
	for elem := range s.elements {
		if !otherSet.Contains(elem) {
			return false
		}
	}
	return true
}

// IsSubset checks if this set is a subset of otherSet.
// Returns true if all elements of this set are in otherSet, false otherwise.
// If otherSet is nil and this set is not empty, false is returned.
// If otherSet is nil and this set is empty, true is returned.
// The values are not considered when checking for subset.
func (s *Set[T, V]) IsSubset(otherSet *Set[T, V]) bool {
	if otherSet == nil && s.Size() > 0 {
		return false
	}
	if otherSet == nil && s.Size() == 0 {
		return true
	}
	if s.Size() > otherSet.Size() {
		return false
	}
	for elem := range s.elements {
		if !otherSet.Contains(elem) {
			return false
		}
	}
	return true
}

// String returns a string representation of the set.
// The elements are separated by commas.
// The elements are converted to strings using the fmt package.
// The values are not included in the string representation.
// The order of the elements is not defined.
// If the set is empty, an empty string is returned.
func (s *Set[T, V]) String() string {
	strElems := make([]string, 0, s.Size())
	for elem := range s.elements {
		strElems = append(strElems, fmt.Sprintf("%v", elem))
	}
	return strings.Join(strElems, ", ")
}

// StringWithValues returns a string representation of the set including values.
// The elements are separated by commas.
// The elements and values are converted to strings using the fmt package.
// Each element's value is given in braces after the element.
// The order of the elements is not defined.
// If the set is empty, an empty string is returned.
func (s *Set[T, V]) StringWithValues() string {
	strElems := make([]string, 0, s.Size())
	for elem, value := range s.elements {
		strElems = append(strElems, fmt.Sprintf("%v (%v)", elem, value))
	}
	return strings.Join(strElems, ", ")
}

// Copy returns a new set containing all elements (including the values) of this set.
func (s *Set[T, V]) Copy() *Set[T, V] {
	newSet := NewWithValues[T, V]()
	newSet.AddAll(s)
	return newSet
}

// Intersect returns a new set containing only elements (including the values) that are in both, this set and otherSet.
// If there are no common elements or otherSet is nil, a new empty set is returned.
// Values of elements that are in both sets are taken from otherSet.
// Neither this set nor otherSet are changed.
// The values are not considered when creating the intersection.
func (s *Set[T, V]) Intersect(otherSet *Set[T, V]) *Set[T, V] {
	if otherSet == nil {
		return NewWithValues[T, V]()
	}
	newSet := NewWithValues[T, V]()
	for elem, value := range otherSet.elements {
		if s.Contains(elem) {
			newSet.AddWithValue(elem, value)
		}
	}
	return newSet
}

// Unite returns a new set containing all elements (including the values) of both, this set and otherSet.
// If otherSet is nil, a new set containing all elements of this set is returned.
// Values of elements that are in both sets are taken from otherSet.
// Neither this set nor otherSet are changed.
// The values are not considered when creating the union.
func (s *Set[T, V]) Unite(otherSet *Set[T, V]) *Set[T, V] {
	newSet := NewWithValues[T, V]()
	newSet.AddAll(s)
	newSet.AddAll(otherSet)
	return newSet
}

// UniteDisjunctively returns a new set containing all elements (including the values) that are in either this set or otherSet, but not in both (symmetric difference).
// If otherSet is nil, a new set containing all elements of this set is returned.
// Neither this set nor otherSet are changed.
// The values are not considered when creating the disjunctive union.
func (s *Set[T, V]) UniteDisjunctively(otherSet *Set[T, V]) *Set[T, V] {
	newSet := NewWithValues[T, V]()
	if otherSet == nil {
		newSet.AddAll(s)
		return newSet
	}
	for elem, value := range s.elements {
		if !otherSet.Contains(elem) {
			newSet.AddWithValue(elem, value)
		}
	}
	for elem, value := range otherSet.elements {
		if !s.Contains(elem) {
			newSet.AddWithValue(elem, value)
		}
	}
	return newSet
}

// Subtract returns a new set containing all elements (including the values) that are in this set but not in otherSet.
// If otherSet is nil, a new set containing all elements of this set is returned.
// Neither this set nor otherSet are changed.
// The values are not considered when creating the subtraction.
func (s *Set[T, V]) Subtract(otherSet *Set[T, V]) *Set[T, V] {
	newSet := NewWithValues[T, V]()
	if otherSet == nil {
		newSet.AddAll(s)
		return newSet
	}
	for elem, value := range s.elements {
		if !otherSet.Contains(elem) {
			newSet.AddWithValue(elem, value)
		}
	}
	return newSet
}

// Filter returns a new set containing only elements (including the values) for which the filter function returns true.
// If the filter function is nil, a copy of this set is returned (all elements are included because there is no filter).
// This set remains unchanged.
func (s *Set[T, V]) Filter(filter FilterFunc[T]) *Set[T, V] {
	if filter == nil {
		return s.Copy()
	}
	newSet := NewWithValues[T, V]()
	for elem, value := range s.elements {
		if filter(elem) {
			newSet.AddWithValue(elem, value)
		}
	}
	return newSet
}
