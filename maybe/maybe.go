package maybe

import "fmt"

// Some ...
func Some[T any](v T) Maybe[T] {
	return Maybe[T]{v: v}
}

// None ...
func None[T any]() Maybe[T] {
	return Maybe[T]{none: true}
}

// Maybe ...
type Maybe[T any] struct {
	v    T
	none bool
}

// Unwrap ...
func (t Maybe[T]) Unwrap() (T, bool) {
	return t.v, !t.none
}

// DUnwrap is dangerous unwrap.
func (t Maybe[T]) DUnwrap() T {
	if t.none {
		panic(fmt.Sprintf("Attempted to unwrap Maybe, but found None. Variable %T does not contain a value.", t.v))
	}
	return t.v
}

// HasSome ...
func (t Maybe[T]) HasSome() bool {
	return !t.none
}

// None ...
func (t Maybe[T]) None() bool {
	return t.none
}

// Run ...
func Run[T, R any](
	arg Maybe[T],
	transform func(T) Maybe[R],
) Maybe[R] {
	if arg.None() {
		return None[R]()
	}
	return transform(arg.v)
}
