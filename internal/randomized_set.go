package internal

import "math/rand"

// RandomizedSet is a data structure that allows for O(1) average time complexity for insert, remove, and get random element operations.
type RandomizedSet[T comparable] struct {
	values  []T
	indexes map[T]int
}

func NewRandomizedSet[T comparable]() *RandomizedSet[T] {
	return &RandomizedSet[T]{
		values:  make([]T, 0),
		indexes: make(map[T]int),
	}
}

// Add adds a value to the set. Returns true if the value was added, false if it was already present.
func (rs *RandomizedSet[T]) Add(value T) bool {
	if _, ok := rs.indexes[value]; ok {
		return false
	}

	rs.indexes[value] = len(rs.values)
	rs.values = append(rs.values, value)
	return true
}

// Remove removes a value from the set. Returns true if the value was removed, false if it was not present.
func (rs *RandomizedSet[T]) Remove(value T) bool {
	if idx, ok := rs.indexes[value]; ok {
		rs.values[idx] = rs.values[len(rs.values)-1]
		rs.indexes[rs.values[idx]] = idx
		rs.values = rs.values[:len(rs.values)-1]
		delete(rs.indexes, value)
		return true
	}

	return false
}

// GetRandom returns a random value from the set. Returns the value and true if the set is not empty, otherwise returns zero value and false.
func (rs *RandomizedSet[T]) GetRandom() (T, bool) {
	if len(rs.values) == 0 {
		var zero T
		return zero, false
	}

	idx := rand.Intn(len(rs.values))
	return rs.values[idx], true
}

func (rs *RandomizedSet[T]) Len() int {
	return len(rs.values)
}

func (rs *RandomizedSet[T]) IsEmpty() bool {
	return len(rs.values) == 0
}
