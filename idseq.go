package gdt

type IDSeq[T integer] struct {
	Current T
}

func (g *IDSeq[T]) Next() T {
	g.Current++
	id := g.Current
	return id
}
