package list

import types "github.com/verniyyy/extend-types-go"

type List[T any] interface {
	ReadOnlyList[T]
	Stack[T]
	Queue[T]
	types.Enumerator[T, int, List[T]]

	Add(T)
	Insert(int, T)
	Concat(List[T])
	Overwrite(int, T)
	Fill(T)
	Remove(int)
}

type ReadOnlyList[T any] interface {
	Print()
	Size() int
	At(int) T
	Include(T) bool
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
