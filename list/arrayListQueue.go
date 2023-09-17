package list

import "github.com/verniyyy/extend-types-go/lib"

func NewQueue[T any]() Queue[T] {
	return lib.Ptr(make(arrayList[T], 0))
}

func (l *arrayList[T]) Enqueue(v T) {
	*l = append(*l, v)
}

func (l *arrayList[T]) Dequeue() T {
	popper := (*l)[0]
	*l = (*l)[1:]
	return popper
}
