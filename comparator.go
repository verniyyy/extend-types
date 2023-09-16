package types

type Comparator[T comparable] interface {
	Equal(T) bool
	DeepEqual(T) bool
}
