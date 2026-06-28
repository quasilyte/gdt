package gdt

type IDSeq[T integer] struct {
	Current T
}

func (g *IDSeq[T]) Next() T {
	return g.NextOr(1)
}

func (g *IDSeq[T]) NextOr(firstValue T) T {
	old := g.Current
	g.Current++
	if g.Current < old {
		g.Current = firstValue
	}
	id := g.Current
	return id
}
