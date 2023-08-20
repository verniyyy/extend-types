package types

type Enumerable[T any, R any] interface {
	Find(func(T) bool) (T, bool)
	Filter(func(T, int) bool) R
	Each(func(*T)) R
	EachWithIndex(func(*T, int)) R
	Reject(func(T, int) bool) R
	Count(func(T) bool) int
}
