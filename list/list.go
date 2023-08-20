package list

import types "github.com/verniyyy/extend-types"

type List[T any] interface {
	ReadOnlyList[T]
	Stack[T]
	Queue[T]
	types.Enumerable[T, List[T]]

	Add(T)
	Insert(int, T)
	Concat(List[T])
	Overwrite(int, T)
	Remove(int)
}

type ReadOnlyList[T any] interface {
	Print()
	Size() int
	At(int) T
	IsEmpty() bool
	Duplicate() List[T]
	DeepDuplicate() (List[T], error)
}

type Stack[T any] interface {
	Push(T)
	Pop() T
}

type Queue[T any] interface {
	Enqueue(T)
	Dequeue() T
}
