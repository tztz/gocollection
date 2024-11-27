// An API to handle sets providing methods from set theory.
package set

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"
)

type InternalEmptyType struct{}

type FilterFunc[T comparable, V any] func(T, V) bool
type MapFunc[T comparable, V any] func(T, V) (T, V)
type MapFreeFunc[T comparable, V any, TOut comparable, VOut any] func(T, V) (TOut, VOut)
type MapToListFunc[T comparable, V any, EOut any] func(T, V) EOut
type ReduceFunc[T comparable, V any, Acc any] func(T, V, Acc) Acc

// Set is a collection of unique elements having the same type T.
// Values of type V can be associated with the elements - but don't have to.
// A Set is actually a map[T]V with keys of type T and values of type V, where the values are just associated data.
// If you don't need values, you can omit them in the Set to save memory. Then it's just a set of elements (like a set of labels).
// A Set can, of course, be empty.
// The zero value of a Set is an empty set.
type Set[T comparable, V any] interface {
	// Internal methods

	randIndex() int64

	// Public methods

	GetElements() map[T]V

	AddWithValue(T, V)
	AddWithoutValue(T)
	Remove(T)
	AddAll(Set[T, V])
	RemoveAll(Set[T, V])
	Clear()

	Size() int
	List() []T
	Contains(T) bool
	Equals(Set[T, V]) bool
	IsSubset(Set[T, V]) bool
	String() string
	StringWithValues() string

	Copy() Set[T, V]
	Intersect(Set[T, V]) Set[T, V]
	Unite(Set[T, V]) Set[T, V]
	UniteDisjunctively(Set[T, V]) Set[T, V]
	Subtract(Set[T, V]) Set[T, V]
	Filter(FilterFunc[T, V]) Set[T, V]
	Map(MapFunc[T, V]) Set[T, V]

	OneR() (T, V, error)
}

type tzSet[T comparable, V any] struct {
	elements map[T]V
}

var internalEmptyValue = InternalEmptyType{}

// NewWithValues creates a new, empty set that can contain elements of type T having values of type V (like a map).
func NewWithValues[T comparable, V any]() Set[T, V] {
	return &tzSet[T, V]{
		elements: createNewWithValues[T, V](),
	}
}

// NewWithoutValues creates a new, empty set that can contain elements of type T (like a set of labels).
func NewWithoutValues[T comparable]() Set[T, InternalEmptyType] {
	return &tzSet[T, InternalEmptyType]{
		elements: createNewWithValues[T, InternalEmptyType](),
	}
}

// createNewWithValues creates a new map with keys of type T and values of type V.
func createNewWithValues[T comparable, V any]() map[T]V {
	return make(map[T]V)
}

// randIndex returns a random index in the range of the elements.
// The index is used to get a random element from the set.
// If the set is empty, -1 is returned.
// The values are not considered when choosing the index.
// This method is not part of the public API.
func (s *tzSet[T, V]) randIndex() int64 {
	if len(s.elements) == 0 {
		return -1
	}
	n, _ := rand.Int(rand.Reader, big.NewInt(int64(len(s.elements))))
	return n.Int64()
}

// GetElements returns the internal map of elements.
func (s *tzSet[T, V]) GetElements() map[T]V {
	return s.elements
}

// AddWithValue adds an element with an associated value to the set.
func (s *tzSet[T, V]) AddWithValue(element T, value V) {
	s.elements[element] = value
}

// AddWithoutValue adds an element (without an associated value) to the set.
func (s *tzSet[T, V]) AddWithoutValue(element T) {
	var empty V
	s.elements[element] = empty
}

// Remove removes an element from the set.
func (s *tzSet[T, V]) Remove(element T) {
	delete(s.elements, element)
}

// AddAll adds all elements (including the value) from otherSet to this set.
// If otherSet is nil, nothing happens.
// If an element already exists in this set, the value is overwritten with the value from otherSet.
// The otherSet remains unchanged.
func (s *tzSet[T, V]) AddAll(otherSet Set[T, V]) {
	if otherSet == nil {
		return
	}
	for elem, value := range otherSet.GetElements() {
		s.elements[elem] = value
	}
}

// RemoveAll removes all elements from otherSet from this set.
// If otherSet is nil, nothing happens.
// The otherSet remains unchanged.
func (s *tzSet[T, V]) RemoveAll(otherSet Set[T, V]) {
	if otherSet == nil {
		return
	}
	for elem := range otherSet.GetElements() {
		delete(s.elements, elem)
	}
}

// Clear removes all elements from the set.
func (s *tzSet[T, V]) Clear() {
	clear(s.elements)
}

// Size returns the number of elements in the set.
func (s *tzSet[T, V]) Size() int {
	return len(s.elements)
}

