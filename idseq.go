package gdt

type IDSeq[T integer] struct {
	Current T
}

func (g *IDSeq[T]) Next() T {
	old := g.Current
	g.Current++
	if g.Current < old {
		// Overflow happened, should start with 1.
		g.Current = 1
	}
	id := g.Current
	return id
}
