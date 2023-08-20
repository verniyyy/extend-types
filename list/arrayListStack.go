package list

import "github.com/verniyyy/extend-types/lib"

func NewStack[T any]() Stack[T] {
	return lib.Ptr(make(arrayList[T], 0))
}

func (l *arrayList[T]) Push(v T) {
	*l = append(*l, v)
}

func (l *arrayList[T]) Pop() T {
	lastIndex := l.index(-1)
	popper := (*l)[lastIndex]
	*l = (*l)[:lastIndex]
	return popper
}
