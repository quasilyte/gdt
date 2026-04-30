package gdt

import (
	"math"
)

// GenMap is an array LUT replacement that has O(1) complexity for Reset().
//
// It's keyed with uint values that are < of its size, just like an array.
// It has a very comparable get/set performance as an array too.
// It uses more memory though.
//
// When to use: instead of an array with relatively large size (e.g. 300+ elems)
// that needs to be re-used with a very cheap Reset().
// For example, if an array of size N is only ever accessed for a few reads and writes,
// so the Reset time might dominate, this GenMap can be considered.
type GenMap[T any] struct {
	elems []genMapElem[T]
	gen   uint32
}

type genMapElem[T any] struct {
	value T
	gen   uint32
}

func NewGenMap[T any](size int) *GenMap[T] {
	return &GenMap[T]{
		elems: make([]genMapElem[T], size),
		gen:   1,
	}
}

// Contains checks whether key k is bound in the array.
func (m *GenMap[T]) Contains(k uint) bool {
	if k < uint(len(m.elems)) {
		return m.elems[k].gen == m.gen
	}
	return false
}

// Get returns the value bound by the key k.
// It returns false if element is not bound.
func (m *GenMap[T]) Get2(k uint) (T, bool) {
	if k < uint(len(m.elems)) {
		el := m.elems[k]
		if el.gen == m.gen {
			return el.value, true
		}
	}
	var empty T
	return empty, false
}

// Get is a one value wrapper around Get2 that drops the bool return value.
func (m *GenMap[T]) Get(k uint) T {
	v, _ := m.Get2(k)
	return v
}

// Set assigns the value v under the specified key k.
func (m *GenMap[T]) Set(k uint, v T) {
	if k < uint(len(m.elems)) {
		m.elems[k] = genMapElem[T]{value: v, gen: m.gen}
	}
}

// Reset clears the underlying array and it's O(1).
func (m *GenMap[T]) Reset() {
	if m.gen == math.MaxUint32 {
		// For most users, this will never happen.
		// But to be safe, we need to handle this correctly.
		// m.gen will be 1, element gen will be 0.
		m.clear()
	} else {
		m.gen++
	}
}

// clear does a real array data reset.
// m.gen becomes 1.
// Every element gen becomes 0.
// This is identical to the initial array state.
//
//go:noinline - called on a cold path, therefore it should not be inlined.
func (m *GenMap[T]) clear() {
	m.gen = 1
	clear(m.elems)
}
