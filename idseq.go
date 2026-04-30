package gdt

import (
	"golang.org/x/exp/constraints"
)

type IDSeq[T constraints.Integer] struct {
	Current T
}

func (g *IDSeq[T]) Next() T {
	g.Current++
	id := g.Current
	return id
}
