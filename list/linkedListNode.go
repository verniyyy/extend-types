package list

func newNode[T any](v T) *node[T] {
	return &node[T]{
		v:    &v,
		next: nil,
	}
}

type node[T any] struct {
	v    *T
	next *node[T]
}