// List returns all elements (without values) of the set as a slice.
// The returned slice is a copy, changes to that copy do not interfere with the original set.
func (s *tzSet[T, V]) List() []T {
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
func (s *tzSet[T, V]) Contains(element T) bool {
	_, exists := s.elements[element]
	return exists
}

// Equals checks if this set is equal to otherSet ignoring the values.
// Returns true if both sets are of equal size and contain the same elements (ignoring the values), false otherwise.
func (s *tzSet[T, V]) Equals(otherSet Set[T, V]) bool {
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
func (s *tzSet[T, V]) IsSubset(otherSet Set[T, V]) bool {
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
func (s *tzSet[T, V]) String() string {
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
func (s *tzSet[T, V]) StringWithValues() string {
	strElems := make([]string, 0, s.Size())
	for elem, value := range s.elements {
		strElems = append(strElems, fmt.Sprintf("%v (%v)", elem, value))
	}
	return strings.Join(strElems, ", ")
}

// Copy returns a new set containing all elements (including the values) of this set.
func (s *tzSet[T, V]) Copy() Set[T, V] {
	newSet := NewWithValues[T, V]()
	newSet.AddAll(s)
	return newSet
}

// Intersect returns a new set containing only elements (including the values) that are in both, this set and otherSet.
// If there are no common elements or otherSet is nil, a new empty set is returned.
// Values of elements that are in both sets are taken from otherSet.
// Neither this set nor otherSet are changed.
// The values are not considered when creating the intersection.
func (s *tzSet[T, V]) Intersect(otherSet Set[T, V]) Set[T, V] {
	if otherSet == nil {
		return NewWithValues[T, V]()
	}
	newSet := NewWithValues[T, V]()
	for elem, value := range otherSet.GetElements() {
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
func (s *tzSet[T, V]) Unite(otherSet Set[T, V]) Set[T, V] {
	newSet := NewWithValues[T, V]()
	newSet.AddAll(s)
	newSet.AddAll(otherSet)
	return newSet
}

// UniteDisjunctively returns a new set containing all elements (including the values) that are in either this set or otherSet, but not in both (symmetric difference).
// If otherSet is nil, a new set containing all elements of this set is returned.
// Neither this set nor otherSet are changed.
// The values are not considered when creating the disjunctive union.
func (s *tzSet[T, V]) UniteDisjunctively(otherSet Set[T, V]) Set[T, V] {
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
	for elem, value := range otherSet.GetElements() {
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
func (s *tzSet[T, V]) Subtract(otherSet Set[T, V]) Set[T, V] {
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

// Filter returns a new set containing only elements (including the values) of this set for which the filter function returns true.
// If the filter function is nil, a copy of this set is returned (all elements with their values are included because no filter applies).
// This set remains unchanged.
func (s *tzSet[T, V]) Filter(filterFunc FilterFunc[T, V]) Set[T, V] {
	if filterFunc == nil {
		return s.Copy()
	}
	newSet := NewWithValues[T, V]()
	for elem, value := range s.elements {
		if filterFunc(elem, value) {
			newSet.AddWithValue(elem, value)
		}
	}
	return newSet
}

// Map returns a new Set[T, V] containing all elements of type T (including the values of type V) returned by the map function which is applied to each element of this set.
// If the map function is nil, a copy of this set is returned (all elements with their values are included because having no mapping function behaves like identity mapping).
// This set remains unchanged.
func (s *tzSet[T, V]) Map(mapFunc MapFunc[T, V]) Set[T, V] {
	if mapFunc == nil {
		return s.Copy()
	}
	newSet := NewWithValues[T, V]()
	for elem, value := range s.elements {
		newElem, newValue := mapFunc(elem, value)
		newSet.AddWithValue(newElem, newValue)
	}
	return newSet
}

// MapFree is passed a Set[T, V] together with a map function and returns a new map[TOut]VOut containing all objects returned by the map function which is applied to each element of the set.
// If the map function is nil, an empty map is returned.
// The given set remains unchanged.
func MapFree[T comparable, V any, TOut comparable, VOut any](set Set[T, V], mapFunc MapFreeFunc[T, V, TOut, VOut]) map[TOut]VOut {
	if mapFunc == nil {
		return make(map[TOut]VOut)
	}
	newSet := make(map[TOut]VOut)
	for elem, value := range set.GetElements() {
		newElem, newValue := mapFunc(elem, value)
		newSet[newElem] = newValue
	}
	return newSet
}

// MapToList is passed a Set[T, V] together with a map function and returns a new slice containing all objects returned by the map function which is applied to each element of the set.
// If the map function is nil, an empty slice is returned.
// The given set remains unchanged.
func MapToList[T comparable, V any, EOut any](set Set[T, V], mapFunc MapToListFunc[T, V, EOut]) []EOut {
	if mapFunc == nil {
		return make([]EOut, 0)
	}
	newList := make([]EOut, 0, set.Size())
	for elem, value := range set.GetElements() {
		newListElem := mapFunc(elem, value)
		newList = append(newList, newListElem)
	}
	return newList
}

// Reduce is passed a Set[T, V] together with a reduce function as well as an initial value and returns the accumulated/reduced value successively calculated by applying the reduce function to each element of the set.
// If the reduce function is nil, the initial value is returned.
// The given set remains unchanged.
func Reduce[T comparable, V any, Acc any](set Set[T, V], reduceFunc func(T, V, Acc) Acc, initial Acc) Acc {
	if reduceFunc == nil {
		return initial
	}
	var acc = initial
	for elem, value := range set.GetElements() {
		acc = reduceFunc(elem, value, acc)
	}
	return acc
}

// OneR returns one random element (and its value) from the set.
// If the set is empty, an error is returned.
// If there is only one element in the set, that element and its value are returned.
// The values are not considered when choosing the element.
func (s *tzSet[T, V]) OneR() (T, V, error) {
	if len(s.elements) != 0 {
		rndIndex := s.randIndex()
		var counter int64 = 0

		for elem, value := range s.elements {
			if counter == rndIndex {
				return elem, value, nil
			}
			counter++
		}
	}

	var emptyT T
	var emptyV V
	return emptyT, emptyV, fmt.Errorf("cannot get a random element from set, set is empty")
}
